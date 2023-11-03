package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func NewMetricsHandler() (http.Handler, promauto.Factory) {
	registry := prometheus.NewRegistry()

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	factory := promauto.With(registry)

	return handler, factory
}
