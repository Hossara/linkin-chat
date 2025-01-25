package port

import (
	"context"
	"github.com/Hossara/linkin-chat/internal/user/domain"
)

type Service interface {
	GetUserByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (domain.UserID, error)
	GetUserByID(ctx context.Context, id domain.UserID) (*domain.User, error)
	RunMigrations() error
}
