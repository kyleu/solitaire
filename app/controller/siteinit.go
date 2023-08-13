// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
)

// Initialize system dependencies for the marketing site.
func initSite(_ *app.State, _ util.Logger) {
}

// Configure marketing site data for each request.
func initSiteRequest(*app.State, *cutil.PageState) error {
	return nil
}
