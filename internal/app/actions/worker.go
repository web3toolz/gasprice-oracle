package actions

import (
	"context"
	"gasprice-oracle/internal/adapter/storage"
	"gasprice-oracle/internal/config"
	"gasprice-oracle/pkg/ethclient"
	"gasprice-oracle/pkg/gasprice"
	"go.uber.org/zap"
)

type WorkerHandler struct {
	ctx     context.Context
	cfg     config.Config
	logger  *zap.Logger
	storage storage.IStorage
}

func NewWorkerHandler(ctx context.Context, cfg config.Config, logger *zap.Logger, storage storage.IStorage) *WorkerHandler {
	return &WorkerHandler{ctx: ctx, cfg: cfg, logger: logger, storage: storage}
}

func (w *WorkerHandler) RunIteration() error {
	client, err := ethclient.New(w.ctx, w.cfg)

	if err != nil {
		w.logger.Error("error while initializing eth client", zap.Error(err))
		return err
	}

	block, err := client.GetPendingBlockData(w.ctx)

	if err != nil {
		w.logger.Error("failed to get block data", zap.Error(err))
		return err
	}

	data, err := block.GetTransactionsGasPrices()

	if err != nil {
		w.logger.Error("failed to get gasprice data", zap.Error(err))
		return err
	}

	distribution, err := gasprice.DistributionFomSlice(data)

	if err != nil {
		w.logger.Error("failed to get gasprice distribution", zap.Error(err))
		return err
	}

	w.logger.Debug("saving distribution to storage")
	w.logger.Debug("new distributions", zap.Any("values", distribution))
	_ = w.storage.Set(storage.BlockNumber, block.GetBlockNumberAsInt())
	_ = w.storage.Set(storage.P40, distribution.P40)
	_ = w.storage.Set(storage.P60, distribution.P60)
	_ = w.storage.Set(storage.P75, distribution.P75)
	_ = w.storage.Set(storage.P95, distribution.P95)

	return nil
}
