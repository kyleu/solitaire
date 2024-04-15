package controller

import (
	"net/http"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/game"
	"github.com/kyleu/solitaire/app/game/card"
	"github.com/kyleu/solitaire/app/game/pile"
	"github.com/kyleu/solitaire/app/game/poker"
	"github.com/kyleu/solitaire/app/game/rules/gamerules"
	"github.com/kyleu/solitaire/app/game/suit"
	"github.com/kyleu/solitaire/views"
	"github.com/kyleu/solitaire/views/vgame"
)

func Game(w http.ResponseWriter, r *http.Request) {
	Act("game", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = []string{"/game/test/html", "/game/test/wasm"}
		return Render(r, as, &vgame.Index{}, ps, "game")
	})
}

type TestJSONResponse struct {
	Game        any `json:"game,omitempty"`
	PokerHand   any `json:"pokerHand,omitempty"`
	PokerResult any `json:"pokerResult,omitempty"`
}

func GameTestJSON(w http.ResponseWriter, r *http.Request) {
	Act("game.test.json", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		g := game.New(gamerules.Example)
		hand := pile.Pile{ID: "test", Options: &pile.Options{}, Cards: card.RandomCards(5, suit.SuitsCommon, 1)}
		result := poker.Check(hand.Cards)
		ret := &TestJSONResponse{Game: g, PokerHand: hand, PokerResult: result}
		ps.Data = ret
		ps.Title = "JSON Test"
		return Render(r, as, &views.Debug{}, ps, "game", "json")
	})
}

func GameTestHTML(w http.ResponseWriter, r *http.Request) {
	Act("game.test.html", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		g := game.New(gamerules.Example)
		ps.Data = g
		ps.Title = "HTML Test"
		return Render(r, as, &vgame.HTML{Game: g}, ps, "game", "html")
	})
}

func GameTestWASM(w http.ResponseWriter, r *http.Request) {
	Act("game.test.wasm", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		g := game.New(gamerules.Example)
		ps.Data = g
		ps.Title = "WASM Test"
		return Render(r, as, &vgame.WASM{Game: g}, ps, "game", "wasm")
	})
}
