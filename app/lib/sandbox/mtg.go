package sandbox

import (
	"context"
	"encoding/json"
	"os"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/lib/log"
	"github.com/kyleu/solitaire/app/util"
)

var mtg = &Sandbox{Key: "mtg", Title: "Magic: The Gathering", Icon: "mobile", Run: onMTG}

func onMTG(ctx context.Context, _ *app.State, logger util.Logger) (any, error) {
	logger.Info("[mtg::0] starting file read...")
	t := log.NewTimer("mtg", logger)
	b, err := os.ReadFile("./tmp/scryfall.json")
	if err != nil {
		return nil, err
	}

	t.Lap("file read complete, parsing JSON...")
	messages := make([]json.RawMessage, 0, 10000)
	err = util.FromJSON(b, &messages)
	if err != nil {
		return nil, err
	}

	t.Lap("parse complete, handling [%d] rows...", len(messages))
	ret := lo.Map(messages, func(js json.RawMessage, idx int) string {
		if idx > 0 && idx%10000 == 0 {
			t.Lap("completed [%d] rows...", idx)
		}
		return ""
	})

	t.Complete()
	return len(ret), nil
}
