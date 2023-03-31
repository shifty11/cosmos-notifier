package admin

import (
	"context"
	"fmt"
	"github.com/shifty11/cosmos-notifier/common"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/notifier"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/admin_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//goland:noinspection GoNameStartsWithPackageName
type AdminServer struct {
	pb.UnimplementedAdminServiceServer
	notifier     notifier.GeneralNotifier
	statsManager *database.StatsManager
}

func NewAdminServer(notifier notifier.GeneralNotifier, managers *database.DbManagers) pb.AdminServiceServer {
	return &AdminServer{
		notifier:     notifier,
		statsManager: managers.StatsManager,
	}
}

func (server *AdminServer) BroadcastMessage(req *pb.BroadcastMessageRequest, stream pb.AdminService_BroadcastMessageServer) error {
	entUser, ok := stream.Context().Value(common.ContextKeyUser).(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return status.Errorf(codes.NotFound, "invalid user")
	}

	if req.Message == "" {
		log.Sugar.Error("message is empty")
		return status.Errorf(codes.InvalidArgument, "message is empty")
	}

	waitc := make(chan notifier.BroadcastMessageResult)

	go func() {
		defer close(waitc)
		result := server.notifier.BroadcastMessage(req.Message, req.Type, entUser, waitc)
		waitc <- result
	}()

	for result := range waitc {
		if result.Error != nil {
			log.Sugar.Errorf("error while broadcasting message: %v", result.Error)
			return status.Errorf(codes.Internal, "error while broadcasting message: %v", result.Error)
		}
		if result.IsSending {
			var response = ""
			if req.Type == pb.BroadcastMessageRequest_TELEGRAM {
				response = fmt.Sprintf("Broadcast message to %v telegram chats", result.ChatCnt)
			} else if req.Type == pb.BroadcastMessageRequest_DISCORD {
				response = fmt.Sprintf("Broadcast message to %v discord channels", result.ChatCnt)
			}
			err := stream.Send(&pb.BroadcastMessageResponse{Status: pb.BroadcastMessageResponse_SENDING, Response: response})
			if err != nil {
				log.Sugar.Errorf("error while sending response: %v", err)
				return err
			}
		} else {
			var response = ""
			if req.Type == pb.BroadcastMessageRequest_TELEGRAM_TEST {
				response = fmt.Sprintf("Sent message to telegram chat %v", result.SingleChatName)
			} else if req.Type == pb.BroadcastMessageRequest_DISCORD_TEST {
				response = fmt.Sprintf("Sent message to discord channel %v", result.SingleChatName)
			} else if req.Type == pb.BroadcastMessageRequest_TELEGRAM {
				response = fmt.Sprintf("Broadcasted message to %v telegram chats\n%v chats were deleted", result.ChatCnt, result.ErrorCnt)
			} else if req.Type == pb.BroadcastMessageRequest_DISCORD {
				response = fmt.Sprintf("Broadcasted message to %v discord channels\n%v channels were deleted", result.ChatCnt, result.ErrorCnt)
			}
			err := stream.Send(&pb.BroadcastMessageResponse{Status: pb.BroadcastMessageResponse_SENT, Response: response})
			if err != nil {
				log.Sugar.Errorf("error while sending response: %v", err)
				return err
			}
		}
	}
	return nil
}

func (server *AdminServer) GetStats(_ context.Context, _ *emptypb.Empty) (*pb.GetStatsResponse, error) {
	stats, err := server.statsManager.GetStats()
	if err != nil {
		log.Sugar.Errorf("error while getting stats: %v", err)
		return nil, status.Errorf(codes.Internal, "error while getting stats: %v", err)
	}

	return &pb.GetStatsResponse{
		TelegramChats:         int32(stats.TelegramChats),
		DiscordChannels:       int32(stats.DiscordChannels),
		TelegramUsers:         int32(stats.TelegramUsers),
		DiscordUsers:          int32(stats.DiscordUsers),
		TelegramSubscriptions: int32(stats.TelegramSubscriptions),
		DiscordSubscriptions:  int32(stats.DiscordSubscriptions),
		Subscriptions:         int32(stats.Subscriptions),
		Users:                 int32(stats.Users),
		Contracts:             int32(stats.Contracts),
		Chains:                int32(stats.Chains),
	}, nil
}
