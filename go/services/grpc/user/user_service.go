package user

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/shifty11/cosmos-notifier/common"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/log"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/user_service"
	"github.com/shifty11/cosmos-notifier/services/grpc/types"
)

//goland:noinspection GoNameStartsWithPackageName
type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() pb.UserServiceServer {
	return &UserServer{}
}

func (u UserServer) GetUser(ctx context.Context, _ *empty.Empty) (*pb.User, error) {
	userEnt, ok := ctx.Value(common.ContextKeyUser).(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, types.UserNotFoundErr
	}

	return &pb.User{
		Name:   userEnt.Name,
		Avatar: "",
	}, nil
}
