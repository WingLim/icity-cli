package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"github.com/xeonx/timeago"
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
	user := getUser()
	diaries := user.GetWorld()

	funcMap := promptui.FuncMap
	funcMap["timeAgo"] = TimeAgo

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}: ",
		Active:   "> {{ .Nickname }} {{ .Date | timeAgo }}",
		Inactive: "  {{ .Nickname }} {{ .Date | timeAgo }}",
		Selected: "> {{ .Nickname }}",
		FuncMap:  funcMap,
		Details: `
--------- Diary ----------
{{ .Nickname }} {{ .Date | timeAgo }} {{ .Location }}
{{ .Title }}
{{ .Content }}
`,
	}

	worldBrowse := promptui.Select{
		Label:     "World",
		Items:     diaries,
		Templates: templates,
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
	}
}

func TimeAgo(date time.Time) string {
	return timeago.Chinese.Format(date)
}
