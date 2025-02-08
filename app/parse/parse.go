package parse

import (
	"context"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/app/util"
)

func ParseJSONRules(_ context.Context, _ util.Logger) (any, error) {
	fs, err := filesystem.NewFileSystem("./tmp", false, "")
	if err != nil {
		return nil, err
	}
	b, err := fs.ReadFile("rules.json")
	if err != nil {
		return nil, err
	}
	var ret []*JSONRules
	err = util.FromJSONStrict(b, &ret)
	return lo.Map(ret, func(x *JSONRules, _ int) *rules.Rules {
		return x.ToRules()
	}), err
}
