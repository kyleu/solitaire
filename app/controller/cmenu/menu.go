// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"context"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/lib/filter"
	"github.com/kyleu/solitaire/app/lib/menu"
	"github.com/kyleu/solitaire/app/lib/sandbox"
	"github.com/kyleu/solitaire/app/lib/user"
	"github.com/kyleu/solitaire/app/util"
)

func MenuFor(
	ctx context.Context, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, as *app.State, logger util.Logger, //nolint:revive
) (menu.Items, any, error) {
	var ret menu.Items
	var data any
	// $PF_SECTION_START(routes_start)$
	// $PF_SECTION_END(routes_start)$
	// $PF_SECTION_START(routes_end)$
	// This is your menu, feel free to customize it
	admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
	ret = append(ret, gameMenu(ctx), menu.Separator, graphQLMenu(ctx, as.GraphQL), sandbox.Menu(ctx), admin)
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(routes_end)$
	return ret, data, nil
}
