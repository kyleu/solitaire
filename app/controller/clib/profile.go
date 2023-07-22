// Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/url"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller"
	"github.com/kyleu/solitaire/app/controller/csession"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/lib/theme"
	"github.com/kyleu/solitaire/views/vprofile"
)

func Profile(rc *fasthttp.RequestCtx) {
	controller.Act("profile", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(rc, as, ps)
	})
}

func ProfileSite(rc *fasthttp.RequestCtx) {
	controller.ActSite("profile", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(rc, as, ps)
	})
}

func profileAction(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (string, error) {
	ps.Title = "Profile"
	ps.Data = ps.Profile
	thm := as.Themes.Get(ps.Profile.Theme, ps.Logger)

	redir := "/"
	ref := string(rc.Request.Header.Peek("Referer"))
	if ref != "" {
		u, err := url.Parse(ref)
		if err == nil && u != nil && u.Path != cutil.DefaultProfilePath {
			redir = u.Path
		}
	}

	page := &vprofile.Profile{Profile: ps.Profile, Theme: thm, Referrer: redir}
	return controller.Render(rc, as, page, ps, "Profile")
}

func ProfileSave(rc *fasthttp.RequestCtx) {
	controller.Act("profile.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}

		n := ps.Profile.Clone()

		referrerDefault := frm.GetStringOpt("referrer")
		if referrerDefault == "" {
			referrerDefault = cutil.DefaultProfilePath
		}

		n.Name = frm.GetStringOpt("name")
		n.Mode = frm.GetStringOpt("mode")
		n.Theme = frm.GetStringOpt("theme")
		if n.Theme == theme.ThemeDefault.Key {
			n.Theme = ""
		}

		err = csession.SaveProfile(n, rc, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		return controller.ReturnToReferrer("Saved profile", referrerDefault, rc, ps)
	})
}
