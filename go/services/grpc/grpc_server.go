package grpc

import (
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
	"github.com/shifty11/dao-dao-notifier/notifier"
	"github.com/shifty11/dao-dao-notifier/services/contract_crawler"
	"github.com/shifty11/dao-dao-notifier/services/grpc/admin"
	"github.com/shifty11/dao-dao-notifier/services/grpc/auth"
	"github.com/shifty11/dao-dao-notifier/services/grpc/protobuf/go/admin_service"
	"github.com/shifty11/dao-dao-notifier/services/grpc/protobuf/go/auth_service"
	"github.com/shifty11/dao-dao-notifier/services/grpc/protobuf/go/subscription_service"
	"github.com/shifty11/dao-dao-notifier/services/grpc/subscription"
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
	dbManagers      *database.DbManagers
	config          *Config
	crawlerClient   *contract_crawler.ContractCrawler
	generalNotifier *notifier.GeneralNotifier
}

func NewGRPCServer(dbManagers *database.DbManagers, config *Config, crawlerClient *contract_crawler.ContractCrawler, generalNotifier *notifier.GeneralNotifier) *GRPCServer {
	return &GRPCServer{dbManagers: dbManagers, config: config, crawlerClient: crawlerClient, generalNotifier: generalNotifier}
}

func (s GRPCServer) Run() {
	log.Sugar.Info("Starting GRPC server")
	jwtManager := auth.NewJWTManager([]byte(s.config.JwtSecretKey), s.config.AccessTokenDuration, s.config.RefreshTokenDuration)
	interceptor := auth.NewAuthInterceptor(jwtManager, s.dbManagers.UserManager, auth.AccessibleRoles())

	authServer := auth.NewAuthServer(s.dbManagers.UserManager, jwtManager, s.config.TelegramToken, s.config.DiscordOAuth2Config)
	subsServer := subscription.NewSubscriptionsServer(s.dbManagers, s.crawlerClient)
	adminServer := admin.NewAdminServer(s.generalNotifier)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	auth_service.RegisterAuthServiceServer(server, authServer)
	subscription_service.RegisterSubscriptionServiceServer(server, subsServer)
	admin_service.RegisterAdminServiceServer(server, adminServer)

	lis, err := net.Listen("tcp", s.config.Port)
	if err != nil {
		log.Sugar.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Sugar.Fatalf("failed to serve grpc: %v", err)
	}
}
