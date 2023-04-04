package tracker

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/shifty11/cosmos-notifier/common"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/services/grpc/protobuf/pbcommon"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/tracker_service"
	"github.com/shifty11/cosmos-notifier/services/grpc/types"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

//goland:noinspection GoNameStartsWithPackageName
type TrackerServer struct {
	pb.UnimplementedTrackerServiceServer
	addressTrackerManager *database.AddressTrackerManager
	validatorManager      *database.ValidatorManager
}

func NewTrackerServer(managers *database.DbManagers) pb.TrackerServiceServer {
	return &TrackerServer{
		addressTrackerManager: managers.AddressTrackerManager,
		validatorManager:      managers.ValidatorManager,
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

func toValidatorBundles(validators []*ent.Validator, trackedMonikers map[string]bool) []*pb.ValidatorBundle {
	bundles := make(map[string]*pb.ValidatorBundle)
	for _, validator := range validators {
		var pbValidator = &pb.Validator{
			Id:        int64(validator.ID),
			Address:   validator.Address,
			ChainName: validator.Edges.Chain.PrettyName,
		}
		if _, ok := bundles[validator.Moniker]; ok {
			bundles[validator.Moniker].Validators = append(bundles[validator.Moniker].Validators, pbValidator)
		} else {
			var isTracked = trackedMonikers[validator.Moniker]
			bundles[validator.Moniker] = &pb.ValidatorBundle{
				Moniker:    strings.Trim(validator.Moniker, " \t"), // remove leading and trailing spaces
				Validators: []*pb.Validator{pbValidator},
				IsTracked:  isTracked,
			}
		}
	}
	validatorBundles := make([]*pb.ValidatorBundle, 0, len(bundles))
	for _, bundle := range bundles {
		validatorBundles = append(validatorBundles, bundle)
	}
	//sort.Slice(validatorBundles, func(i, j int) bool {
	//	return len(validatorBundles[i].Validators) > len(validatorBundles[j].Validators) ||
	//		validatorBundles[i].Moniker < validatorBundles[j].Moniker
	//})
	return validatorBundles
}

func (server *TrackerServer) ListTrackers(ctx context.Context, _ *empty.Empty) (*pb.ListTrackersResponse, error) {
	userEnt, ok := ctx.Value(common.ContextKeyUser).(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, types.UserNotFoundErr
	}

	trackers, err := server.addressTrackerManager.QueryByUser(userEnt)
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
		var validatorMoniker string
		if tracker.Edges.Validator != nil {
			validatorMoniker = tracker.Edges.Validator.Moniker
		}
		pbTrackers = append(pbTrackers, &pb.Tracker{
			Id:                   int64(tracker.ID),
			Address:              tracker.Address,
			NotificationInterval: &duration.Duration{Seconds: tracker.NotificationInterval},
			ChatRoom:             chatRoom,
			UpdatedAt:            &timestamp.Timestamp{Seconds: tracker.UpdateTime.Unix()},
			ValidatorMoniker:     validatorMoniker,
		})
	}

	var pbTrackerChatRooms []*pb.TrackerChatRoom
	var trackedMonikers = make(map[string]bool)
	discordChannels, telegramChats, err := server.addressTrackerManager.QueryChatRooms(userEnt)
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
		if trackerChatRoom.Edges.Validators != nil {
			for _, validator := range trackerChatRoom.Edges.Validators {
				trackedMonikers[validator.Moniker] = true
			}
		}
	}
	for _, trackerChatRoom := range telegramChats {
		pbTrackerChatRooms = append(pbTrackerChatRooms, &pb.TrackerChatRoom{
			Name: trackerChatRoom.Name,
			Type: &pb.TrackerChatRoom_Telegram{Telegram: &pbcommon.TelegramType{
				Id:     int32(trackerChatRoom.ID),
				ChatId: trackerChatRoom.ChatID,
			}},
		})
		if trackerChatRoom.Edges.Validators != nil {
			for _, validator := range trackerChatRoom.Edges.Validators {
				trackedMonikers[validator.Moniker] = true
			}
		}
	}

	validators := server.validatorManager.QueryActive()
	pbValidatorBundles := toValidatorBundles(validators, trackedMonikers)

	return &pb.ListTrackersResponse{
		Trackers:         pbTrackers,
		ChatRooms:        pbTrackerChatRooms,
		ValidatorBundles: pbValidatorBundles,
	}, nil
}

func (server *TrackerServer) IsAddressValid(_ context.Context, req *pb.IsAddressValidRequest) (*pb.IsAddressValidResponse, error) {
	isValid, _ := server.addressTrackerManager.QueryIsValid(req.Address)
	return &pb.IsAddressValidResponse{
		IsValid: isValid,
	}, nil
}

