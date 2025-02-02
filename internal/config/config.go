package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	AppPort string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	config.Database.Host = viper.GetString("DB_HOST")
	config.Database.Port = viper.GetString("DB_PORT")
	config.Database.User = viper.GetString("DB_USER")
	config.Database.Password = viper.GetString("DB_PASSWORD")
	config.Database.Name = viper.GetString("DB_NAME")

	config.AppPort = viper.GetString("APP_PORT")

	return &config, nil
}
