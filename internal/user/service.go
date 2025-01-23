package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/Hossara/linkin-chat/internal/user/domain"
	"github.com/Hossara/linkin-chat/internal/user/port"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"log"
)

var (
	ErrUserOnCreate      = errors.New("error on creating new user")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidPassword   = errors.New("password is invalid")
	ErrPasswordTooLong   = errors.New("password too long")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) RunMigrations() error {
	return s.repo.RunMigrations()
}

func (s *service) GetUserByUsernamePassword(ctx context.Context, username string, password string) (*domain.User, error) {
	user, err := s.repo.FindByUsernamePassword(ctx, username, password)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, ErrInvalidPassword):
			return nil, ErrUserNotFound
		default:
			return nil, errors.New(fmt.Sprintf("failed to authenticate user: %s", err))
		}
	}

	return user, nil
}

func (s *service) CreateUser(ctx context.Context, user *domain.User) (domain.UserID, error) {
	userID, err := s.repo.Insert(ctx, user)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, ErrUserAlreadyExists
		}

		log.Println("error on creating new user : ", err.Error())

		return 0, ErrUserOnCreate
	}

	return userID, nil
}

func (s *service) GetUserByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return user, nil
}
