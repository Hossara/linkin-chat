package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/Hossara/linkin-chat/server/http/helpers"
	"github.com/Hossara/linkin-chat/server/http/services"
	"github.com/Hossara/linkin-chat/server/http/types"
)

func Login(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		body := new(types.LoginRequest)

		if err := helpers.ParseRequestBody(c, body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		response, err := svc.Login(c.UserContext(), *body)

		if err != nil {
			fmt.Printf("Error while logging in %v\n", err)

			switch {
			case errors.Is(err, services.ErrUserNotFound):
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": "username or password incorrect",
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

func Register(svcGetter helpers.ServiceGetter[*services.AccountService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		body := new(types.RegisterRequest)

		if err := helpers.ParseRequestBody(c, body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		response, err := svc.Register(c.UserContext(), *body)

		if err != nil {
			switch {
			case errors.Is(err, services.ErrUserOnCreate):
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Error while creating new user",
					"message": err.Error(),
				})
			case errors.Is(err, services.ErrUserAlreadyExists):
				return c.Status(http.StatusConflict).JSON(fiber.Map{
					"error": "User already exists",
				})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Internal server error",
					"message": err.Error(),
				})
			}
		}

		return c.JSON(response)
	}
}
