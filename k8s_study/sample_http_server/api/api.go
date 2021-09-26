package api

import (
	"net/http"
	"sample_http_server/api/health"
	"sample_http_server/api/hello"
)

func SetHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", health.Healthz)
	mux.HandleFunc("/", hello.Hello)
	return mux
}
