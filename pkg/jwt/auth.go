package jwt

import (
	"errors"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/Hossara/linkin-chat/internal/user/domain"
)

const UserClaimKey = "User-Claims"

func CreateToken(secret []byte, claims *UserClaims) (string, error) {
	return jwt2.NewWithClaims(jwt2.SigningMethodHS512, claims).SignedString(secret)
}

func GenerateUserClaims(user *domain.User, exp time.Time) *UserClaims {
	return &UserClaims{
		RegisteredClaims: jwt2.RegisteredClaims{
			ExpiresAt: &jwt2.NumericDate{
				Time: exp,
			},
		},
		Username: user.Username,
		UserID:   uint(user.ID),
	}
}

func ParseToken(tokenString string, secret []byte) (*UserClaims, error) {
	token, err := jwt2.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt2.Token) (interface{}, error) {
		return secret, nil
	})

	if token == nil {
		return nil, errors.New("invalid token (nil)")
	}

	var claim *UserClaims
	if token.Claims != nil {
		cc, ok := token.Claims.(*UserClaims)
		if ok {
			claim = cc
		}
	}

	if err != nil {
		return claim, err
	}

	if !token.Valid {
		return claim, errors.New("token is not valid")
	}

	return claim, nil
}
