package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	icity "github.com/WingLim/icity-sdk"
	"github.com/manifoldco/promptui"
)

type bellSkipper struct{}

// Write implements an io.WriterCloser over os.Stderr, but it skips the terminal
// bell character.
func (bs *bellSkipper) Write(b []byte) (int, error) {
	const charBell = 7 // c.f. readline.CharBell
	if len(b) == 1 && b[0] == charBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

// Close implements an io.WriterCloser over os.Stderr.
func (bs *bellSkipper) Close() error {
	return os.Stderr.Close()
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func browseDiaries(user *icity.User, label string, diaries []icity.Diary) {
	var cursorPos int

browse:
	worldBrowse := promptui.Select{
		Label:     label,
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
