package service

import (
	"context"
	"gasprice-oracle/internal/adapter/storage"
	"gasprice-oracle/internal/app"
	"gasprice-oracle/internal/config"
	"gasprice-oracle/internal/logger"
	"golang.org/x/sync/errgroup"
)

func RunApplication(cfg config.Config) error {
	logger_, loggerCleanup := logger.New(cfg.LogLevel)
	defer loggerCleanup()

	logger_.Info("initializing application")

	ctx := context.Background()

	storage_ := storage.NewMemoryStorage()

	app_ := &app.Application{Context: ctx, Config: cfg, Logger: logger_, Storage: &storage_}

	logger_.Info("application initialized")

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		logger_.Info("running worker")
		return app_.RunWorker()
	})

	group.Go(func() error {
		logger_.Info("running http app")
		return app_.RunHTTPServer()
	})

	return group.Wait()
}
