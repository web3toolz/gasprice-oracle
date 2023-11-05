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
	cfg      config.ServerConfig
	logger   *zap.Logger
	storage  storage.IStorage
	networks []string
}

func NewHTTPServerHandler(cfg config.ServerConfig, logger *zap.Logger, storage storage.IStorage, networks []string) *HTTPServerHandler {
	return &HTTPServerHandler{cfg: cfg, logger: logger, storage: storage, networks: networks}
}

func (h *HTTPServerHandler) Storage() storage.IStorage {
	return h.storage
}

func (h *HTTPServerHandler) Run() error {
	handler := ports.RootHandler{
		Logger:   h.logger,
		Storage:  h.storage,
		Networks: h.networks,
	}

	server := http.Server{
		Addr: net.JoinHostPort(h.cfg.Host, h.cfg.Port),

		Handler: handler.ServeHTTP(),

		ReadHeaderTimeout: h.cfg.Timeout,
		ReadTimeout:       h.cfg.Timeout,
		WriteTimeout:      h.cfg.Timeout,
		IdleTimeout:       h.cfg.Timeout * 10,
	}

	h.logger.Info("starting http server", zap.String("address", server.Addr))

	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}