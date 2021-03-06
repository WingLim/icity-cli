package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var WorldCommand = &cli.Command{
	Name:  "world",
	Usage: "Browser diaries in iCity world",
	Action: func(context *cli.Context) error {
		getWorld()
		return nil
	},
}

func getWorld() {
	fmt.Println("getting diaries...")
	clear()

	user := getUser()
	diaries := user.GetWorld()

	browseDiaries(user, "World", diaries)
}
