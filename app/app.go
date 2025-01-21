package app

import (
	"github.com/Hossara/linkin-chat/config"
	"gorm.io/gorm"
)

type app struct {
	cfg config.ServerConfig
	db  *gorm.DB
}

func (a *app) DB() *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (a *app) Config() config.ServerConfig {
	return a.cfg
}

func NewApp(cfg config.ServerConfig) (App, error) {
	a := &app{cfg: cfg}

	return a, nil
}

func MustNewApp(cfg config.ServerConfig) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}
