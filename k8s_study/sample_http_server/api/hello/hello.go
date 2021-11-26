package hello

import (
	"fmt"
	// "log"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	statusCode := 200
	rAddress := r.RemoteAddr

	for key := range r.Header {
		w.Header().Set(key, r.Header.Get(key))
	}

	w.Header().Set("VERSION", os.Getenv("VERSION"))

	hName := r.URL.Path[1:]
	wMsg := hName
	if hName == "" {
		statusCode = 500
		wMsg = "access deny"
	} else {
		wMsg = "hello " + hName
	}

	w.WriteHeader(statusCode)
	fmt.Fprintln(w, wMsg)
	// time.Sleep(10 * time.Second)
	// defer log.Println("defer")

	// log.Printf("apiPath:%s Status:%d RemoteAddress:%s", r.URL.Path, statusCode, rAddress)
	requestLogger := log.WithFields(log.Fields{"apiPath": r.URL.Path, "Status": statusCode, "RemoteAddress": rAddress})
	requestLogger.Infof("response: %s", wMsg)
}
