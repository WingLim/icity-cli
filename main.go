package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "icity"
	app.Version = "0.0.1"
	app.UseShortOptionHandling = true

	app.Commands = []*cli.Command{
		initCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
