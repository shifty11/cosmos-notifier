package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shifty11/cosmos-notifier/ent"
	"github.com/shifty11/cosmos-notifier/ent/user"
	"os"
	"time"
)

type Role string

var (
	Unauthenticated Role = "unauthenticated"
	User                 = Role(user.RoleUser.String())
	Admin                = Role(user.RoleAdmin.String())
)

type TokenType string

const (
	AccessToken  TokenType = "AccessToken"
	RefreshToken TokenType = "RefreshToken"
)

func AccessibleRoles() map[string][]Role {
	const path = "/cosmos_notifier_grpc"
	const authService = path + ".AuthService/"
	const subsService = path + ".SubscriptionService/"
	const adminService = path + ".AdminService/"
	const trackerService = path + ".TrackerService/"
	const devService = path + ".DevService/"

	roles := map[string][]Role{
		authService + "TelegramLogin":              {Unauthenticated, User, Admin},
		authService + "DiscordLogin":               {Unauthenticated, User, Admin},
		authService + "RefreshAccessToken":         {Unauthenticated, User, Admin},
		authService + "CannySSO":                   {User, Admin},
		subsService + "ListSubscriptions":          {User, Admin},
		subsService + "ToggleChainSubscription":    {User, Admin},
		subsService + "ToggleContractSubscription": {User, Admin},
		subsService + "AddDao":                     {Admin},
		subsService + "EnableChain":                {Admin},
		subsService + "DeleteDao":                  {Admin},
		trackerService + "ListTrackers":            {User, Admin},
		trackerService + "CreateTracker":           {User, Admin},
		trackerService + "UpdateTracker":           {User, Admin},
		trackerService + "DeleteTracker":           {User, Admin},
		trackerService + "IsAddressValid":          {User, Admin},
		trackerService + "TrackValidators":         {User, Admin},
		adminService + "BroadcastMessage":          {Admin},
		adminService + "GetStats":                  {Admin},
	}
	if os.Getenv("DEV") == "true" {
		roles[devService+"Login"] = []Role{Unauthenticated, User, Admin}
	}
	return roles
}

type Claims struct {
	jwt.RegisteredClaims
	UserId int64     `json:"user_id"`
	Type   user.Type `json:"type"`
	Role   Role      `json:"role,omitempty"`
}

type JWTManager struct {
	jwtSecretKey         []byte
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewJWTManager(jwtSecretKey []byte, accessTokenDuration time.Duration, refreshTokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		jwtSecretKey:         jwtSecretKey,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}

func (manager *JWTManager) GenerateToken(entUser *ent.User, tokenType TokenType) (string, error) {
	expirationTime := time.Now().Add(manager.accessTokenDuration)
	if tokenType == RefreshToken {
		expirationTime = time.Now().Add(manager.refreshTokenDuration)
	}

	claims := &Claims{
		UserId: entUser.UserID,
		Type:   entUser.Type,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		Role: Role(entUser.Role.String()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.jwtSecretKey)
}

func (manager *JWTManager) Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return manager.jwtSecretKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
