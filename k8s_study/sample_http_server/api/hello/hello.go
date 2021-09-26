package hello

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	log.Printf("apiPath:%s Status:%d RemoteAddress:%s", r.URL.Path, statusCode, rAddress)
}
