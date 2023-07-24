package sandbox

import (
	"context"
	"github.com/pkg/errors"
	"os"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/app/util"
)

var parseRules = &Sandbox{Key: "parseRules", Title: "Parse Rules", Icon: "heart", Run: onParseRules}

func onParseRules(_ context.Context, _ *app.State, logger util.Logger) (any, error) {
	fs := filesystem.NewFileSystem("../solitaire.gg/shared/src/main/scala/models/rules/impl")
	files := lo.Map(fs.ListFiles(".", nil, logger), func(e os.DirEntry, _ int) string {
		return e.Name()
	})
	ret := lo.Map(files, func(filename string, _ int) *rules.Rules {
		b, err := fs.ReadFile(filename)
		if err != nil {
			name := strings.TrimSuffix(filename, ".scala")
			return &rules.Rules{Name: name, Error: err.Error()}
		}
		content := string(b)
		return parseFile(content, logger)
		// return parseLines(util.StringSplitAndTrim(content, "\n"), logger)
	})
	return ret, nil
}

func parseFile(content string, logger util.Logger) *rules.Rules {
	sections := splitSections(content)
	ret := &rules.Rules{}
	for _, section := range sections {
		k, v := util.StringSplit(section, '=', true)
		k, v = strings.TrimSpace(k), cleanLine(v)
		v = strings.TrimSuffix(strings.TrimPrefix(v, `"`), `"`)
		err := parseKV(k, v, ret, logger)
		if err != nil {
			ret.Context = append(ret.Context, err.Error())
		}
	}
	return ret
}

func parseKV(k string, v string, r *rules.Rules, logger util.Logger) error {
	switch k {
	case "id":
		r.Key = v
	case "title":
		if v != "" {
			r.Name = v
		}
	case "completed":
		r.Completed = v == "true"
	case "layout":
		r.Layout = v
	case "like":
		r.Like = v
	default:
		return errors.Errorf("[%s]: %s", k, v)
	}
	return nil
}

func cleanLine(l string) string {
	l = strings.TrimSuffix(strings.TrimSpace(l), ",")
	if strings.HasPrefix(l, "Some(") {
		l = strings.TrimPrefix(strings.TrimSuffix(l, ")"), "Some(")
	}
	return l
}

func splitSections(content string) []string {
	lines := strings.Split(content, "\n")
	var ret []string
	var curr string
	for idx := range lines {
		line := lines[idx]
		clean := cleanLine(line)
		if len(clean) == 0 || strings.HasPrefix(clean, "package ") || strings.HasPrefix(clean, "import ") || strings.HasPrefix(clean, "object ") {
			continue
		}
		if strings.HasPrefix(line, "  ") {
			if strings.HasPrefix(line, "    ") {
				curr += "\n" + line
				continue
			}
			if clean == ")" {
				curr += "\n" + line
				ret = append(ret, curr)
				curr = ""
				continue
			}
			if strings.Contains(line, "(") {
				if strings.Contains(line, ")") {
					ret = append(ret, line)
				} else {
					curr = line
				}
				continue
			}
			ret = append(ret, line)
		}
	}
	if curr != "" {
		ret = append(ret, curr)
	}
	return ret
}
