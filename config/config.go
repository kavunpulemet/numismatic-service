package config

import (
	"github.com/spf13/viper"
)

const (
	configPath    = "config"
	configName    = "config"
	dbDatabaseKey = "db.database"
	dbPortKey     = "db.port"
	dbHostKey     = "db.host"
)

type Settings struct {
	Database string
	Port     string
	Host     string
}

func NewSettings() (Settings, error) {
	err := initConfig()
	if err != nil {
		return Settings{}, err
	}

	return Settings{
		Database: viper.GetString(dbDatabaseKey),
		Port:     viper.GetString(dbPortKey),
		Host:     viper.GetString(dbHostKey),
	}, nil
}

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
