package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "icity"
	app.Version = "0.2.0"
	app.UseShortOptionHandling = true

	app.Commands = []*cli.Command{
		initCommand,
		loginCommand,
		newCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
