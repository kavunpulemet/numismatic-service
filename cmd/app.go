package main

import (
	"NumismaticClubApi/config"
	"NumismaticClubApi/pkg/api"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/repository"
	"NumismaticClubApi/pkg/service/coin"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

type App struct {
	ctx        utils.MyContext
	server     *api.Server
	repository *mongo.Database
	settings   config.Settings
}

func NewApp(ctx context.Context, logger *zap.SugaredLogger, settings config.Settings) *App {
	return &App{
		ctx:      utils.NewMyContext(ctx, logger),
		settings: settings,
	}
}

func (a *App) InitDatabase() error {
	mongoURL := fmt.Sprintf("mongodb://%s:%s", a.settings.Host, a.settings.Port)
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURL))
	if err != nil {
		a.ctx.Logger.Fatalf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		a.ctx.Logger.Fatalf("failed to ping MongoDB: %v", err)
	}

	a.repository = client.Database(a.settings.Database)
	return nil
}

func (a *App) InitService() {
	s := coin.NewCoinService(repository.NewRepository(a.repository))
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

	err = a.repository.Client().Disconnect(a.ctx.Ctx)
	if err != nil {
		a.ctx.Logger.Errorf("failed to disconnect from bd %v", err)
	}

	a.ctx.Logger.Info("server shut down successfully")
	return nil
}
