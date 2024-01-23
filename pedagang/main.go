package main

import (
	"log"
	"net/http"
	"time"

	"github.com/nanwp/jajan-yuk/pedagang/config"
	"github.com/nanwp/jajan-yuk/pedagang/pkg/conn"
	"github.com/nanwp/jajan-yuk/pedagang/pkg/middleware"
)

func main() {
	cfg := config.Get()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	db := conn.InitPostgreSQL(&cfg)

	defer conn.DbClose(db)

	if cfg.LogMode {
		db = db.Debug()
	}

	router, _ := middleware.InitRouter(cfg, db)

	log.Printf("starting server at %v on port %v", time.Now().Format("2006-01-02 15:04:05"), cfg.HttpPort)
	err := http.ListenAndServe(":"+cfg.HttpPort, router)
	if err != nil {
		return
	}
}
