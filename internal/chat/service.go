package chat

import (
	"context"
	"errors"
	"fmt"
	chatDomain "github.com/Hossara/linkin-chat/internal/chat/domain"
	"github.com/Hossara/linkin-chat/internal/chat/port"
	"github.com/Hossara/linkin-chat/internal/user"
	userDomain "github.com/Hossara/linkin-chat/internal/user/domain"
	chatRepo "github.com/Hossara/linkin-chat/pkg/adapters/database"
)

var (
	ErrInvalidUserID           = user.ErrInvalidUserID
	ErrMaximumChatReached      = chatRepo.ErrMaximumChatReached
	ErrChatNotFound            = chatRepo.ErrChatNotFound
	ErrUserNotExistsInChat     = chatRepo.ErrUserNotExistsInChat
	ErrUserAlreadyExistsInChat = chatRepo.ErrUserAlreadyExistsInChat
	ErrInvalidChatInfo         = errors.New("chat room must have a valid code and title")
)

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

func (s *service) GetUserChatRooms(ctx context.Context, userID userDomain.UserID) ([]*chatDomain.ChatRoom, error) {
	if userID == 0 {
		return nil, ErrInvalidUserID
	}

	chatRooms, err := s.chatRepo.FindAllByUserID(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("failed to query chat rooms for user %d: %w", userID, err)
	}

	return chatRooms, nil
}

func (s *service) CreateChatRoom(ctx context.Context, room chatDomain.ChatRoom) (chatDomain.ChatRoomCode, error) {
	if room.OwnerID == 0 {
		return "", fmt.Errorf("invalid user ID")
	}

	if room.Title == "" {
		return "", ErrInvalidChatInfo
	}

	code, err := s.chatRepo.GenerateNewCode(ctx)

	if err != nil {
		return "", fmt.Errorf("failed to generate code: %w", err)
	}

	room.Code = code

	err = s.chatRepo.Insert(ctx, room)

	if err != nil {
		if errors.Is(err, ErrMaximumChatReached) {
			return "", ErrMaximumChatReached
		}

		return "", fmt.Errorf("failed to create chat room: %w", err)
	}

	return room.Code, nil
}

func (s *service) JoinChatRoom(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID) (chatDomain.ChatRoomCode, error) {
	if code == "" {
		return "", ErrInvalidChatInfo
	}

	if userID == 0 {
		return "", ErrInvalidUserID
	}

	err := s.chatRepo.InsertUserToChat(ctx, code, userID)
	if err != nil {
		if errors.Is(err, ErrChatNotFound) || errors.Is(err, ErrUserAlreadyExistsInChat) {
			return "", err
		}

		return "", fmt.Errorf("failed to join chat room: %w", err)
	}

	return code, nil
}

func (s *service) DeleteChatRoom(ctx context.Context, code chatDomain.ChatRoomCode) error {
	if code == "" {
		return ErrInvalidChatInfo
	}

	err := s.chatRepo.Delete(ctx, code)

	if err != nil {
		if errors.Is(err, ErrChatNotFound) {
			return err
		}

		return fmt.Errorf("failed to delete chat room: %w", err)
	}

	return nil
}

func (s *service) ExitChatRoom(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID) (chatDomain.ChatRoomCode, error) {
	if code == "" {
		return "", ErrInvalidChatInfo
	}

	if userID == 0 {
		return "", ErrInvalidUserID
	}

	err := s.chatRepo.DeleteUserFromChat(ctx, code, userID)

	if err != nil {
		if errors.Is(err, ErrChatNotFound) || errors.Is(err, ErrUserNotExistsInChat) {
			return "", err
		}

		return "", fmt.Errorf("failed to exit chat room: %w", err)
	}

	return code, nil
}

func (s *service) NewMessage(ctx context.Context, code chatDomain.ChatRoomCode, userID userDomain.UserID, content string) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetChatRoomMessages(ctx context.Context, code chatDomain.ChatRoomCode) []*chatDomain.Message {
	//TODO implement me
	panic("implement me")
}

func (s *service) RunChatMigrations() error {
	return s.chatRepo.RunMigrations()
}

func (s *service) RunMessageMigrations() error {
	return s.messageRepo.RunMigrations()
}
