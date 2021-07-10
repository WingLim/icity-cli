package commands

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
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
{{ with .Title }}
{{ . }}
{{ end }}
{{ with .Photos }}
{{ . | renderPhotos }}
{{ end }}
{{ .Content | renderHTML }}
`,
	}
)

func init() {
	funcMap["timeAgo"] = TimeAgo
	funcMap["renderHTML"] = RenderHTML
	funcMap["renderPhotos"] = RenderPhotos
}

func TimeAgo(date time.Time) string {
	return timeago.Chinese.Format(date)
}

func RenderPhotos(photos []string) string {
	var datas string
	for _, photo := range photos {
		resp, err := http.Get(photo)
		if err != nil {
			return "fail get"
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return "fail read"
		}

		str := fmt.Sprintf("\033]1337;File=inline=1:%s\a\n", base64.StdEncoding.EncodeToString(data))
		datas += str
	}
	return datas
}

func RenderHTML(raw string) string {
	var result string
	result = strings.ReplaceAll(raw, "<br/>", "\n")
	return result
}
