package main

import (
	"context"
	"gasprice-oracle/internal/config"
	"gasprice-oracle/internal/service"
)

func main() {
	cfg, err := config.LoadConfig("APP")

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	_ = service.RunApplication(ctx, *cfg)
}
