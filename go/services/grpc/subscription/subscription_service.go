package subscription

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/services/contract_crawler"
	pb "github.com/shifty11/dao-dao-notifier/services/grpc/protobuf/go/subscription_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//goland:noinspection GoNameStartsWithPackageName
type SubscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
	subscriptionManager *database.SubscriptionManager
	contractManager     database.IContractManager
	crawlerClient       *contract_crawler.ContractCrawler
}

func NewSubscriptionsServer(managers *database.DbManagers, crawlerClient *contract_crawler.ContractCrawler) pb.SubscriptionServiceServer {
	return &SubscriptionServer{
		subscriptionManager: managers.SubscriptionManager,
		contractManager:     managers.ContractManager,
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

	chatRooms := server.subscriptionManager.GetSubscriptions(entUser)

	var res = &pb.GetSubscriptionsResponse{ChatRooms: chatRooms}
	return res, nil
}

func (server *SubscriptionServer) ToggleSubscription(ctx context.Context, req *pb.ToggleSubscriptionRequest) (*pb.ToggleSubscriptionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	log.Sugar.Debugf("ToggleSubscription %v for user %v (%v)", req.ContractId, entUser.Name, entUser.UserID)

	isSubscribed, err := server.subscriptionManager.ToggleSubscription(entUser, req.ChatRoomId, int(req.ContractId))
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

	if len(req.ContractAddress) != 63 {
		log.Sugar.Error("invalid contract address")
		return status.Errorf(codes.InvalidArgument, "invalid contract address")
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
		contract, err = server.crawlerClient.AddContract(req.ContractAddress)
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
