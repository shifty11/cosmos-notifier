package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"github.com/shifty11/cosmos-notifier/log"
	pb "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/go/auth_service"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"time"
)

//goland:noinspection GoNameStartsWithPackageName
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userManager         *database.UserManager
	jwtManager          *JWTManager
	telegramToken       string
	discordOAuth2Config *oauth2.Config
	cannyPrivateKey     string
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
	cannyPrivateKey string,
) pb.AuthServiceServer {
	return &AuthServer{
		userManager:         userManager,
		jwtManager:          jwtManager,
		telegramToken:       telegramToken,
		discordOAuth2Config: discordOAuth2Config,
		cannyPrivateKey:     cannyPrivateKey,
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

var (
	ErrorLoginFailed  = status.Error(codes.Unauthenticated, "login failed")
	ErrorUserNotFound = status.Error(codes.NotFound, "user not found")
	ErrorLoginExpired = status.Error(codes.Unauthenticated, "login expired")
	ErrorInternal     = status.Error(codes.Internal, "internal error")
)

func (s *AuthServer) login(entUser *ent.User, username string) (*pb.LoginResponse, error) {
	accessToken, err := s.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		log.Sugar.Errorf("Could not generate accessToken for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, ErrorLoginFailed
	}

	refreshToken, err := s.jwtManager.GenerateToken(entUser, RefreshToken)
	if err != nil {
		log.Sugar.Errorf("Could not generate refreshToken for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, ErrorInternal
	}

	if username != entUser.Name {
		entUser, err = s.userManager.SetName(entUser, username)
		if err != nil {
			log.Sugar.Errorf("Could not update username of user %v (%v): %v", entUser.Name, entUser.ID, err)
			return nil, ErrorInternal
		}
	}

	res := &pb.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}
	return res, nil
}

func (s *AuthServer) TelegramLogin(_ context.Context, req *pb.TelegramLoginRequest) (*pb.LoginResponse, error) {
	if !s.isValid(req.DataStr, s.secretKey1(), req.Hash) && !s.isValid(req.DataStr, s.secretKey2(), req.Hash) {
		return nil, ErrorLoginFailed
	}

	if time.Now().Sub(time.Unix(req.AuthDate, 0)) > time.Hour {
		return nil, ErrorLoginExpired
	}

	entUser, err := s.userManager.Get(req.UserId, user.TypeTelegram)
	if err != nil {
		return nil, ErrorUserNotFound
	}

	return s.login(entUser, req.Username)
}

func (s *AuthServer) DiscordLogin(_ context.Context, req *pb.DiscordLoginRequest) (*pb.LoginResponse, error) {
	token, err := s.discordOAuth2Config.Exchange(context.Background(), req.GetCode())
	if err != nil {
		log.Sugar.Infof("Error exchanging code for token: %v", err)
		return nil, ErrorLoginFailed
	}

	res, err := s.discordOAuth2Config.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
	if err != nil || res.StatusCode != 200 {
		log.Sugar.Infof("Error getting user (%v): %v", res.StatusCode, err)
		return nil, ErrorLoginFailed
	}

	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Sugar.Infof("Error reading response body: %v", err)
		return nil, ErrorInternal
	}

	var identity DiscordIdentity
	err = json.Unmarshal(body, &identity)
	if err != nil {
		log.Sugar.Infof("Error unmarshalling response body: %v", err)
		return nil, ErrorInternal
	}

	id, err := identity.ID.Int64()
	if err != nil {
		log.Sugar.Infof("Error converting id to int64: %v", err)
		return nil, ErrorInternal
	}
	entUser, err := s.userManager.Get(id, user.TypeDiscord)
	if err != nil {
		return nil, ErrorUserNotFound
	}

	return s.login(entUser, identity.Username)
}

func (s *AuthServer) RefreshAccessToken(_ context.Context, req *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	claims, err := s.jwtManager.Verify(req.RefreshToken)
	if err != nil {
		return nil, ErrorLoginFailed
	}

	entUser, err := s.userManager.Get(claims.UserId, claims.Type)
	if err != nil {
		log.Sugar.Errorf("Could not find user %v (%v): %v", claims.UserId, claims.UserId, err)
		return nil, ErrorUserNotFound
	}

	accessToken, err := s.jwtManager.GenerateToken(entUser, AccessToken)
	if err != nil {
		log.Sugar.Errorf("Could not generate accessToken for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, ErrorInternal
	}

	res := &pb.RefreshAccessTokenResponse{AccessToken: accessToken}
	return res, nil
}

func (s *AuthServer) CannySSO(ctx context.Context, _ *emptypb.Empty) (*pb.CannySSOResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.NotFound, "invalid user")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": fmt.Sprintf("%v-%v@cosmos-notifier.com", entUser.Type, entUser.UserID),
		"id":    entUser.UserID,
		"name":  entUser.Name,
		//"email": "raphael.thurnherr1990@gmail.com",	# uncomment to set up SSO in canny
		//"id":    "6111ae303ab6ab4927a638f8",			# uncomment to set up SSO in canny
	})
	signedToken, err := token.SignedString([]byte(s.cannyPrivateKey))
	if err != nil {
		log.Sugar.Errorf("Could not sign token: %v", err)
		return nil, ErrorInternal
	}

	return &pb.CannySSOResponse{AccessToken: signedToken}, nil
}
