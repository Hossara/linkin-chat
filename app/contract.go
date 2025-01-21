package app

import (
	"github.com/Hossara/linkin-chat/config"
	"gorm.io/gorm"
)

type App interface {
	DB() *gorm.DB
	Config() config.ServerConfig
}
