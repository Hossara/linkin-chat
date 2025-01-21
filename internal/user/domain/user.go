package domain

import (
	"github.com/google/uuid"
	"time"
)

type UserID uint

type User struct {
	ID        UserID
	UUID      uuid.UUID
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	DeletedAt time.Time
}
