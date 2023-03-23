package tracker

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/pbcommon"
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

func getTrackerChatRoom(tracker *ent.AddressTracker) (*pb.TrackerChatRoom, error) {
	chatRoom := &pb.TrackerChatRoom{
		Name: "",
	}
	if tracker.Edges.DiscordChannel != nil {
		chatRoom.Name = tracker.Edges.DiscordChannel.Name
		chatRoom.Type = &pb.TrackerChatRoom_Discord{
			Discord: &pbcommon.DiscordType{
				Id:        int32(tracker.Edges.DiscordChannel.ID),
				ChannelId: tracker.Edges.DiscordChannel.ChannelID,
			},
		}
	}
	if tracker.Edges.TelegramChat != nil {
		chatRoom.Name = tracker.Edges.TelegramChat.Name
		chatRoom.Type = &pb.TrackerChatRoom_Telegram{
			Telegram: &pbcommon.TelegramType{
				Id:     int32(tracker.Edges.TelegramChat.ID),
				ChatId: tracker.Edges.TelegramChat.ChatID,
			},
		}
	}
	if tracker.Edges.DiscordChannel != nil && tracker.Edges.TelegramChat != nil {
		err := errors.New(fmt.Sprintf("tracker %d has both discord channel and telegram chat", tracker.ID))
		log.Sugar.Error(err)
		return nil, err
	}
	if tracker.Edges.DiscordChannel == nil && tracker.Edges.TelegramChat == nil {
		err := errors.New(fmt.Sprintf("tracker %d has no discord channel or telegram chat", tracker.ID))
		log.Sugar.Error(err)
		return nil, err
	}
	return chatRoom, nil
}

func validateChatRoom(chatRoom *pb.TrackerChatRoom) error {
	if chatRoom == nil {
		return status.Errorf(codes.InvalidArgument, "chat-room must be set")
	}
	var discord = chatRoom.GetDiscord()
	if discord != nil && discord.GetId() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid discord chat-room")
	}
	var telegram = chatRoom.GetTelegram()
	if telegram != nil && telegram.GetId() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid telegram chat-room")
	}
	if chatRoom.GetType() == nil {
		return status.Errorf(codes.InvalidArgument, "chat-room-type must be set")
	}
	return nil
}

func getDiscordChannelIdOrZero(chatRoom *pb.TrackerChatRoom) int {
	switch chatRoom.Type.(type) {
	case *pb.TrackerChatRoom_Discord:
		return int(chatRoom.GetDiscord().GetId())
	default:
		return 0
	}
}

func getTelegramChatIdOrZero(chatRoom *pb.TrackerChatRoom) int {
	switch chatRoom.Type.(type) {
	case *pb.TrackerChatRoom_Telegram:
		return int(chatRoom.GetTelegram().GetId())
	default:
		return 0
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
		chatRoom, err := getTrackerChatRoom(tracker)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error while getting trackers")
		}
		pbTrackers = append(pbTrackers, &pb.Tracker{
			Id:                   int64(tracker.ID),
			Address:              tracker.Address,
			NotificationInterval: &duration.Duration{Seconds: tracker.NotificationInterval},
			ChatRoom:             chatRoom,
			UpdatedAt:            &timestamp.Timestamp{Seconds: tracker.UpdateTime.Unix()},
		})
	}

	var pbTrackerChatRooms []*pb.TrackerChatRoom
	discordChannels, telegramChats, err := server.addressTrackerManager.GetChatRooms(userEnt)
	if err != nil {
		log.Sugar.Error(err)
		return nil, status.Errorf(codes.Internal, "error while getting trackers")
	}
	for _, trackerChatRoom := range discordChannels {
		pbTrackerChatRooms = append(pbTrackerChatRooms, &pb.TrackerChatRoom{
			Name: trackerChatRoom.Name,
			Type: &pb.TrackerChatRoom_Discord{Discord: &pbcommon.DiscordType{
				Id:        int32(trackerChatRoom.ID),
				ChannelId: trackerChatRoom.ChannelID,
			}},
		})
	}
	for _, trackerChatRoom := range telegramChats {
		pbTrackerChatRooms = append(pbTrackerChatRooms, &pb.TrackerChatRoom{
			Name: trackerChatRoom.Name,
			Type: &pb.TrackerChatRoom_Telegram{Telegram: &pbcommon.TelegramType{
				Id:     int32(trackerChatRoom.ID),
				ChatId: trackerChatRoom.ChatID,
			}},
		})
	}
	return &pb.GetTrackersResponse{Trackers: pbTrackers, ChatRooms: pbTrackerChatRooms}, nil
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
	if req.NotificationInterval == nil || req.NotificationInterval.Seconds < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification-interval must be greater than 0")
	}
	if err := validateChatRoom(req.ChatRoom); err != nil {
		return nil, err
	}

	tracker, err := server.addressTrackerManager.AddTracker(
		userEnt,
		req.Address,
		getDiscordChannelIdOrZero(req.ChatRoom),
		getTelegramChatIdOrZero(req.ChatRoom),
		req.NotificationInterval.Seconds,
	)
	if err != nil {
		log.Sugar.Errorf("error while adding tracker: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	chatRoom, err := getTrackerChatRoom(tracker)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while adding tracker")
	}

	return &pb.Tracker{
		Id:                   int64(tracker.ID),
		Address:              tracker.Address,
		NotificationInterval: &duration.Duration{Seconds: tracker.NotificationInterval},
		ChatRoom:             chatRoom,
		UpdatedAt:            &timestamp.Timestamp{Seconds: tracker.UpdateTime.Unix()},
	}, nil
}

func (server *TrackerServer) UpdateTracker(ctx context.Context, req *pb.UpdateTrackerRequest) (*pb.Tracker, error) {
	userEnt, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}
	if req.NotificationInterval == nil || req.NotificationInterval.Seconds < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification-interval must be greater than 0")
	}
	if err := validateChatRoom(req.ChatRoom); err != nil {
		return nil, err
	}

	tracker, err := server.addressTrackerManager.UpdateTracker(
		userEnt,
		int(req.TrackerId),
		getDiscordChannelIdOrZero(req.ChatRoom),
		getTelegramChatIdOrZero(req.ChatRoom),
		req.NotificationInterval.Seconds,
	)
	if err != nil {
		log.Sugar.Errorf("error while adding tracker: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	chatRoom, err := getTrackerChatRoom(tracker)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while updating tracker")
	}

	return &pb.Tracker{
		Id:                   int64(tracker.ID),
		Address:              tracker.Address,
		NotificationInterval: &duration.Duration{Seconds: tracker.NotificationInterval},
		ChatRoom:             chatRoom,
		UpdatedAt:            &timestamp.Timestamp{Seconds: tracker.UpdateTime.Unix()},
	}, nil
}

func (server *TrackerServer) DeleteTracker(ctx context.Context, req *pb.DeleteTrackerRequest) (*empty.Empty, error) {
	userEnt, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	if err := server.addressTrackerManager.DeleteTracker(userEnt, int(req.TrackerId)); err != nil {
		log.Sugar.Errorf("error while deleting tracker: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}

	return &empty.Empty{}, nil
}
