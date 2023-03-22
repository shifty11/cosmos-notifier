package tracker

import (
	"context"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/tracker_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//goland:noinspection GoNameStartsWithPackageName
type TrackerServer struct {
	pb.UnimplementedTrackerServiceServer
	addressTrackerManager *database.AddressTrackerManager
}

func NewTrackerServer(managers *database.DbManagers) pb.TrackerServiceServer {
	return &TrackerServer{
		addressTrackerManager: managers.AddressTrackerManager,
	}
}

func (server *TrackerServer) GetTrackers(ctx context.Context, _ *empty.Empty) (*pb.GetTrackersResponse, error) {
	userEnt, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	trackers, err := server.addressTrackerManager.GetTrackers(userEnt)
	if err != nil {
		log.Sugar.Error(err)
		return nil, status.Errorf(codes.Internal, "error while getting trackers")
	}

	var pbTrackers []*pb.Tracker
	for _, tracker := range trackers {
		discordChannelId := int64(0)
		if tracker.Edges.DiscordChannel != nil {
			discordChannelId = int64(tracker.Edges.DiscordChannel.ID)
		}
		telegramChatId := int64(0)
		if tracker.Edges.TelegramChat != nil {
			telegramChatId = int64(tracker.Edges.TelegramChat.ID)
		}
		pbTrackers = append(pbTrackers, &pb.Tracker{
			Id:                   int64(tracker.ID),
			Address:              tracker.Address,
			NotificationInterval: &duration.Duration{Seconds: tracker.NotificationInterval},
			DiscordChannelId:     discordChannelId,
			TelegramChatId:       telegramChatId,
			UpdatedAt:            &timestamp.Timestamp{Seconds: tracker.UpdateTime.Unix()},
		})
	}
	return &pb.GetTrackersResponse{Trackers: pbTrackers}, nil
}

func (server *TrackerServer) IsAddressValid(_ context.Context, req *pb.IsAddressValidRequest) (*pb.IsAddressValidResponse, error) {
	isValid, _ := server.addressTrackerManager.IsValid(req.Address)
	return &pb.IsAddressValidResponse{
		IsValid: isValid,
	}, nil
}

func (server *TrackerServer) AddTracker(ctx context.Context, req *pb.AddTrackerRequest) (*pb.Tracker, error) {
	userEnt, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}
	if req.DiscordChannelId == 0 && req.TelegramChatId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "no discord-channel-id or telegram-chat-id provided")
	}
	if req.DiscordChannelId != 0 && req.TelegramChatId != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "both discord-channel-id and telegram-chat-id provided")
	}
	if req.NotificationInterval.Seconds <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification-interval must be greater than 0")
	}

	tracker, err := server.addressTrackerManager.AddTracker(
		userEnt,
		req.Address,
		int(req.DiscordChannelId),
		int(req.TelegramChatId),
		req.NotificationInterval.Seconds,
	)
	if err != nil {
		log.Sugar.Errorf("error while adding tracker: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}

	return &pb.Tracker{
		Id:                   int64(tracker.ID),
		Address:              tracker.Address,
		NotificationInterval: &duration.Duration{Seconds: tracker.NotificationInterval},
		DiscordChannelId:     req.DiscordChannelId,
		TelegramChatId:       req.TelegramChatId,
		UpdatedAt:            &timestamp.Timestamp{Seconds: tracker.UpdateTime.Unix()},
	}, nil
}

func (server *TrackerServer) UpdateTracker(ctx context.Context, req *pb.UpdateTrackerRequest) (*pb.Tracker, error) {
	userEnt, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}
	if req.DiscordChannelId == 0 && req.TelegramChatId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "no discord-channel-id or telegram-chat-id provided")
	}
	if req.DiscordChannelId != 0 && req.TelegramChatId != 0 {
		return nil, status.Errorf(codes.InvalidArgument, "both discord-channel-id and telegram-chat-id provided")
	}
	if req.NotificationInterval.Seconds < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification-interval must be greater than 0")
	}

	tracker, err := server.addressTrackerManager.UpdateTracker(
		userEnt,
		int(req.TrackerId),
		int(req.DiscordChannelId),
		int(req.TelegramChatId),
		req.NotificationInterval.Seconds,
	)
	if err != nil {
		log.Sugar.Errorf("error while adding tracker: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}

	return &pb.Tracker{
		Id:                   int64(tracker.ID),
		Address:              tracker.Address,
		NotificationInterval: &duration.Duration{Seconds: tracker.NotificationInterval},
		DiscordChannelId:     req.DiscordChannelId,
		TelegramChatId:       req.TelegramChatId,
		UpdatedAt:            &timestamp.Timestamp{Seconds: tracker.UpdateTime.Unix()},
	}, nil
}

func (server *TrackerServer) DeleteTracker(context.Context, *pb.DeleteTrackerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTracker not implemented")
}
