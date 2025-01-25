package database

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	chatDomain "github.com/Hossara/linkin-chat/internal/chat/domain"
	"github.com/Hossara/linkin-chat/internal/chat/port"
	"github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/helpers"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/mapper"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/models"
	"github.com/Hossara/linkin-chat/pkg/utils"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var (
	ErrMaximumChatReached      = errors.New("maximum chat reached")
	ErrUserAlreadyExistsInChat = errors.New("user already exists in chat")
	ErrUserNotExistsInChat     = errors.New("user not exists in chat")
	ErrChatNotFound            = errors.New("chat not found")
)

const (
	codeLength = 8
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) port.ChatRepo {
	return &chatRepo{db}
}

func (r *chatRepo) Insert(ctx context.Context, room chatDomain.ChatRoom) error {
	count, err := r.Count(ctx, room.OwnerID)
	if err != nil {
		return fmt.Errorf("failed to count chats for user %d: %w", room.ID, err)
	}

	if count >= 5 {
		return ErrMaximumChatReached
	}

	chat := mapper.ToModelChat(&room)

	chat.Owner = models.User{ID: uint(room.OwnerID)}

	// Insert the chat into the database
	if err := r.db.WithContext(ctx).Create(&chat).Error; err != nil {
		return fmt.Errorf("failed to insert chat room: %w", err)
	}

	return nil
}

func generateRandomCode() (string, error) {
	b := make([]byte, codeLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random code: %w", err)
	}

	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}

	return string(b), nil
}

func (r *chatRepo) GenerateNewCode(ctx context.Context) (chatDomain.ChatRoomCode, error) {
	for {
		code, err := generateRandomCode()
		if err != nil {
			return "", err
		}

		// Check if the code is unique in the database
		var count int64
		err = r.db.WithContext(ctx).Model(&models.Chat{}).Where("code = ?", code).Count(&count).Error
		if err != nil {
			return "", fmt.Errorf("failed to verify code uniqueness: %w", err)
		}

		if count == 0 {
			return chatDomain.ChatRoomCode(code), nil
		}
	}
}

func (r *chatRepo) Count(ctx context.Context, userID domain.UserID) (int, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&models.Chat{}).
		Where("owner_id = ?", userID).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *chatRepo) Delete(ctx context.Context, code chatDomain.ChatRoomCode) error {
	var chat models.Chat

	// Fetch the chat by its code to ensure it exists
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&chat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrChatNotFound
		}
		return fmt.Errorf("failed to fetch chat: %w", err)
	}

	// Delete the chat
	err = r.db.WithContext(ctx).Delete(&chat).Error
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	return nil
}

func (r *chatRepo) FindAllByUserID(ctx context.Context, userID domain.UserID) ([]*chatDomain.ChatRoom, error) {
	var chats []models.Chat

	err := r.db.WithContext(ctx).
		Joins("LEFT JOIN chat_users ON chat_users.chat_id = chats.id").
		Where("chats.owner_id = ? OR chat_users.user_id = ?", userID, userID).
		Preload("Users").
		Preload("Owner").
		Find(&chats).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user chats: %w", err)
	}

	chatRooms := utils.Map(chats, func(chat models.Chat) *chatDomain.ChatRoom {
		return mapper.ToDomainChat(&chat)
	})

	return chatRooms, nil
}

func (r *chatRepo) InsertUserToChat(ctx context.Context, code chatDomain.ChatRoomCode, userID domain.UserID) error {
	var chat models.Chat

	// Fetch the chat by its code
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&chat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrChatNotFound
		}
		return fmt.Errorf("failed to fetch chat: %w", err)
	}

	var exists bool
	err = r.db.WithContext(ctx).
		Model(&chat).
		Select("COUNT(*) > 0").
		Where("user_id = ?", userID).
		Association("Users").
		Find(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if user exists in chat: %w", err)
	}
	if exists {
		return ErrUserAlreadyExistsInChat
	}

	// Add the user to the chat
	err = r.db.WithContext(ctx).Model(&chat).Association("Users").Append(&models.User{ID: uint(userID)})
	if err != nil {
		return fmt.Errorf("failed to add user to chat: %w", err)
	}

	return nil
}

func (r *chatRepo) DeleteUserFromChat(ctx context.Context, code chatDomain.ChatRoomCode, userID domain.UserID) error {
	var chat models.Chat

	err := r.db.WithContext(ctx).Where("code = ?", code).First(&chat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrChatNotFound
		}
		return fmt.Errorf("failed to fetch chat: %w", err)
	}

	// Check if the user is part of the chat
	var count int64
	err = r.db.WithContext(ctx).
		Table("chat_users").
		Where("chat_id = ? AND user_id = ?", chat.ID, userID).
		Count(&count).Error

	if err != nil {
		return fmt.Errorf("failed to check if user exists in chat: %w", err)
	}
	if count == 0 {
		return ErrUserNotExistsInChat
	}

	// Remove the user from the chat
	err = r.db.WithContext(ctx).Model(&chat).Association("Users").Delete(&models.User{ID: uint(userID)})
	if err != nil {
		return fmt.Errorf("failed to remove user from chat: %w", err)
	}

	return nil
}

func (r *chatRepo) RunMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.Chat]("chats", &models.Chat{}),
	)

	return migrator.Migrate()
}
