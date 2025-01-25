package jwt

import (
	"crypto/ecdsa"
	"errors"
	"github.com/gofiber/fiber/v2"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/Hossara/linkin-chat/internal/user/domain"
)

const UserClaimKey = "user"

func CreateToken(secret *ecdsa.PrivateKey, claims *UserClaims) (string, error) {
	return jwt2.NewWithClaims(jwt2.SigningMethodES256, claims).SignedString(secret)
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

func ParseToken(tokenString string, secret string) (*UserClaims, error) {
	token, err := jwt2.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt2.Token) (interface{}, error) {
		return secret, nil
	})

	if token == nil {
		return nil, errors.New("invalid token (null)")
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

func GetUserClaims(ctx *fiber.Ctx) *UserClaims {
	if u := ctx.Locals(UserClaimKey); u != nil {
		userClaims, ok := u.(*jwt2.Token).Claims.(*UserClaims)

		if ok {
			return userClaims
		}
	}
	return nil
}
