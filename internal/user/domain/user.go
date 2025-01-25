package domain

import (
	"time"
)

type UserID uint

type User struct {
	ID        UserID
	Username  string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	DeletedAt time.Time
}
