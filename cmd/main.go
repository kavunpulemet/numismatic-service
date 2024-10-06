package main

import (
	"NumismaticClubApi/config"
	_ "NumismaticClubApi/docs"
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// @title Numismatic Club API
// @version 1.0
// @description This is a sample API for a Numismatic Club.
// @host localhost:81
// @BasePath /api/v1
func main() {
	prdLogger, _ := zap.NewProduction()
	defer prdLogger.Sync()
	logger := prdLogger.Sugar()

	fmt.Println(logger.Level())

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()

	settings, err := config.NewSettings()
	if err != nil {
		logger.Fatalf("failed to read settings: %s", err.Error())
	}

	app := NewApp(ctx, logger, settings)
	if err := app.InitDatabase(); err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}

	app.InitService()

	if err = app.Run(); err != nil {
		logger.Errorf(err.Error())
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan

	if err = app.Shutdown(ctx); err != nil {
		logger.Errorf(err.Error())
		return
	}
}
