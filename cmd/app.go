package main

import (
	"NumismaticClubApi/config"
	"NumismaticClubApi/pkg/api"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/repository"
	"NumismaticClubApi/pkg/service/coin"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type App struct {
	ctx        utils.MyContext
	server     *api.Server
	repository *sqlx.DB
	settings   config.Settings
}

func NewApp(ctx context.Context, logger *zap.SugaredLogger, settings config.Settings) *App {
	return &App{
		ctx:      utils.NewMyContext(ctx, logger),
		settings: settings,
	}
}

func (a *App) InitDatabase() error {
	var err error
	a.repository, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		a.settings.Host, a.settings.Port, a.settings.Username, a.settings.DBName, a.settings.Password, a.settings.SSLMode))
	if err != nil {
		return err
	}

	err = a.repository.Ping()
	if err != nil {
		return err
	}

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

func (a *App) Shutdown(ctx context.Context) error {
	err := a.server.Shutdown(ctx)
	if err != nil {
		a.ctx.Logger.Errorf("Failed to disconnect from server %v", err)
		return err
	}

	err = a.repository.Close()
	if err != nil {
		a.ctx.Logger.Errorf("failed to disconnect from bd %v", err)
	}

	a.ctx.Logger.Info("server shut down successfully")
	return nil
}
