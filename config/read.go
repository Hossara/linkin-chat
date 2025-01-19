package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig(configPath string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			err = viper.WriteConfigAs(configPath)

			if err != nil {
				return fmt.Errorf("error creating config file: %w", err)
			}
		} else {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	return nil
}
