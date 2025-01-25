package handlers

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"strings"
)

func SendMessageHandler(msg *nats.Msg) {
	subject := strings.Split(msg.Subject, ".")
	chatroomID := subject[1]
	username := subject[2]

	println(chatroomID)
	println(username)

	// Log the chatroom ID and message payload
	fmt.Printf("Message received in chatroom %s: %s\n", chatroomID, string(msg.Data))
}
