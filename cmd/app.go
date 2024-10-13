package main

import (
	"NumismaticClubApi/config"
	"NumismaticClubApi/models"
	"NumismaticClubApi/pkg/api"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/database"
	"NumismaticClubApi/pkg/database/cache"
	"NumismaticClubApi/pkg/service/coin"
	"context"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
	"time"
)

const (
	cacheKey = "coin:%s"
	ttl      = 10 * time.Minute
)

type App struct {
	ctx      utils.MyContext
	server   *api.Server
	mongo    *mongo.Database
	redis    *redis.Client
	settings config.Settings
}

func NewApp(ctx context.Context, logger *zap.SugaredLogger, settings config.Settings) *App {
	return &App{
		ctx:      utils.NewMyContext(ctx, logger),
		settings: settings,
	}
}

func (a *App) InitDatabase() error {
	client, err := mongo.Connect(options.Client().ApplyURI(a.settings.Mongo.MongoURL))
	if err != nil {
		a.ctx.Logger.Fatalf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(a.ctx.Ctx, nil)
	if err != nil {
		a.ctx.Logger.Fatalf("failed to ping MongoDB: %v", err)
	}

	a.mongo = client.Database(a.settings.Mongo.Database)

	a.redis = redis.NewClient(&redis.Options{
		Addr:     a.settings.Redis.Address,
		Password: a.settings.Redis.Password,
		DB:       a.settings.Redis.DB,
	})

	_, err = a.redis.Ping(a.ctx.Ctx).Result()
	if err != nil {
		a.ctx.Logger.Fatalf("failed to ping Redis: %v", err)
	}

	return nil
}

func (a *App) InitService() {
	s := coin.NewCoinService(database.NewMongoRepository(a.mongo), cache.NewRedisCache[string, models.Coin](a.redis, cacheKey, ttl))

	a.server = api.NewServer(a.ctx)
	a.server.HandleCoins(a.ctx, s)
}

func (a *App) Run() error {
	go func() {
		if err := a.server.Run(); err != nil {
			a.ctx.Logger.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	a.ctx.Logger.Info("run server")
	return nil
}

func (a *App) Shutdown() error {
	err := a.server.Shutdown(a.ctx.Ctx)
	if err != nil {
		a.ctx.Logger.Errorf("Failed to disconnect from server %v", err)
		return err
	}

	err = a.mongo.Client().Disconnect(a.ctx.Ctx)
	if err != nil {
		a.ctx.Logger.Errorf("failed to disconnect from bd %v", err)
	}

	a.ctx.Logger.Info("server shut down successfully")
	return nil
}
