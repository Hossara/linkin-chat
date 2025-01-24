package chat

import (
	"context"
	"fmt"
	"github.com/Hossara/linkin-chat/internal/user"

	"github.com/Hossara/linkin-chat/internal/chat/port"

	chatDomain "github.com/Hossara/linkin-chat/internal/chat/domain"
	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
)

var ()

type service struct {
	chatRepo    port.ChatRepo
	messageRepo port.MessageRepo
}

func NewService(chatRepo port.ChatRepo, messageRepo port.MessageRepo) port.Service {
	return &service{
		chatRepo:    chatRepo,
		messageRepo: messageRepo,
	}
}

func (s *service) GetChatRoomMessages(ctx context.Context, code chatDomain.ChatRoomCode) []*chatDomain.Message {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetUserChatRooms(ctx context.Context, userID userDomain.UserID) ([]*chatDomain.ChatRoom, error) {
	if userID == 0 {
		return nil, user.ErrInvalidUserID
	}

	chatRooms, err := s.chatRepo.FindAllByUserID(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to query chat rooms for user %d: %w", userID, err)
	}

	return chatRooms, nil
}

func (s *service) CreateChatRoom(ctx context.Context, room chatDomain.ChatRoom) (chatDomain.ChatRoomCode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) JoinChatRoom(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID) (chatDomain.ChatRoomCode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteChatRoom(ctx context.Context, code chatDomain.ChatRoomCode) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) ExitChatRoom(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID) (chatDomain.ChatRoomCode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) NewMessage(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID, content string) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) RunChatMigrations() error {
	return s.chatRepo.RunMigrations()
}

func (s *service) RunMessageMigrations() error {
	return s.messageRepo.RunMigrations()
}
