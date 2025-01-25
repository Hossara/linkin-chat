package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
}
