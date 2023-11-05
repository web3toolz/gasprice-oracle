package server

import (
	"gasprice-oracle/internal/config"
	"gasprice-oracle/internal/service"
	"github.com/urfave/cli/v2"
	"log"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Run http server",
		Action: func(cCtx *cli.Context) error {
			cfg, err := config.LoadConfigFromFile(cCtx.String("config"))

			if err != nil {
				log.Fatal("error to load config:", err)
			}

			return service.RunApplication(*cfg)
		},
	}
}
