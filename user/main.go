package main

import (
	"github.com/nanwp/jajan-yuk/user/config"
	"github.com/nanwp/jajan-yuk/user/pkg/conn"
	"github.com/nanwp/jajan-yuk/user/pkg/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.Get()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	log.Println("starting server ... at", time.Now().Format("2006-01-02 15:04:05"))

	db := conn.InitPostgreSQL(&cfg)

	defer conn.DbClose(db)

	if cfg.LogMode {
		db = db.Debug()
	}

	router, _ := middleware.InitRouter(cfg, db)

	http.ListenAndServe(":"+cfg.HttpPort, &router)
}
