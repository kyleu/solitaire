// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/http"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views"
)

func About(w http.ResponseWriter, r *http.Request) {
	controller.Act("about", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.SetTitleAndData("About "+util.AppName, util.AppName+" v"+as.BuildInfo.Version)
		return controller.Render(w, r, as, &views.About{}, ps, "about")
	})
}
