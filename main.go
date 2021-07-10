package main

import (
	"log"
	"os"

	"github.com/WingLim/icity-cli/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "icity"
	app.Usage = "An iCity command line interface client."
	app.Version = "0.4.0"
	app.UseShortOptionHandling = true

	app.Commands = []*cli.Command{
		commands.InitCommand,
		commands.LoginCommand,
		commands.NewCommand,
		commands.WorldCommand,
		commands.FriendsCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
