package app

import (
	"gorm.io/gorm"

	"github.com/Hossara/linkin-chat/config"

	userPort "github.com/Hossara/linkin-chat/internal/user/port"
)

type App interface {
	DB() *gorm.DB
	Config() config.ServerConfig
	UserService() userPort.Service
}
