package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	configPath       = "config"
	configName       = "config"
	mongoDatabaseKey = "mongo.database"
	mongoPortKey     = "mongo.port"
	mongoHostKey     = "mongo.host"
	redisPortKey     = "redis.port"
	redisHostKey     = "redis.host"
	redisPasswordKey = "redis.password"
	redisDBKey       = "redis.db"
)

type Settings struct {
	Mongo MongoSettings
	Redis RedisSettings
}

type MongoSettings struct {
	Database string
	MongoURL string
}

type RedisSettings struct {
	Address  string
	Password string
	DB       int
}

func NewSettings() (Settings, error) {
	err := initConfig()
	if err != nil {
		return Settings{}, err
	}

	return Settings{
		Mongo: MongoSettings{
			Database: viper.GetString(mongoDatabaseKey),
			MongoURL: fmt.Sprintf("mongodb://%s:%s", viper.GetString(mongoHostKey), viper.GetString(mongoPortKey)),
		},
		Redis: RedisSettings{
			Address:  fmt.Sprintf("%s:%s", viper.GetString(redisHostKey), viper.GetString(redisPortKey)),
			Password: viper.GetString(redisPasswordKey),
			DB:       viper.GetInt(redisDBKey),
		},
	}, nil
}

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
