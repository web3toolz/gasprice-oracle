package app

import (
	"context"
	"gasprice-oracle/internal/adapter/storage"
	actions2 "gasprice-oracle/internal/app/actions"
	"gasprice-oracle/internal/config"
	"go.uber.org/zap"
)

type Application struct {
	Context context.Context
	Config  config.Config
	Logger  *zap.Logger
	Storage storage.IStorage
}

func (app *Application) RunHTTPServer() error {
	var networks []string

	for _, network := range app.Config.Networks {
		networks = append(networks, network.NetworkName)
	}

	http := actions2.NewHTTPServerHandler(app.Config.Server, app.Logger, app.Storage, networks)

	return http.Run()
}

func (app *Application) RunWorker() error {
	worker := actions2.NewWorkerHandler(app.Context, app.Config.Worker, app.Logger, app.Storage, app.Config.Networks)

	return worker.Run()
}
