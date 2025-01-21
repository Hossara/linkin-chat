package config

type ServerConfig struct {
	DB     DBConfig `json:"db"`
	Server Server   `json:"server"`
	Nats   Nats     `json:"nats"`
}

type DBConfig struct {
	Host   string `json:"host"`
	Port   uint   `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type Server struct {
	Port                  uint   `json:"port"`
	Secret                string `json:"secret"`
	PasswordSecret        string `json:"password_secret"`
	OtpTtlMinutes         uint   `json:"otp_ttl_minutes"`
	MaxRequestsPerSecond  uint   `json:"maxRequestsPerSecond"`
	AuthExpirationMinutes uint   `json:"auth_expiration_minutes"`
	AuthRefreshMinutes    uint   `json:"auth_refresh_minutes"`
}

type Nats struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
