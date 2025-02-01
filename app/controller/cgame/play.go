package cgame

import (
	"net/http"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/game"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/assets"
	"github.com/kyleu/solitaire/views/vgame"
)

func PlayLocal(w http.ResponseWriter, r *http.Request) {
	controller.Act("play.local", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Play(as, nil, true, r, ps)
	})
}

func PlayServer(w http.ResponseWriter, r *http.Request) {
	controller.Act("play.server", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		return Play(as, nil, false, r, ps)
	})
}

func Play(as *app.State, g *game.Game, offline bool, r *http.Request, ps *cutil.PageState) (string, error) {
	applyAssets(ps)
	sp := util.Choose(offline, "", "/TODO")
	// ps.HideMenu = true
	return controller.Render(r, as, &vgame.Play{Game: g, SocketPath: sp}, ps, "Play")
}

func applyAssets(ps *cutil.PageState) {
	ps.HeaderContent = assets.ScriptElement(`game.js`, true) + assets.StylesheetElement(`game.css`)
}
