package port

import (
	"context"

	chatDomain "github.com/Hossara/linkin-chat/internal/chat/domain"
	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
)

type Service interface {
	CreateChatRoom(ctx context.Context, room chatDomain.ChatRoom) (chatDomain.ChatRoomCode, error)
	JoinChatRoom(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID) (chatDomain.ChatRoomCode, error)
	ExitChatRoom(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID) (chatDomain.ChatRoomCode, error)
	DeleteChatRoom(ctx context.Context, code chatDomain.ChatRoomCode) error
	GetUserChatRooms(ctx context.Context, userID userDomain.UserID) ([]*chatDomain.ChatRoom, error)

	GetChatRoomMessages(ctx context.Context, code chatDomain.ChatRoomCode) []*chatDomain.Message
	NewMessage(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID, content string) error

	RunChatMigrations() error
	RunMessageMigrations() error
}
