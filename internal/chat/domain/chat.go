package domain

import (
	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
	"time"
)

type ChatRoomID uint
type ChatRoomCode string
type ChatRoomUserRole int

const (
	ChatRoomUserAmin ChatRoomUserRole = iota
	ChatRoomUserNormal
	ChatRoomUserBan
)

type ChatRoomUser struct {
	UserID userDomain.UserID
	Role   ChatRoomUserRole
}

type ChatRoom struct {
	ID        ChatRoomID
	Code      ChatRoomCode
	Users     []userDomain.User
	CreatedAt time.Time
	DeletedAt time.Time
}
