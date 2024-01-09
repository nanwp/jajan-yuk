package conn

import (
	"fmt"
	"github.com/nanwp/jajan-yuk/user/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func InitPostgreSQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	} else {
		log.Println("Success connect to database")
	}

	rdb, err := db.DB()

	if err != nil {
		log.Fatalf("error at %v", err.Error())
		panic(err)
	}

	rdb.SetMaxIdleConns(cfg.DBMaxIdleConns)
	rdb.SetMaxOpenConns(cfg.DBMaxOpenConns)
	rdb.SetConnMaxLifetime(time.Duration(int(time.Minute) * cfg.DBConnMaxLifeTime))

	return db
}

func DbClose(db *gorm.DB) {
	rdb, err := db.DB()
	if err != nil {
		log.Fatalf("error at %v", err.Error())
		panic(err)
	}
	_ = rdb.Close()
}
