package sandbox

import (
	"context"
	"os"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/app/util"
)

var parseRules = &Sandbox{Key: "parseRules", Title: "Parse Rules", Icon: "heart", Run: onParseRules}

type Rules struct {
	Key       string            `json:"key"`
	Completed bool              `json:"completed,omitempty"`
	Name      string            `json:"name"`
	AKA       map[string]string `json:"aka,omitempty"`
	Links     map[string]string `json:"links,omitempty"`
	Layout    string            `json:"layout"`
	Like      string            `json:"like,omitempty"`
	Error     string            `json:"error,omitempty"`
}

func onParseRules(_ context.Context, _ *app.State, logger util.Logger) (any, error) {
	fs := filesystem.NewFileSystem("../solitaire.gg/shared/src/main/scala/models/rules/impl")
	files := lo.Map(fs.ListFiles(".", nil, logger), func(e os.DirEntry, _ int) string {
		return e.Name()
	})
	ret := lo.Map(files, func(filename string, _ int) *Rules {
		b, err := fs.ReadFile(filename)
		if err != nil {
			name := strings.TrimSuffix(filename, ".scala")
			return &Rules{Name: name, Error: err.Error()}
		}
		return parseLines(util.StringSplitAndTrim(string(b), "\n"))
	})
	return ret, nil
}

func parseLines(lines []string) *Rules {
	ret := &Rules{}
	mode := ""
	section := ""
	for _, line := range lines {
		if mode == "" {
			mode, section = parseNormalLine(line, ret, mode, section)
		} else {
			println("???", line)
		}
	}
	return ret
}

func parseNormalLine(line string, rl *Rules, mode string, section string) (string, string) {
	if strings.Contains(line, "=") {
		k, v := util.StringSplit(line, '=', true)
		k, v = strings.TrimSpace(k), strings.TrimSuffix(strings.TrimSpace(v), ",")
		if strings.HasPrefix(v, "Some(") {
			v = strings.TrimSuffix(strings.TrimPrefix(v, `Some(`), `)`)
		}
		v = strings.TrimSuffix(strings.TrimPrefix(v, `"`), `"`)
		switch k {
		case "id":
			rl.Key = v
		case "title":
			if v != "" {
				rl.Name = v
			}
		case "completed":
			rl.Completed = v == "true"
		case "layout":
			rl.Layout = v
		case "like":
			rl.Like = v
		default:
			println(k, "/", v)
		}
	}
	return mode, section
}
