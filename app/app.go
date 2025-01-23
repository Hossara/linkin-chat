package app

import (
	"fmt"
	"github.com/Hossara/linkin-chat/internal/user"
	"github.com/Hossara/linkin-chat/pkg/adapters/database"
	"github.com/Hossara/linkin-chat/pkg/cache"
	"gorm.io/gorm"
	"log"

	userPort "github.com/Hossara/linkin-chat/internal/user/port"
	redisAdapter "github.com/babyhando/order-service/pkg/adapters/cache"

	"github.com/Hossara/linkin-chat/config"
	"github.com/Hossara/linkin-chat/pkg/postgres"
)

type app struct {
	cfg config.ServerConfig
	db  *gorm.DB

	redisProvider cache.Provider

	userService userPort.Service
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.ServerConfig {
	return a.cfg
}

func (a *app) Cache() cache.Provider {
	return a.redisProvider
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

func (a *app) UserService() userPort.Service {
	if a.userService == nil {
		a.userService = user.NewService(database.NewUserRepo(a.db))

		if err := a.userService.RunMigrations(); err != nil {
			log.Fatalf("failed to run migrations for user service: %v", err)
		}

		return a.userService
	}

	return a.userService
}

func (a *app) setRedis() {
	a.redisProvider = redisAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))
}

func NewApp(cfg config.ServerConfig) (App, error) {
	a := &app{cfg: cfg}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.setRedis()

	return a, nil
}

func MustNewApp(cfg config.ServerConfig) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}
