package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var initCommand = &cli.Command{
	Name:  "init",
	Usage: "init iCity directory",
	Action: func(context *cli.Context) error {
		return initCLI()
	},
}

func checkExist() bool {
	_, err := os.Stat(iCityPath)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("iCity has initialized")
			return true
		}
		return false
	}
	return true
}

func initCLI() error {
	if !checkExist() {
		err := os.Mkdir(iCityPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
