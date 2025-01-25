package domain

import (
	"time"

	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
)

type MessageID uint

type Message struct {
	ID         MessageID
	ChatRoomID ChatRoomID
	SenderID   userDomain.UserID
	Content    string
	CreatedAt  time.Time
	DeletedAt  time.Time
}
