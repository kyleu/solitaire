package sandbox

import (
	"context"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/parse"
	"github.com/kyleu/solitaire/app/util"
)

var parseRules = &Sandbox{Key: "parseRules", Title: "Parse Rules", Icon: "heart", Run: onParseRules}

func onParseRules(ctx context.Context, as *app.State, _ util.ValueMap, logger util.Logger) (any, error) {
	return parse.ParseJSONRules(ctx, logger)
}
