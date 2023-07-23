package controller

import (
	"github.com/kyleu/solitaire/app/card"
	"github.com/kyleu/solitaire/app/pile"
	"github.com/kyleu/solitaire/app/rules"
	"github.com/kyleu/solitaire/app/suit"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/vgame"
)

func Game(rc *fasthttp.RequestCtx) {
	Act("game", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = []string{"/game/test/html", "/game/test/wasm"}
		return Render(rc, as, &vgame.Index{}, ps, "game")
	})
}

func GameTestJSON(rc *fasthttp.RequestCtx) {
	Act("game.test.json", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		pl := pile.Pile{ID: "test", Options: &pile.Options{}, Cards: card.RandomCards(5, suit.SuitsCommon, 1)}
		outcome := rules.CheckPoker(pl.Cards)
		ret := util.ValueMap{"cards": pl, "outcome": outcome}
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
