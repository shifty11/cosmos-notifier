package subscription

import (
	"context"
	"fmt"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
	pb "github.com/shifty11/dao-dao-notifier/service_grpc/protobuf/go/subscription_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//goland:noinspection GoNameStartsWithPackageName
type SubscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
	subscriptionManager *database.SubscriptionManager
}

func NewSubscriptionsServer(subscriptionManager *database.SubscriptionManager) pb.SubscriptionServiceServer {
	return &SubscriptionServer{subscriptionManager: subscriptionManager}
}

func convertSubscriptionToProtobuf(entUser *ent.User, subscriptions []*database.ChatRoom) []*pb.ChatRoom {
	var rooms []*pb.ChatRoom
	for _, chatRoom := range subscriptions {
		var subs []*pb.Subscription
		for _, sub := range chatRoom.Subscriptions {
			thumbUrl := sub.ThumbnailUrl
			if thumbUrl == "" {
				thumbUrl = "images/logo.png"
			} else {
				//TODO: make this configurable
				thumbUrl = fmt.Sprintf("https://api.decrypto.online/%v", thumbUrl)
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

	log.Sugar.Debugf("ToggleSubscription on %v for user %v (%v)", req.ContractId, entUser.Name, entUser.UserID)

	isSubscribed, err := server.subscriptionManager.ToggleSubscription(entUser, req.ChatRoomId, int(req.ContractId))
	if err != nil {
		log.Sugar.Errorf("error while toggling subscription: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	var res = &pb.ToggleSubscriptionResponse{IsSubscribed: isSubscribed}
	return res, nil
}
