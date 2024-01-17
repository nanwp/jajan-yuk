package main

import (
	"fmt"
	"github.com/nanwp/jajan-yuk/product/config"
)

func main() {

	cfg := config.Get()

	fmt.Printf("Db host, %s", cfg.DBHost)
}
