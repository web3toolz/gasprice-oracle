package main

import (
	"gasprice-oracle/internal/config"
	"gasprice-oracle/internal/service"
)

func main() {
	cfg, err := config.LoadConfig("APP")

	if err != nil {
		panic(err)
	}

	_ = service.RunApplication(*cfg)
}
