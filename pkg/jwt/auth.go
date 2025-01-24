package jwt

import (
	"crypto/ecdsa"
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

func GetUserClaims(ctx *fiber.Ctx) *UserClaims {
	if u := ctx.Locals(UserClaimKey); u != nil {
		userClaims, ok := u.(*jwt2.Token).Claims.(*UserClaims)

		if ok {
			return userClaims
		}
	}
	return nil
}
