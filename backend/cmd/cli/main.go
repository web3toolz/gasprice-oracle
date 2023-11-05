package main

import (
	"gasprice-oracle/cmd/cli/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:                 "gasprice-oracle",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			server.Cmd(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
