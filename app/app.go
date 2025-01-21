package app

import (
	"github.com/Hossara/linkin-chat/config"
	"github.com/Hossara/linkin-chat/pkg/postgres"
)

type app struct {
	cfg config.ServerConfig
	db  *gorm.DB
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.ServerConfig {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Pass,
		Name:   a.cfg.DB.Name,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func NewApp(cfg config.ServerConfig) (App, error) {
	a := &app{cfg: cfg}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	return a, nil
}

func MustNewApp(cfg config.ServerConfig) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}
