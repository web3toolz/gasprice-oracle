package actions

import (
	"errors"
	"gasprice-oracle/internal/adapter/storage"
	"gasprice-oracle/internal/config"
	"gasprice-oracle/internal/ports/rest"
	"github.com/leonelquinteros/router"
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

	r := router.New("/")

	restHandler := rest.RHandler{
		Logger:   h.logger,
		Storage:  h.storage,
		Networks: h.networks,
	}

	r.Add(h.cfg.Path, &restHandler)

	server := http.Server{
		Addr: net.JoinHostPort(h.cfg.Host, h.cfg.Port),

		Handler: router.Build(r),

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
