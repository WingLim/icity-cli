package commands

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"log"
)

var WorldCommand = &cli.Command{
	Name:  "world",
	Usage: "Browser diaries in iCity world",
	Action: func(context *cli.Context) error {
		clear()
		getWorld()
		return nil
	},
}

func getWorld() {
	user := getUser()
	diaries := user.GetWorld()

	var cursorPos int

browse:
	worldBrowse := promptui.Select{
		Label:     "World",
		Items:     diaries,
		CursorPos: cursorPos,
		Templates: templates,
		Stdout:    &bellSkipper{},
	}

	index, _, err := worldBrowse.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	diaryID := diaries[index].ID
	if diaryID != "" {
		commentInput := promptui.Prompt{
			Label: "Comment",
		}
		comment, err := commentInput.Run()
		if err != nil {
			log.Fatal(err)
			return
		}
		resp := user.NewComment(diaryID, comment)
		if resp.Success {
			fmt.Println("comment success")
		} else {
			fmt.Println("comment failed")
		}
		clear()
		cursorPos = index
		goto browse
	}
}
