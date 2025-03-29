package sandbox

import (
	"context"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/util"
)

var wasm = &Sandbox{Key: "wasm", Title: "WASM", Icon: "gift", Run: onWASM}

func onWASM(_ context.Context, _ *app.State, _ util.ValueMap, _ util.Logger) (any, error) {
	ret := util.ValueMap{"status": "ok"}
	return ret, nil
}
