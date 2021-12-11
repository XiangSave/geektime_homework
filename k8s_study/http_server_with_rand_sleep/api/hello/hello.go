package hello

import (
	"fmt"
	"time"

	// "log"
	"math/rand"
	"net/http"
	"os"
	"sample_http_server/pkg/metrics"

	log "github.com/sirupsen/logrus"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	// get execution time
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()

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

	// random sleep 10mm ~ 2000mm
	delay := randInt(10, 2000)
	//delay := randInt(10, 200)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// add access status code header
	w.WriteHeader(statusCode)

	//response msg
	_, err := fmt.Fprintln(w, wMsg)
	if err != nil {
		log.Errorf("response msg %s error: %v", wMsg, err)
	}

	requestLogger := log.WithFields(log.Fields{"apiPath": r.URL.Path, "Status": statusCode, "RemoteAddress": rAddress})
	requestLogger.Infof("response: %s", wMsg)
	log.Println("aaa")
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
