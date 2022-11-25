package subscription

import (
	"context"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/services/contract_crawler"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/subscription_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

//goland:noinspection GoNameStartsWithPackageName
type SubscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
	subscriptionManager *database.SubscriptionManager
	contractManager     database.IContractManager
	chainManager        *database.ChainManager
	crawlerClient       *contract_crawler.ContractCrawler
}

func NewSubscriptionsServer(managers *database.DbManagers, crawlerClient *contract_crawler.ContractCrawler) pb.SubscriptionServiceServer {
	return &SubscriptionServer{
		subscriptionManager: managers.SubscriptionManager,
		contractManager:     managers.ContractManager,
		chainManager:        managers.ChainManager,
		crawlerClient:       crawlerClient,
	}
}

func (server *SubscriptionServer) GetSubscriptions(ctx context.Context, _ *emptypb.Empty) (*pb.GetSubscriptionsResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	log.Sugar.Debugf("GetSubscriptions for user %v (%v)", entUser.Name, entUser.UserID)

	res := server.subscriptionManager.GetSubscriptions(entUser)
	return res, nil
}

func (server *SubscriptionServer) ToggleChainSubscription(ctx context.Context, req *pb.ToggleChainSubscriptionRequest) (*pb.ToggleSubscriptionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	log.Sugar.Debugf("ToggleChainSubscription %v for user %v (%v)", req.ChainId, entUser.Name, entUser.UserID)

	isSubscribed, err := server.subscriptionManager.ToggleChainSubscription(entUser, req.ChatRoomId, int(req.ChainId))
	if err != nil {
		log.Sugar.Errorf("error while toggling subscription: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	var res = &pb.ToggleSubscriptionResponse{IsSubscribed: isSubscribed}
	return res, nil
}

func (server *SubscriptionServer) ToggleContractSubscription(ctx context.Context, req *pb.ToggleContractSubscriptionRequest) (*pb.ToggleSubscriptionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	log.Sugar.Debugf("ToggleContractSubscription %v for user %v (%v)", req.ContractId, entUser.Name, entUser.UserID)

	isSubscribed, err := server.subscriptionManager.ToggleContractSubscription(entUser, req.ChatRoomId, int(req.ContractId))
	if err != nil {
		log.Sugar.Errorf("error while toggling subscription: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	var res = &pb.ToggleSubscriptionResponse{IsSubscribed: isSubscribed}
	return res, nil
}

func (server *SubscriptionServer) AddDao(req *pb.AddDaoRequest, stream pb.SubscriptionService_AddDaoServer) error {
	entUser, ok := stream.Context().Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return status.Errorf(codes.NotFound, "invalid user")
	}

	if req.ContractAddress == "" {
		log.Sugar.Error("invalid contract address")
		return status.Errorf(codes.InvalidArgument, "invalid contract address")
	}

	var chain *ent.Chain
	for _, c := range server.chainManager.All() {
		if c.Display != "" && strings.HasPrefix(req.ContractAddress, c.Display) {
			chain = c
			break
		}
	}

	if chain == nil {
		log.Sugar.Error("invalid contract address: no chain found")
		return status.Errorf(codes.InvalidArgument, "invalid contract address: no chain found")
	}

	log.Sugar.Debugf("AddDao %v by user %v (%v)", req.ContractAddress, entUser.Name, entUser.UserID)

	contract, err := server.crawlerClient.ByAddress(req.ContractAddress)
	if err != nil && !ent.IsNotFound(err) {
		log.Sugar.Errorf("error while getting dao: %v", err)
		return status.Errorf(codes.Internal, "Unknown error occurred")
	}
	if contract != nil {
		err = stream.Send(&pb.AddDaoResponse{
			Status:          pb.AddDaoResponse_ALREADY_ADDED,
			Name:            contract.Name,
			ContractAddress: contract.Address,
		})
		if err != nil {
			log.Sugar.Errorf("error while sending response: %v", err)
		}
		return err
	}

	err = stream.Send(&pb.AddDaoResponse{Status: pb.AddDaoResponse_IS_ADDING})
	if err != nil {
		log.Sugar.Errorf("error while sending response: %v", err)
		return err
	}

	waitc := make(chan *ent.Contract)
	go func() {
		contract, err = server.crawlerClient.AddContract(chain, req.ContractAddress, req.CustomQuery)
		if err != nil {
			log.Sugar.Errorf("error while adding dao: %v", req.ContractAddress)
		}
		waitc <- contract
		close(waitc)
	}()

	contract = <-waitc
	if contract == nil {
		err = stream.Send(&pb.AddDaoResponse{Status: pb.AddDaoResponse_FAILED, ContractAddress: req.GetContractAddress()})
		if err != nil {
			log.Sugar.Errorf("error while sending response: %v", err)
		}
		return err
	} else {
		err = stream.Send(&pb.AddDaoResponse{
			Status:          pb.AddDaoResponse_ADDED,
			Name:            contract.Name,
			ContractAddress: contract.Address,
		})
		if err != nil {
			log.Sugar.Errorf("error while sending response: %v", err)
		}
		return err
	}
}

func (server *SubscriptionServer) DeleteDao(_ context.Context, req *pb.DeleteDaoRequest) (*emptypb.Empty, error) {
	if req.GetContractId() == 0 {
		log.Sugar.Error("invalid contract id")
		return nil, status.Errorf(codes.InvalidArgument, "invalid contract id")
	}

	log.Sugar.Debugf("Delete DAO %v", req.ContractId)
	err := server.contractManager.Delete(int(req.ContractId))
	if err != nil && ent.IsNotFound(err) {
		return nil, status.Errorf(codes.NotFound, "contract not found")
	}
	if err != nil {
		log.Sugar.Errorf("error while deleting dao: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}

	return &emptypb.Empty{}, nil
}

func (server *SubscriptionServer) EnableChain(_ context.Context, req *pb.EnableChainRequest) (*emptypb.Empty, error) {
	if req.GetChainId() == 0 {
		log.Sugar.Error("invalid chain id")
		return nil, status.Errorf(codes.InvalidArgument, "invalid chain id")
	}

	log.Sugar.Debugf("Delete DAO %v", req.ChainId)
	err := server.chainManager.Enable(int(req.ChainId), req.IsEnabled)
	if err != nil && ent.IsNotFound(err) {
		return nil, status.Errorf(codes.NotFound, "chain not found")
	}
	if err != nil {
		log.Sugar.Errorf("error while deleting dao: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}

	return &emptypb.Empty{}, nil
}
