package grpc

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	crawler "github.com/shifty11/dao-dao-notifier/service_crawler"
	"github.com/shifty11/dao-dao-notifier/service_grpc/auth"
	"github.com/shifty11/dao-dao-notifier/service_grpc/protobuf/go/auth_service"
	"github.com/shifty11/dao-dao-notifier/service_grpc/protobuf/go/subscription_service"
	"github.com/shifty11/dao-dao-notifier/service_grpc/subscription"
	"golang.org/x/oauth2"
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
	DiscordOAuth2Config  *oauth2.Config
}

//goland:noinspection GoNameStartsWithPackageName
type GRPCServer struct {
	dbManagers    *database.DbManagers
	config        *Config
	crawlerClient *crawler.Crawler
}

func NewGRPCServer(dbManagers *database.DbManagers, config *Config, crawlerClient *crawler.Crawler) *GRPCServer {
	return &GRPCServer{dbManagers: dbManagers, config: config, crawlerClient: crawlerClient}
}

func (s GRPCServer) Run() {
	log.Sugar.Info("Starting GRPC server")
	jwtManager := auth.NewJWTManager([]byte(s.config.JwtSecretKey), s.config.AccessTokenDuration, s.config.RefreshTokenDuration)
	interceptor := auth.NewAuthInterceptor(jwtManager, s.dbManagers.UserManager, auth.AccessibleRoles())

	authServer := auth.NewAuthServer(s.dbManagers.UserManager, jwtManager, s.config.TelegramToken, s.config.DiscordOAuth2Config)
	subsServer := subscription.NewSubscriptionsServer(s.dbManagers, s.crawlerClient)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	auth_service.RegisterAuthServiceServer(server, authServer)
	subscription_service.RegisterSubscriptionServiceServer(server, subsServer)

	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
