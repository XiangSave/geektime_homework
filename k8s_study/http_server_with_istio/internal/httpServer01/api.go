package httpServer01

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sample_http_server/api/health"
	"sample_http_server/api/hello"
	"sample_http_server/pkg/metrics"
)

func SetHandlers() *http.ServeMux {

	metrics.Register()

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello.Hello)
	mux.HandleFunc("/healthz", health.Healthz)
	mux.Handle("/metrics", promhttp.Handler())

	return mux
}
