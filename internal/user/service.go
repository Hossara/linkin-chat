package user

import (
	"errors"
	"github.com/Hossara/linkin-chat/internal/user/port"
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
