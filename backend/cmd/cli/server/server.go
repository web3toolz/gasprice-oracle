package server

import (
	"gasprice-oracle/internal/config"
	"gasprice-oracle/internal/service"
	"github.com/urfave/cli/v2"
	"log"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "Config file path",
		Value:   "config.yaml",
	},
}

func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Run http server",
		Flags: flags,
		Action: func(cCtx *cli.Context) error {
			cfg, err := config.LoadConfigFromFile(cCtx.String("config"))

			if err != nil {
				log.Fatal("error to load config:", err)
			}

			return service.RunApplication(*cfg)
		},
	}
}
