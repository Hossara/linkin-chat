package models

import (
	"gorm.io/gorm"
)

type UserID uint

type User struct {
	gorm.Model
	ID        UserID `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
}
