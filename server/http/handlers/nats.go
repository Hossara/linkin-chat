package handlers

import (
	"errors"
	"fmt"
	"github.com/Hossara/linkin-chat/pkg/jwt"
	"github.com/Hossara/linkin-chat/pkg/utils"
	"github.com/Hossara/linkin-chat/server/http/helpers"
	"github.com/Hossara/linkin-chat/server/http/services"
	"github.com/Hossara/linkin-chat/server/http/types"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"net/http"
)

func HandleNatsAuth(
	accountSvcGetter helpers.ServiceGetter[*services.AccountService],
	chatSvcGetter helpers.ServiceGetter[*services.ChatService],
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//accountSvc := accountSvcGetter(c.UserContext())
		chatSvc := chatSvcGetter(c.UserContext())
		body := new(types.NatsAuthRequest)

		if err := helpers.ParseRequestBody(c, body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		claims, err := jwt.ParseToken(body.Token, viper.GetString("public_key"))

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
		}

		chats, err := chatSvc.GetAllChats(c.UserContext(), claims.UserID)

		if err != nil {
			switch {
			case errors.Is(err, services.ErrInvalidUserID):
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error":   "Unauthorized",
					"message": err.Error(),
				})
			default:
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Internal server error",
					"message": err.Error(),
				})
			}
		}

		chatsString := func(receive bool) []string {
			return utils.Map(chats.Chats, func(t types.ResponseChatRoom) string {
				return fmt.Sprintf("chatroom.%s.%s", utils.IfThenElse(
					receive, "receive", "send",
				), t.Code)
			})
		}

		authResponse := types.NatsAuthResponse{
			Publish:   chatsString(false),
			Subscribe: chatsString(true),
		}

		return c.JSON(authResponse)
	}
}
