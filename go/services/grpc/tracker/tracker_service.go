package tracker

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/shifty11/cosmos-notifier/database"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/tracker_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//goland:noinspection GoNameStartsWithPackageName
type TrackerServer struct {
	pb.UnimplementedTrackerServiceServer
	addressTrackerManager *database.AddressTrackerManager
}

func NewTrackersServer(managers *database.DbManagers) pb.TrackerServiceServer {
	return &TrackerServer{
		addressTrackerManager: managers.AddressTrackerManager,
	}
}

func (server *TrackerServer) IsAddressValid(_ context.Context, req *pb.IsAddressValidRequest) (*pb.IsAddressValidResponse, error) {
	isValid, _ := server.addressTrackerManager.IsValid(req.Address)
	return &pb.IsAddressValidResponse{
		IsValid: isValid,
	}, nil
}
func (server *TrackerServer) AddTracker(context.Context, *pb.AddTrackerRequest) (*pb.AddTrackerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTracker not implemented")
}
func (server *TrackerServer) DeleteTracker(context.Context, *pb.DeleteTrackerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTracker not implemented")
}
