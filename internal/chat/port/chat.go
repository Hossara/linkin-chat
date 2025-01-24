package port

import (
	"context"
	chatDomain "github.com/Hossara/linkin-chat/internal/chat/domain"
	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
)

type ChatRepo interface {
	Insert(ctx context.Context, room chatDomain.ChatRoom) error
	Count(ctx context.Context, userID userDomain.UserID) (int, error)
	Delete(ctx context.Context, code chatDomain.ChatRoomCode) error

	GenerateNewCode(ctx context.Context) (chatDomain.ChatRoomCode, error)

	FindAllByUserID(ctx context.Context, userID userDomain.UserID) ([]*chatDomain.ChatRoom, error)
	InsertUserToChat(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID)
	DeleteUserFromChat(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID)

	RunMigrations() error
}

type MessageRepo interface {
	Insert(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID, content string) error
	FindAllByChatCode(ctx context.Context, code chatDomain.ChatRoomCode) ([]chatDomain.Message, error)

	RunMigrations() error
}
