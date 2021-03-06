package commands

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"strings"

	icity "github.com/WingLim/icity-sdk"
	"github.com/urfave/cli/v2"
)

var NewCommand = &cli.Command{
	Name:  "new",
	Usage: "Post a new diary",
	Action: func(context *cli.Context) error {
		ok := newDiary(context)
		if !ok {
			return errors.New("post new diary failed")
		}
		fmt.Println("post successful")
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "privacy",
			Aliases:     []string{"p"},
			Usage:       "Diary privacy, options are:\n \tpublic | friend | private",
			DefaultText: "public",
		},
	},
}

func convertPrivacy(privacy string) icity.DiaryPrivacy {
	privacy = strings.ToLower(privacy)
	switch privacy {
	case "friend":
		return icity.OnlyFriend
	case "private":
		return icity.Private
	default:
		return icity.Public
	}
}

func newDiary(context *cli.Context) bool {
	user := getUser()
	p := context.String("privacy")
	privacy := convertPrivacy(p)

	titleInput := promptui.Prompt{
		Label: "Title",
	}
	title, err := titleInput.Run()
	if err != nil {
		log.Fatal(err)
		return false
	}

	contentInput := promptui.Prompt{
		Label: "Content",
	}
	content, err := contentInput.Run()
	if err != nil {
		log.Fatal(err)
		return false
	}

	resp := user.NewDiary(title, content, privacy)
	return resp.Success
}
