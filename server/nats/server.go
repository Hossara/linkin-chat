package nats

import (
	"fmt"
	"github.com/Hossara/linkin-chat/app"
	"github.com/Hossara/linkin-chat/config"
	"github.com/Hossara/linkin-chat/server/nats/handlers"
	"github.com/nats-io/nats.go"
	"log"
)

func Bootstrap(ac app.App, cfg config.Nats) error {
	addr := fmt.Sprintf("nats://%s:%s@%s:%d", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	nc, err := nats.Connect(addr)

	if err != nil {
		return fmt.Errorf("error connecting to NATS: %v", err)
	}

	log.Printf("Nats server started successfully: %s", addr)

	defer nc.Close()

	ac.ChatService()

	subscriptions(nc)

	select {}
}

func subscriptions(nc *nats.Conn) {
	// Subscribe to all chatrooms using wildcard
	_, err := nc.Subscribe("chatroom.send.*.*", handlers.SendMessageHandler)

	if err != nil {
		log.Fatalf("Error subscribing to chatroom.*: %v", err)
	}
}
