package app

import (
	"github.com/Hossara/linkin-chat/pkg/cache"
	"gorm.io/gorm"

	"github.com/Hossara/linkin-chat/config"

	chatPort "github.com/Hossara/linkin-chat/internal/chat/port"
	userPort "github.com/Hossara/linkin-chat/internal/user/port"
)

type App interface {
	DB() *gorm.DB
	Config() config.ServerConfig
	Cache() cache.Provider
	UserService() userPort.Service
	ChatService() chatPort.Service
}
