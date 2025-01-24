package config

type ServerConfig struct {
	DB     DBConfig    `mapstructure:"db"`
	Server Server      `mapstructure:"server"`
	Redis  RedisConfig `mapstructure:"redis"`
	Nats   Nats        `mapstructure:"nats"`
}

type DBConfig struct {
	Host   string `mapstructure:"host"`
	Port   uint   `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	Name   string `mapstructure:"name"`
	Schema string `mapstructure:"schema"`
}

type Server struct {
	Port                  uint   `mapstructure:"port"`
	PasswordSecret        string `mapstructure:"password_secret"`
	MaxRequestsPerSecond  uint   `mapstructure:"maxRequestsPerSecond"`
	AuthExpirationMinutes uint   `mapstructure:"auth_expiration_minutes"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
}
type Nats struct {
	Host string `mapstructure:"host"`
	Port uint   `mapstructure:"port"`
}
