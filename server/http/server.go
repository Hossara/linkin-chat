package http

import (
	"encoding/json"
	"fmt"
	"github.com/Hossara/linkin-chat/app"
	"github.com/Hossara/linkin-chat/config"
	"github.com/gofiber/fiber/v2"
)

func Bootstrap(ac app.App, cfg config.Server) error {
	fiberApp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	return fiberApp.Listen(fmt.Sprintf(":%d", cfg.Port))
}
