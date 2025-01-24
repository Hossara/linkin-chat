package database

import (
	"context"
	chatDomain "github.com/Hossara/linkin-chat/internal/chat/domain"
	"github.com/Hossara/linkin-chat/internal/chat/port"
	"github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/helpers"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type messageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) port.MessageRepo {
	return &messageRepo{db}
}

func (r *messageRepo) Insert(ctx context.Context, code chatDomain.ChatRoomCode, userID domain.UserID, content string) error {
	//TODO implement me
	panic("implement me")
}

func (r *messageRepo) FindAllByChatCode(ctx context.Context, code chatDomain.ChatRoomCode) ([]chatDomain.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (r *messageRepo) RunMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.Message]("messages", &models.Message{}),
	)

	return migrator.Migrate()
}
