package app

import (
	"context"
	"gasprice-oracle/internal/adapter/storage"
	"gasprice-oracle/internal/app/actions"
	"gasprice-oracle/internal/config"
	"go.uber.org/zap"
	"time"
)

type Application struct {
	Context context.Context
	Config  config.Config
	Logger  *zap.Logger
	Storage storage.IStorage
}

func (app *Application) RunHTTPServer() error {
	http := actions.NewHTTPServerHandler(app.Config, app.Logger, app.Storage)

	return http.Run()
}

func (app *Application) RunWorker() error {
	worker := actions.NewWorkerHandler(app.Context, app.Config, app.Logger, app.Storage)

	ticker := time.NewTicker(app.Config.WorkerInterval)
	defer ticker.Stop()

	go func() {
		_ = worker.RunIteration()
	}()

	for {
		select {
		case <-ticker.C:
			app.Logger.Debug("new worker iteration")

			err := worker.RunIteration()

			if err != nil {
				app.Logger.Error("error while running worker iteration", zap.Error(err))
				return err
			}

		case <-app.Context.Done():
			app.Logger.Info("stopping worker")
			return nil
		}
	}

}
