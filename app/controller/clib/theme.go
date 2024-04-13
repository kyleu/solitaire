// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller"
	"github.com/kyleu/solitaire/app/controller/csession"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/lib/theme"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/vtheme"
)

const themeIcon = "gift"

func ThemeList(w http.ResponseWriter, r *http.Request) {
	controller.Act("theme.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		th := as.Themes.All(ps.Logger)
		ps.SetTitleAndData("Themes", th)
		ps.DefaultNavIcon = themeIcon
		return controller.Render(w, r, as, &vtheme.List{Themes: th}, ps, "Themes||/theme")
	})
}

func ThemeEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("theme.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}
		if key == theme.Default.Key {
			return controller.FlashAndRedir(false, "Unable to edit default theme", "/theme", w, ps)
		}
		var t *theme.Theme
		if key == theme.KeyNew {
			t = theme.Default.Clone(key)
		} else {
			t = as.Themes.Get(key, ps.Logger)
			if t == nil {
				if pal := r.URL.Query().Get("palette"); pal != "" {
					themes, err := theme.PaletteThemes(pal)
					if err != nil {
						return "", err
					}
					t = themes.Get(key)
				}
			}
		}
		if t == nil {
			return "", errors.Errorf("invalid theme [%s]", key)
		}
		ps.SetTitleAndData("Edit theme ["+t.Key+"]", t)
		ps.DefaultNavIcon = themeIcon
		page := &vtheme.Edit{Theme: t, Icon: "app", Exists: as.Themes.FileExists(t.Key)}
		return controller.Render(w, r, as, page, ps, "Themes||/theme", t.Key)
	})
}

func ThemeSave(w http.ResponseWriter, r *http.Request) {
	controller.Act("theme.save", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", err
		}

		orig := as.Themes.Get(key, ps.Logger)

		newKey, err := frm.GetString("key", false)
		if err != nil {
			return "", err
		}
		if newKey == theme.KeyNew {
			newKey = util.RandomString(12)
		}

		l := orig.Light.Clone().ApplyMap(frm, "light-")
		d := orig.Dark.Clone().ApplyMap(frm, "dark-")

		t := &theme.Theme{Key: newKey, Light: l, Dark: d}

		err = as.Themes.Save(t, frm.GetStringOpt("originalKey"), ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to save theme")
		}

		ps.Profile.Theme = newKey
		err = csession.SaveProfile(ps.Profile, w, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		return controller.ReturnToReferrer("saved changes to theme ["+newKey+"]", "/theme", w, ps)
	})
}

func ThemeRemove(w http.ResponseWriter, r *http.Request) {
	controller.Act("theme.remove", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Themes.Remove(key, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to remove theme")
		}
		return controller.ReturnToReferrer("removed theme ["+key+"]", "/theme", w, ps)
	})
}
