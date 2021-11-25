package main

import (
	"os"
	"sample_http_server/api"
	"sample_http_server/pkg/server"

	log "github.com/sirupsen/logrus"
	// "log"
)

func init() {
	err := setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}

func main() {
	s := server.New("0.0.0.0:8080", api.SetHandlers())
	log.Info("http starting")
	err := s.Run()
	if err != nil {
		log.Println(err)
	}
}

func setupLogger() error {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
