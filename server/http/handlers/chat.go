package handlers

import (
	"errors"
	"github.com/Hossara/linkin-chat/internal/chat"
	"github.com/Hossara/linkin-chat/pkg/jwt"
	"github.com/Hossara/linkin-chat/server/http/helpers"
	"github.com/Hossara/linkin-chat/server/http/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetAllChats(svcGetter helpers.ServiceGetter[*services.ChatService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		userClaims := jwt.GetUserClaims(c)

		response, err := svc.GetAllChats(c.UserContext(), userClaims.UserID)

		if err != nil {
			switch {
			case errors.Is(err, chat.ErrInvalidUserID):
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{})
			default:
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Internal server error",
					"message": err.Error(),
				})
			}
		}

		return c.JSON(response)
	}
}
