package database

import (
	"context"
	"errors"
	"fmt"
	userService "github.com/Hossara/linkin-chat/internal/user"
	"github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/internal/user/port"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/helpers"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/mapper"
	"github.com/Hossara/linkin-chat/pkg/adapters/database/models"
	"github.com/Hossara/linkin-chat/pkg/hash"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{db}
}

func (r *userRepo) FindByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error) {
	var user models.User

	// Retrieve the user by username
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Validate the plain password against the hashed password
	bcryptHasher := hash.NewBcryptHasher()
	if !bcryptHasher.Validate(user.Password, password) {
		return nil, userService.ErrInvalidPassword
	}

	// Map the user to domain and return
	return mapper.ToDomainUser(&user), nil
}

func (r *userRepo) RunMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.User]("users", &models.User{}),
	)

	return migrator.Migrate()
}

func (r *userRepo) Insert(ctx context.Context, user *domain.User) (domain.UserID, error) {
	newU := mapper.ToModelUser(user)

	return domain.UserID(newU.ID), r.db.WithContext(ctx).Create(newU).Error
}

func (r *userRepo) FindByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	return &user, nil
}

func (r *userRepo) Update(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
