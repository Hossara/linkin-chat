package http

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"

	"github.com/Hossara/linkin-chat/app"
	"github.com/Hossara/linkin-chat/config"
	"github.com/Hossara/linkin-chat/server/http/handlers"
	"github.com/Hossara/linkin-chat/server/http/services"

	middlewares "github.com/Hossara/linkin-chat/server/http/middleware"
)

func Bootstrap(ac app.App, cfg config.Server) error {
	fiberApp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api := fiberApp.Group("/auth")

	accountSvcGetter := services.AccountServiceGetter(ac, cfg)

	api.Post("/login", middlewares.RateLimiter(), handlers.Login(accountSvcGetter))
	api.Post("/register", middlewares.RateLimiter(), handlers.Register(accountSvcGetter))

	//api.Post("/nats", NatsAuth(accountSvcGetter))

	return fiberApp.Listen(fmt.Sprintf(":%d", cfg.Port))
}
