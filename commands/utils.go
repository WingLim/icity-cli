package commands

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/xeonx/timeago"
)

var (
	funcMap   = promptui.FuncMap
	templates = &promptui.SelectTemplates{
		Label:    "{{ . }}: ",
		Active:   "> {{ .Nickname }} {{ .Date | timeAgo }}",
		Inactive: "  {{ .Nickname }} {{ .Date | timeAgo }}",
		Selected: "> {{ .Nickname }}",
		FuncMap:  funcMap,
		Details: `
--------- Diary ----------
{{ .Nickname }} {{ .Date | timeAgo }} {{ .Location }}
{{ .Title }}
{{ .Content | renderHTML }}
`,
	}
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

func init() {
	funcMap["timeAgo"] = TimeAgo
	funcMap["renderHTML"] = RenderHTML
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func TimeAgo(date time.Time) string {
	return timeago.Chinese.Format(date)
}

func RenderHTML(raw string) string {
	var result string
	result = strings.ReplaceAll(raw, "<br/>", "\n")
	return result
}
