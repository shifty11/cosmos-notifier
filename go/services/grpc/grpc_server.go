package grpc

import (
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/shifty11/cosmos-notifier/notifier"
	"github.com/shifty11/cosmos-notifier/services/contract_crawler"
	"github.com/shifty11/cosmos-notifier/services/grpc/admin"
	"github.com/shifty11/cosmos-notifier/services/grpc/auth"
	"github.com/shifty11/cosmos-notifier/services/grpc/dev"
	"github.com/shifty11/cosmos-notifier/services/grpc/protobuf/admin_service"
	"github.com/shifty11/cosmos-notifier/services/grpc/protobuf/auth_service"
	"github.com/shifty11/cosmos-notifier/services/grpc/protobuf/dev_service"
	"github.com/shifty11/cosmos-notifier/services/grpc/protobuf/subscription_service"
	"github.com/shifty11/cosmos-notifier/services/grpc/protobuf/tracker_service"
	"github.com/shifty11/cosmos-notifier/services/grpc/subscription"
	"github.com/shifty11/cosmos-notifier/services/grpc/tracker"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

type Config struct {
	Port                 string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	JwtSecretKey         string
	TelegramToken        string
	DiscordOAuth2Config  *oauth2.Config
	CannyPrivateKey      string
}

//goland:noinspection GoNameStartsWithPackageName
type GRPCServer struct {
	dbManagers      *database.DbManagers
	config          *Config
	crawlerClient   *contract_crawler.ContractCrawler
	generalNotifier notifier.GeneralNotifier
}

func NewGRPCServer(
	dbManagers *database.DbManagers,
	config *Config,
	crawlerClient *contract_crawler.ContractCrawler,
	generalNotifier notifier.GeneralNotifier,
) *GRPCServer {
	return &GRPCServer{
		dbManagers:      dbManagers,
		config:          config,
		crawlerClient:   crawlerClient,
		generalNotifier: generalNotifier,
	}
}

func (s GRPCServer) Run() {
	log.Sugar.Info("Starting GRPC server")
	jwtManager := auth.NewJWTManager([]byte(s.config.JwtSecretKey), s.config.AccessTokenDuration, s.config.RefreshTokenDuration)
	interceptor := auth.NewAuthInterceptor(jwtManager, s.dbManagers.UserManager, auth.AccessibleRoles())

	authServer := auth.NewAuthServer(s.dbManagers.UserManager, jwtManager, s.config.TelegramToken, s.config.DiscordOAuth2Config, s.config.CannyPrivateKey)
	subsServer := subscription.NewSubscriptionsServer(s.dbManagers, s.crawlerClient)
	adminServer := admin.NewAdminServer(s.generalNotifier, s.dbManagers)
	trackerServer := tracker.NewTrackerServer(s.dbManagers)
	devServer := dev.NewDevServer(s.dbManagers, jwtManager)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	auth_service.RegisterAuthServiceServer(server, authServer)
	subscription_service.RegisterSubscriptionServiceServer(server, subsServer)
	admin_service.RegisterAdminServiceServer(server, adminServer)
	tracker_service.RegisterTrackerServiceServer(server, trackerServer)
	if os.Getenv("DEV") == "true" {
		dev_service.RegisterDevServiceServer(server, devServer)
	}

	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
