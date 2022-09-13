package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/ent"
	"github.com/shifty11/dao-dao-notifier/ent/user"
	"github.com/shifty11/dao-dao-notifier/log"
	pb "github.com/shifty11/dao-dao-notifier/service_grpc/protobuf/go/auth_service"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"time"
)

//goland:noinspection GoNameStartsWithPackageName
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userManager         *database.UserManager
	jwtManager          *JWTManager
	telegramToken       string
	discordOAuth2Config *oauth2.Config
}

type DiscordIdentity struct {
	ID       json.Number `json:"id"`
	Username string      `json:"username"`
}

func NewAuthServer(
	userManager *database.UserManager,
	jwtManager *JWTManager,
	telegramToken string,
	discordOAuth2Config *oauth2.Config,
) pb.AuthServiceServer {
	return &AuthServer{
		userManager:         userManager,
		jwtManager:          jwtManager,
		telegramToken:       telegramToken,
		discordOAuth2Config: discordOAuth2Config,
	}
}

func (s *AuthServer) secretKey1() []byte {
	sha := sha256.New()
	sha.Write([]byte(s.telegramToken))
	secretKey := sha.Sum(nil)
	return secretKey
}

func (s *AuthServer) secretKey2() []byte {
	h1 := hmac.New(sha256.New, []byte("WebAppData"))
	h1.Write([]byte(s.telegramToken))
	secretKey := h1.Sum(nil)
	return secretKey
}

func (s *AuthServer) isValid(dataStr string, secretKey []byte, hash string) bool {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(dataStr))
	hh := h.Sum(nil)
	resultHash := hex.EncodeToString(hh)
	return resultHash == hash
}

func (s *AuthServer) login(entUser *ent.User, username string) (*pb.LoginResponse, error) {
	accessToken, err := s.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		log.Sugar.Errorf("Could not generate accessToken for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Error(codes.Internal, "login failed")
	}

	refreshToken, err := s.jwtManager.GenerateToken(entUser, RefreshToken)
	if err != nil {
		log.Sugar.Errorf("Could not generate refreshToken for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Error(codes.Internal, "login failed")
	}

	if username != entUser.Name {
		entUser, err = s.userManager.SetName(entUser, username)
		if err != nil {
			log.Sugar.Errorf("Could not update username of user %v (%v): %v", entUser.Name, entUser.ID, err)
			return nil, status.Error(codes.Unauthenticated, "login failed")
		}
	}

	res := &pb.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return res, nil
}

func (s *AuthServer) TelegramLogin(_ context.Context, req *pb.TelegramLoginRequest) (*pb.LoginResponse, error) {
	if !s.isValid(req.DataStr, s.secretKey1(), req.Hash) && !s.isValid(req.DataStr, s.secretKey2(), req.Hash) {
		return nil, status.Errorf(codes.Unauthenticated, "login failed")
	}

	if time.Now().Sub(time.Unix(req.AuthDate, 0)) > time.Hour {
		return nil, status.Errorf(codes.Unauthenticated, "login failed")
	}

	entUser, err := s.userManager.Get(req.UserId, user.TypeTelegram)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find user: %v", req.GetUserId())
	}

	return s.login(entUser, req.Username)
}

func (s *AuthServer) DiscordLogin(_ context.Context, req *pb.DiscordLoginRequest) (*pb.LoginResponse, error) {
	token, err := s.discordOAuth2Config.Exchange(context.Background(), req.GetCode())
	if err != nil {
		log.Sugar.Infof("Error exchanging code for token: %v", err)
		return nil, status.Error(codes.Unauthenticated, "login failed")
	}

	res, err := s.discordOAuth2Config.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
	if err != nil || res.StatusCode != 200 {
		log.Sugar.Infof("Error getting user (%v): %v", res.StatusCode, err)
		return nil, status.Error(codes.Unauthenticated, "login failed")
	}

	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Sugar.Infof("Error reading response body: %v", err)
		return nil, status.Error(codes.Unauthenticated, "login failed")
	}

	var identity DiscordIdentity
	err = json.Unmarshal(body, &identity)
	if err != nil {
		log.Sugar.Infof("Error unmarshalling response body: %v", err)
		return nil, status.Error(codes.Unauthenticated, "login failed")
	}

	id, err := identity.ID.Int64()
	if err != nil {
		log.Sugar.Infof("Error converting id to int64: %v", err)
		return nil, status.Error(codes.Unauthenticated, "login failed")
	}
	entUser, err := s.userManager.Get(id, user.TypeDiscord)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find user: %v", err)
	}

	return s.login(entUser, identity.Username)
}

func (s *AuthServer) RefreshAccessToken(_ context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	claims, err := s.jwtManager.Verify(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "refresh token invalid: %v", err)
	}

	entUser, err := s.userManager.Get(claims.UserId, claims.Type)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot find user: %v", err)
	}

	accessToken, err := s.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate accessToken: %v", err)
	}

	res := &pb.RefreshAccessTokenResponse{AccessToken: accessToken}
	return res, nil
}
