package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Code    string `gorm:"unique;not null;size:100"`
	Title   string `gorm:"type:text;not null"`
	Users   []User `gorm:"many2many:chat_users;"`
	OwnerID uint   `gorm:"not null;index"`
	Owner   User   `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Message struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	ChatID   uint   `gorm:"not null;index"`
	SenderID uint   `gorm:"not null;index"`
	Content  string `gorm:"type:text;not null"`

	Chat   Chat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ChatID"`   // Relation to Chat
	Sender User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:SenderID"` // Relation to User
}
