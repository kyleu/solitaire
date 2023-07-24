package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/game"
	"github.com/kyleu/solitaire/app/game/card"
	"github.com/kyleu/solitaire/app/game/pile"
	"github.com/kyleu/solitaire/app/game/poker"
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/game/suit"
	"github.com/kyleu/solitaire/views"
	"github.com/kyleu/solitaire/views/vgame"
)

func Game(rc *fasthttp.RequestCtx) {
	Act("game", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = []string{"/game/test/html", "/game/test/wasm"}
		return Render(rc, as, &vgame.Index{}, ps, "game")
	})
}

type TestJSONResponse struct {
	Game        any `json:"game,omitempty"`
	PokerHand   any `json:"pokerHand,omitempty"`
	PokerResult any `json:"pokerResult,omitempty"`
}

func GameTestJSON(rc *fasthttp.RequestCtx) {
	Act("game.test.json", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		g := game.New(rules.Example)
		hand := pile.Pile{ID: "test", Options: &pile.Options{}, Cards: card.RandomCards(5, suit.SuitsCommon, 1)}
		result := poker.PokerCheck(hand.Cards)
		ret := &TestJSONResponse{Game: g, PokerHand: hand, PokerResult: result}
		ps.Data = ret
		ps.Title = "JSON Test"
		return Render(rc, as, &views.Debug{}, ps, "game", "json")
	})
}

func GameTestHTML(rc *fasthttp.RequestCtx) {
	Act("game.test.html", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		g := "TODO"
		ps.Data = g
		ps.Title = "HTML Test"
		return Render(rc, as, &vgame.HTML{}, ps, "game", "html")
	})
}

func GameTestWASM(rc *fasthttp.RequestCtx) {
	Act("game.test.wasm", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		g := "TODO"
		ps.Data = g
		ps.Title = "WASM Test"
		return Render(rc, as, &vgame.WASM{}, ps, "game", "wasm")
	})
}
