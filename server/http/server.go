package http

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"

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

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}

	publicKey := &privateKey.PublicKey
	viper.Set("public_key", publicKey)
	viper.Set("private_key", privateKey)

	authMiddleware := middlewares.Authorization(publicKey)
	reteLimiter := middlewares.RateLimiter()

	authGroup := fiberApp.Group("/auth", reteLimiter)
	chatGroup := fiberApp.Group("/chat", reteLimiter)

	accountSvcGetter := services.AccountServiceGetter(ac, cfg, privateKey)
	chatSvcGetter := services.ChatServiceGetter(ac)

	authGroup.Post("/login", handlers.Login(accountSvcGetter))
	authGroup.Post("/register", handlers.Register(accountSvcGetter))

	chatGroup.Get("/", authMiddleware, handlers.GetAllChats(chatSvcGetter))

	//api.Post("/nats", NatsAuth(accountSvcGetter))

	return fiberApp.Listen(fmt.Sprintf(":%d", cfg.Port))
}
