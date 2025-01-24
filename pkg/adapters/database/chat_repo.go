package database

import (
	"context"
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
	panic("implement me")
	/*var count int64
	if err := r.db.Table("chat_users").Where("user_id = ?", chatID, userID).Count(&count).Error; err != nil {
		return fmt.Errorf("failed to check if user exists in chat: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("user already exists in the chat")
	}

	// Step 3: Add the user to the chat
	user := User{ID: userID} // You only need the ID to associate it
	if err := r.db.Model(&Chat{ID: chatID}).Association("Users").Append(&user); err != nil {
		return fmt.Errorf("failed to add user to chat: %w", err)
	}

	return nil*/
}

func (r *chatRepo) Delete(ctx context.Context, code chatDomain.ChatRoomCode) error {
	//TODO implement me
	panic("implement me")
}

func (r *chatRepo) FindAllByUserID(ctx context.Context, userID domain.UserID) ([]*chatDomain.ChatRoom, error) {
	var chats []models.Chat
	// Query to find all chats associated with the given userID
	err := r.db.WithContext(ctx).Joins("JOIN chat_users ON chat_users.chat_id = chats.id").
		Where("chat_users.user_id = ?", userID).Find(&chats).Error
	if err != nil {
		return nil, err
	}

	// Transform database models to domain models
	chatRooms := utils.Map(chats, func(t models.Chat) *chatDomain.ChatRoom {
		return mapper.ToDomainChat(&t)
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
