package subscription

import (
	"context"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
	crawler "github.com/shifty11/dao-dao-notifier/service_crawler"
	pb "github.com/shifty11/dao-dao-notifier/service_grpc/protobuf/go/subscription_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//goland:noinspection GoNameStartsWithPackageName
type SubscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
	subscriptionManager *database.SubscriptionManager
	contractManager     database.IContractManager
	crawlerClient       *crawler.Crawler
}

func NewSubscriptionsServer(managers *database.DbManagers, crawlerClient *crawler.Crawler) pb.SubscriptionServiceServer {
	return &SubscriptionServer{
		subscriptionManager: managers.SubscriptionManager,
		contractManager:     managers.ContractManager,
		crawlerClient:       crawlerClient,
	}
}

func convertSubscriptionToProtobuf(entUser *ent.User, subscriptions []*database.ChatRoom) []*pb.ChatRoom {
	var rooms []*pb.ChatRoom
	for _, chatRoom := range subscriptions {
		var subs []*pb.Subscription
		for _, sub := range chatRoom.Subscriptions {
			thumbUrl := sub.ThumbnailUrl
			if thumbUrl == "" {
				thumbUrl = "images/logo.png"
			}
			subs = append(subs, &pb.Subscription{
				Id:              sub.Id,
				Name:            sub.Name,
				IsSubscribed:    sub.Notify,
				ThumbnailUrl:    thumbUrl,
				ContractAddress: sub.ContractAddress,
			})
		}
		roomType := pb.ChatRoom_TELEGRAM
		if entUser.Type == user.TypeDiscord {
			roomType = pb.ChatRoom_DISCORD
		}
		rooms = append(rooms, &pb.ChatRoom{
			Id:            chatRoom.Id,
			Name:          chatRoom.Name,
			TYPE:          roomType,
			Subscriptions: subs,
		})
	}
	return rooms
}

func (server *SubscriptionServer) GetSubscriptions(ctx context.Context, _ *emptypb.Empty) (*pb.GetSubscriptionsResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	log.Sugar.Debugf("GetSubscriptions for user %v (%v)", entUser.Name, entUser.UserID)

	subs := server.subscriptionManager.GetSubscriptions(entUser)
	chatRooms := convertSubscriptionToProtobuf(entUser, subs)

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
