package dev

import (
	"context"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/services/grpc/auth"
	authpb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/auth_service"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/dev_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

//goland:noinspection GoNameStartsWithPackageName
type DevServer struct {
	pb.UnimplementedDevServiceServer
	userManager *database.UserManager
	jwtManager  *auth.JWTManager
}

func NewDevServer(managers *database.DbManagers, jwtManager *auth.JWTManager) pb.DevServiceServer {
	return &DevServer{
		userManager: managers.UserManager,
		jwtManager:  jwtManager,
	}
}

func (s *DevServer) login(entUser *ent.User) (*authpb.LoginResponse, error) {
	accessToken, err := s.jwtManager.GenerateToken(entUser, auth.AccessToken)
	if err != nil {
		log.Sugar.Errorf("Could not generate accessToken for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, auth.ErrorLoginFailed
	}

	refreshToken, err := s.jwtManager.GenerateToken(entUser, auth.RefreshToken)
	if err != nil {
		log.Sugar.Errorf("Could not generate refreshToken for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, auth.ErrorInternal
	}

	res := &authpb.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return res, nil
}

func (s *DevServer) setRoleIfNecessary(req *pb.DevLoginRequest, userEnt *ent.User) (*ent.User, error) {
	var desiredRole = user.Role(strings.ToLower(req.GetRole().String()))
	if desiredRole != userEnt.Role {
		_, err := s.userManager.SetRole(userEnt.Name, desiredRole)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to set role: %v", err)
		}
		userEnt, err = s.userManager.Get(userEnt.UserID, userEnt.Type)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to get user: %v", err)
		}
	}
	return userEnt, nil
}

func (s *DevServer) Login(_ context.Context, req *pb.DevLoginRequest) (*authpb.LoginResponse, error) {
	client, ctx := database.ConnectDuringDevelopment()
	if req.GetUserId() != 0 {
		var userQuery = client.User.Query()
		if req.GetType() == pb.DevLoginRequest_TELEGRAM {
			userQuery = userQuery.Where(
				user.And(
					user.HasTelegramChats(),
					user.IDEQ(int(req.GetUserId())),
				),
			)
		} else {
			userQuery = userQuery.Where(
				user.And(
					user.HasDiscordChannels(),
					user.IDEQ(int(req.GetUserId())),
				),
			)
		}
		userEnt, err := userQuery.First(ctx)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to get user: %v", err)
		}
		if userEnt, err = s.setRoleIfNecessary(req, userEnt); err != nil {
			return nil, err
		}
		return s.login(userEnt)
	} else {
		userQuery := client.User.Query()
		if req.GetType() == pb.DevLoginRequest_TELEGRAM {
			userQuery = userQuery.Where(user.HasTelegramChats())
		} else {
			userQuery = userQuery.Where(user.HasDiscordChannels())
		}
		userEnt, err := userQuery.First(ctx)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to get user: %v", err)
		}
		if userEnt, err = s.setRoleIfNecessary(req, userEnt); err != nil {
			return nil, err
		}
		return s.login(userEnt)
	}
}
