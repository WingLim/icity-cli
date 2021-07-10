package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var FriendsCommand = &cli.Command{
	Name:  "friends",
	Usage: "Browser your friend diaries in iCity",
	Action: func(context *cli.Context) error {
		clear()
		getFriends()
		return nil
	},
}

func getFriends() {
	fmt.Println("getting diaries...")
	clear()

	user := getUser()
	diaries := user.GetFriends()

	browseDiaries(user, "World", diaries)
}
