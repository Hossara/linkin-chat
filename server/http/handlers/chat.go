package handlers

import (
	"errors"
	"github.com/Hossara/linkin-chat/pkg/jwt"
	"github.com/Hossara/linkin-chat/server/http/helpers"
	"github.com/Hossara/linkin-chat/server/http/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"net/http"
)

func GetAllChats(svcGetter helpers.ServiceGetter[*services.ChatService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		userClaims := jwt.GetUserClaims(c)

		response, err := svc.GetAllChats(c.UserContext(), userClaims.UserID)

		if err != nil {
			switch {
			case errors.Is(err, services.ErrInvalidUserID):
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

func CreateNewChat(svcGetter helpers.ServiceGetter[*services.ChatService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		userClaims := jwt.GetUserClaims(c)
		title := utils.CopyString(c.Params("title"))

		response, err := svc.CreateNewChat(c.UserContext(), userClaims.UserID, title)

		if err != nil {
			switch {
			case errors.Is(err, services.ErrInvalidUserID):
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{})
			case errors.Is(err, services.ErrMaximumChatReached):
				return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
					"error":   "Maximum chat reached",
					"message": "You can only own 5 chats!",
				})
			case errors.Is(err, services.ErrInvalidChatInfo):
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"error":   "Invalid chat title",
					"message": "",
				})
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
