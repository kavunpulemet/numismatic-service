package config

import (
	"github.com/spf13/viper"
	"os"
)

const (
	configPath    = "config"
	configName    = "config"
	dbHostKey     = "db.host"
	dbPortKey     = "db.port"
	dbUsernameKey = "db.username"
	dbPasswordEnv = "DB_PASSWORD"
	dbNameKey     = "db.dbname"
	dbSSLModeKey  = "db.sslmode"
)

type Settings struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewSettings() (Settings, error) {
	err := initConfig()
	if err != nil {
		return Settings{}, err
	}

	return Settings{
		Host:     viper.GetString(dbHostKey),
		Port:     viper.GetString(dbPortKey),
		Username: viper.GetString(dbUsernameKey),
		Password: os.Getenv(dbPasswordEnv),
		DBName:   viper.GetString(dbNameKey),
		SSLMode:  viper.GetString(dbSSLModeKey),
	}, nil
}

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
