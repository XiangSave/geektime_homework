package health

import (
	"fmt"
	// "log"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	rAddress := r.RemoteAddr

	for key := range r.Header {
		w.Header().Set(key, r.Header.Get(key))
	}
	w.Header().Set("VERSION", os.Getenv("VERSION"))

	statusCode := 200
	w.WriteHeader(statusCode)
	w.Header().Set("Accept", r.Header.Get("Accept"))
	fmt.Fprintln(w, "health")

	requestLogger := log.WithFields(log.Fields{"apiPath": r.URL.Path, "Status": statusCode, "RemoteAddress": rAddress})
	requestLogger.Info("health")
}
