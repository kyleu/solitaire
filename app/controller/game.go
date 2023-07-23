package controller

import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/vgame"
	"github.com/valyala/fasthttp"
)

func Game(rc *fasthttp.RequestCtx) {
	Act("game", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = "TODO"
		return Render(rc, as, &vgame.Game{}, ps)
	})
}
