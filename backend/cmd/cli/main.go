package main

import (
	"gasprice-oracle/cmd/cli/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "Config file path",
		Value:   "config.yaml",
	},
}

func main() {
	app := &cli.App{
		Name:                 "gasprice-oracle",
		EnableBashCompletion: true,
		Flags:                flags,
		Commands: []*cli.Command{
			server.Cmd(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