func (server *TrackerServer) CreateTracker(ctx context.Context, req *pb.CreateTrackerRequest) (*pb.Tracker, error) {
	userEnt, ok := ctx.Value(common.ContextKeyUser).(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, types.UserNotFoundErr
	}
	if req.NotificationInterval == nil || req.NotificationInterval.Seconds < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification-interval must be greater than 0")
	}
	if err := validateChatRoom(req.ChatRoom); err != nil {
		return nil, err
	}

	tracker, err := server.addressTrackerManager.Create(
		ctx,
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
	userEnt, ok := ctx.Value(common.ContextKeyUser).(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, types.UserNotFoundErr
	}
	if req.NotificationInterval == nil || req.NotificationInterval.Seconds < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "notification-interval must be greater than 0")
	}
	if err := validateChatRoom(req.ChatRoom); err != nil {
		return nil, err
	}

	tracker, err := server.addressTrackerManager.Update(
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
	userEnt, ok := ctx.Value(common.ContextKeyUser).(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, types.UserNotFoundErr
	}

	if err := server.addressTrackerManager.Delete(userEnt, int(req.TrackerId)); err != nil {
		log.Sugar.Errorf("error while deleting tracker: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}

	return &empty.Empty{}, nil
}

func (server *TrackerServer) getToBeUntrackedValidators(req *pb.TrackValidatorsRequest, previousValidators []*ent.Validator) []*ent.Validator {
	var validators []*ent.Validator
	for _, previousValidator := range previousValidators {
		if !slices.Contains(req.GetMonikers(), previousValidator.Moniker) {
			validators = append(validators, previousValidator)
		}
	}
	return validators
}

func (server *TrackerServer) getToBeTrackedValidators(req *pb.TrackValidatorsRequest, previousValidators []*ent.Validator) []*ent.Validator {
	var newValidators []*ent.Validator
	for _, moniker := range req.GetMonikers() {
		validators := server.validatorManager.QueryByMoniker(moniker)
		for _, validator := range validators {
			containsFunc := func(prev *ent.Validator) bool {
				return prev.Address == validator.Address
			}
			if !slices.ContainsFunc(previousValidators, containsFunc) {
				newValidators = append(newValidators, validator)
			}
		}
	}
	return newValidators
}

func (server *TrackerServer) TrackValidators(ctx context.Context, req *pb.TrackValidatorsRequest) (*pb.TrackValidatorsResponse, error) {
	userEnt, ok := ctx.Value(common.ContextKeyUser).(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, types.UserNotFoundErr
	}

	var deletedTrackerIds []int64
	previousValidators, err := server.validatorManager.QueryByUser(userEnt)
	if err != nil {
		log.Sugar.Errorf("error while getting previous validators: %v", err)
		return nil, types.UnknownErr
	}

	var toBeTrackedValidators = server.getToBeTrackedValidators(req, previousValidators)
	if len(toBeTrackedValidators) > 0 {
		if req.NotificationInterval == nil || req.NotificationInterval.Seconds < 0 {
			return nil, status.Errorf(codes.InvalidArgument, "notification-interval must be greater than 0")
		}
		if err := validateChatRoom(req.ChatRoom); err != nil {
			return nil, err
		}
	}
	var toBeUntrackedValidators = server.getToBeUntrackedValidators(req, previousValidators)

	ctx, err = server.validatorManager.StartTx(ctx)
	if err != nil {
		log.Sugar.Errorf("error while starting transaction: %v", err)
		return nil, types.UnknownErr
	}
	defer func() {
		_, err := database.RollbackTxIfUncommitted(ctx)
		if err != nil {
			log.Sugar.Errorf("error while rolling back transaction: %v", err)
		}
	}()

	for _, previousValidator := range toBeUntrackedValidators {
		result, err := server.validatorManager.UpdateUntrackValidator(ctx, userEnt, previousValidator)
		if err != nil {
			log.Sugar.Errorf("error while untracking validator: %v", err)
			return nil, types.UnknownErr
		}
		for _, id := range result {
			deletedTrackerIds = append(deletedTrackerIds, int64(id))
		}
	}

	var trackers []*pb.Tracker
	for _, newValidator := range toBeTrackedValidators {
		discordChannelId := getDiscordChannelIdOrZero(req.ChatRoom)
		telegramChatId := getTelegramChatIdOrZero(req.ChatRoom)
		tracker, err := server.validatorManager.UpdateTrackValidator(
			ctx,
			userEnt,
			newValidator,
			discordChannelId,
			telegramChatId,
			req.NotificationInterval.Seconds,
		)
		if err != nil {
			log.Sugar.Errorf("error while tracking validator: %v", err)
			return nil, types.UnknownErr
		}
		if tracker != nil {
			chatRoom, err := getTrackerChatRoom(tracker)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "error while adding tracker")
			}
			var validatorMoniker string
			if tracker.Edges.Validator != nil {
				validatorMoniker = tracker.Edges.Validator.Moniker
			}
			trackers = append(trackers, &pb.Tracker{
				Id:                   int64(tracker.ID),
				Address:              tracker.Address,
				NotificationInterval: &duration.Duration{Seconds: tracker.NotificationInterval},
				ChatRoom:             chatRoom,
				UpdatedAt:            &timestamp.Timestamp{Seconds: tracker.UpdateTime.Unix()},
				ValidatorMoniker:     validatorMoniker,
			})
		}
	}
	ctx, err = database.CommitTx(ctx)
	if err != nil {
		log.Sugar.Errorf("error while committing transaction: %v", err)
		return nil, types.UnknownErr
	}

	return &pb.TrackValidatorsResponse{
		AddedTrackers:     trackers,
		DeletedTrackerIds: deletedTrackerIds,
	}, nil
}
