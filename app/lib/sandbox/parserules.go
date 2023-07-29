package sandbox

import (
	"context"
	"os"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/app/parse"
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
		return parse.RulesContent(content, logger)
	})
	return ret, nil
}
