package actions

import (
	"context"
	storagelib "gasprice-oracle/internal/adapter/storage"
	"gasprice-oracle/internal/config"
	"gasprice-oracle/pkg/ethclient"
	"gasprice-oracle/pkg/gasprice"
	"go.uber.org/zap"
	"time"
)

type result struct {
	network      string
	url          string
	distribution gasprice.Distribution
	updatedAt    int64
}

type WorkerHandler struct {
	ctx      context.Context
	cfg      config.WorkerConfig
	logger   *zap.Logger
	storage  storagelib.IStorage
	networks []config.NetworkConfig
}

func NewWorkerHandler(ctx context.Context, cfg config.WorkerConfig, logger *zap.Logger, storage storagelib.IStorage, networks []config.NetworkConfig) *WorkerHandler {
	return &WorkerHandler{ctx: ctx, cfg: cfg, logger: logger, storage: storage, networks: networks}
}

func (w *WorkerHandler) Run() error {
	ticker := time.NewTicker(w.cfg.Interval)
	defer ticker.Stop()

	resultCh := make(chan result)
	defer func() {
		close(resultCh)
	}()

	go func() {
		w.runIteration(resultCh)
	}()

	for {
		select {
		case <-ticker.C:
			w.logger.Debug("new worker iteration")
			w.runIteration(resultCh)

		case newResult := <-resultCh:
			err := w.saveResultToStorage(newResult)

			if err != nil {
				w.logger.Error("failed to save result to storage", zap.Error(err))
			}

		case <-w.ctx.Done():
			w.logger.Info("stopping worker")
			return nil
		}
	}
}

func (w *WorkerHandler) runIteration(resultCh chan result) {
	for _, network_ := range w.networks {
		go func(network string, url string) {
			newResult, err := w.runIterationForNetwork(network, url)

			if err != nil {
				w.logger.Error("failed to run iteration for network", zap.String("network", network))
			} else {
				resultCh <- *newResult
			}
		}(network_.NetworkName, network_.Url)
	}
}

func (w *WorkerHandler) runIterationForNetwork(network string, url string) (*result, error) {
	w.logger.Debug("trying to get data for network", zap.String("network", network))

	client, err := ethclient.New(url)

	if err != nil {
		w.logger.Error("error while initializing eth client", zap.Error(err), zap.String("network", network))
		return nil, err
	}

	block, err := client.GetPendingBlockData(w.ctx)

	if err != nil {
		w.logger.Error("failed to get block data", zap.Error(err), zap.String("network", network))
		return nil, err
	}

	data, err := block.GetTransactionsGasPrices()

	if err != nil {
		w.logger.Error("failed to get gasprice data", zap.Error(err), zap.String("network", network))
		return nil, err
	}

	distribution, err := gasprice.DistributionFomSlice(data)

	if err != nil {
		w.logger.Error("failed to get gasprice distribution", zap.Error(err), zap.String("network", network))
		return nil, err
	}

	updatedAt := time.Now().Unix()
	return &result{network: network, url: url, distribution: *distribution, updatedAt: updatedAt}, nil

}

func (w *WorkerHandler) saveResultToStorage(resultData result) error {
	if resultData.distribution.IsEmpty() {
		w.logger.Debug("distribution is empty, skipping saving to storage", zap.String("network", resultData.network))
		return nil
	}

	w.logger.Debug("saving data to storage", zap.String("network", resultData.network), zap.Any("data", resultData.distribution))
	_ = w.storage.Set(resultData.network, storagelib.UpdatedAt, resultData.updatedAt)
	_ = w.storage.Set(resultData.network, storagelib.P40, resultData.distribution.P40)
	_ = w.storage.Set(resultData.network, storagelib.P60, resultData.distribution.P60)
	_ = w.storage.Set(resultData.network, storagelib.P75, resultData.distribution.P75)
	_ = w.storage.Set(resultData.network, storagelib.P95, resultData.distribution.P95)
	return nil
}
