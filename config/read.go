package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig(configPath string) (*ServerConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// Unmarshal into the Config struct
	var config ServerConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %v", err)
	}

	return &config, nil
}

func MustReadConfig(path string) ServerConfig {
	config, err := ReadConfig(path)
	if err != nil {
		panic(err)
	}
	return *config
}
