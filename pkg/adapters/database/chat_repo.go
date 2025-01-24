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

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) port.ChatRepo {
	return &chatRepo{db}
}

func (r *chatRepo) Insert(ctx context.Context, room chatDomain.ChatRoom) error {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (r *chatRepo) FindAllByUserID(ctx context.Context, userID domain.UserID) ([]*chatDomain.ChatRoom, error) {
	var chats []models.Chat

	err := r.db.WithContext(ctx).
		Where("owner_id = ?", userID).
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

func (r *chatRepo) InsertUserToChat(ctx context.Context, code chatDomain.ChatRoomCode, userID domain.UserID) {
	//TODO implement me
	panic("implement me")
}

func (r *chatRepo) DeleteUserFromChat(ctx context.Context, code chatDomain.ChatRoomCode, userID domain.UserID) {
	//TODO implement me
	panic("implement me")
}

func (r *chatRepo) RunMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.Chat]("chats", &models.Chat{}),
	)

	return migrator.Migrate()
}
