package grpc

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/service_grpc/auth"
	"github.com/shifty11/dao-dao-notifier/service_grpc/protobuf/go/auth_service"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Config struct {
	Port                 string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	JwtSecretKey         string
	TelegramToken        string
}

//goland:noinspection GoNameStartsWithPackageName
type GRPCServer struct {
	dbManagers *database.DbManagers
	config     *Config
}

func NewGRPCServer(dbManagers *database.DbManagers, config *Config) *GRPCServer {
	return &GRPCServer{dbManagers: dbManagers, config: config}
}

func (s GRPCServer) Run() {
	log.Sugar.Info("Starting GRPC server")
	jwtManager := auth.NewJWTManager([]byte(s.config.JwtSecretKey), s.config.AccessTokenDuration, s.config.RefreshTokenDuration)
	interceptor := auth.NewAuthInterceptor(jwtManager, s.dbManagers.UserManager, auth.AccessibleRoles())

	authServer := auth.NewAuthServer(s.dbManagers.UserManager, jwtManager, s.config.TelegramToken)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	auth_service.RegisterAuthServiceServer(server, authServer)

	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
