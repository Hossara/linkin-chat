package middlewares

import (
	"crypto/ecdsa"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/Hossara/linkin-chat/pkg/jwt"
)

// Authorization middleware using jwtware
func Authorization(secret *ecdsa.PublicKey) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: "ES256",
			Key:    secret,
		},
		Claims:      &jwt.UserClaims{},
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		SuccessHandler: func(c *fiber.Ctx) error {
			userClaims := jwt.GetUserClaims(c)

			if userClaims == nil {
				return fiber.ErrUnauthorized
			}

			return c.Next()
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token", err.Error())
		},
	})
}
