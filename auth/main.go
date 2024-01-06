package main

import (
	"github.com/nanwp/jajan-yuk/auth/config"
	"github.com/nanwp/jajan-yuk/auth/pkg/conn"
	"github.com/nanwp/jajan-yuk/auth/pkg/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	log.Println("starting server ... at", time.Now().Format("2006-01-02 15:04:05"))
	cfg := config.Get()

	db := conn.InitPostgreSQL(&cfg)

	defer conn.DbClose(db)

	if cfg.LogMode {
		db = db.Debug()
	}

	router, _ := middleware.InitRouter(cfg, db)

	http.ListenAndServe(":"+cfg.HttpPort, &router)
}
