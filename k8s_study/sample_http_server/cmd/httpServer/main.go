package main

import (
	"log"
	"sample_http_server/api"
	"sample_http_server/pkg/server"
)

func main() {
	s := server.New("0.0.0.0:80", api.SetHandlers())
	err := s.Run()
	if err != nil {
		log.Println(err)
	}
}
