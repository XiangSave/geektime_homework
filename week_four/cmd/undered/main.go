package main

import (
	"log"
	"week_four/api"
	"week_four/internal/pkg/mysql"
	"week_four/internal/server"
)

func main() {
	dbInfo := &mysql.DBInfo{
		DBType:       "mysql",
		Host:         "",
		UserName:     "",
		Password:     "",
		Charset:      "utf8mb4",
		DatabaseName: "",
	}
	lunchDB := mysql.NewDBModel(dbInfo)
	err := lunchDB.Connect()
	if err != nil {
		log.Panic(err)
	}

	s := server.New("0.0.0.0:8080", api.UnorderHandler(lunchDB))
	err = s.Run()
	if err != nil {
		log.Println(err)
	}
}
