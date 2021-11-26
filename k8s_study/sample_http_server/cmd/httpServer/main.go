package main

import (
	"os"
	"sample_http_server/api"
	"sample_http_server/pkg/server"
	"strconv"

	"sample_http_server/global"
	"sample_http_server/pkg/conf"

	log "github.com/sirupsen/logrus"
	// "log"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger(global.HttpServer.ServiceSetting.LogPath, log.InfoLevel)
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}

func main() {
	log.Info(global.HttpServer.ServiceSetting)

	listenAddress := global.HttpServer.ServiceSetting.BindIP + ":" + strconv.Itoa(global.HttpServer.ServiceSetting.Port)

	s := server.New(listenAddress, api.SetHandlers())
	log.Info("http starting")
	err := s.Run()
	if err != nil {
		// log.Println(err)
		log.Fatal(err)
	}
}

func setupLogger(logPath string, logLevel log.Level) error {
	log.SetFormatter(&log.JSONFormatter{})

	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log.SetOutput(file)
	log.SetLevel(logLevel)

	return nil
}

func setupSetting() error {
	confDirpath := "./configs"
	setting, err := conf.NewSetting(confDirpath)
	if err != nil {
		return err
	}
	err = setting.ReadHttpServer("HttpServer", &global.HttpServer)
	if err != nil {
		return err
	}

	return nil
}
