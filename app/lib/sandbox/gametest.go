package sandbox

import (
	"context"
	"github.com/kyleu/solitaire/app/card"
	"github.com/kyleu/solitaire/app/pile"
	"github.com/kyleu/solitaire/app/rules"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/util"
)

var gametest = &Sandbox{Key: "gametest", Title: "Game Test", Icon: "heart", Run: onGametest}

func onGametest(_ context.Context, _ *app.State, _ util.Logger) (any, error) {
	ret := pile.Pile{ID: "test", Options: &pile.Options{}, Cards: card.RandomCards(5, card.SuitsCommon, 1)}
	outcome := rules.CheckPoker(ret.Cards)
	return util.ValueMap{"cards": ret, "outcome": outcome}, nil
}
