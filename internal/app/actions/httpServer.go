package actions

import (
	"errors"
	"gasprice-oracle/internal/adapter/storage"
	"gasprice-oracle/internal/config"
	"gasprice-oracle/internal/ports"
	"go.uber.org/zap"
	"net"
	"net/http"
)

type HTTPServerHandler struct {
	cfg     config.Config
	logger  *zap.Logger
	storage storage.IStorage
}

func NewHTTPServerHandler(cfg config.Config, logger *zap.Logger, storage storage.IStorage) *HTTPServerHandler {
	return &HTTPServerHandler{cfg: cfg, logger: logger, storage: storage}
}

func (h *HTTPServerHandler) Storage() storage.IStorage {
	return h.storage
}

func (h *HTTPServerHandler) Run() error {
	handler := ports.RootHandler{
		Logger:  h.logger,
		Storage: h.storage,
	}

	server := http.Server{
		Addr: net.JoinHostPort(h.cfg.Host, h.cfg.Port),

		Handler: handler.ServeHTTP(),

		ReadHeaderTimeout: h.cfg.HttpTimeout,
		ReadTimeout:       h.cfg.HttpTimeout,
		WriteTimeout:      h.cfg.HttpTimeout,
		IdleTimeout:       h.cfg.HttpTimeout * 30,
	}

	h.logger.Info("starting http server", zap.String("address", server.Addr))

	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
