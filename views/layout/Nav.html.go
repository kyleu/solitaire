// Code generated by qtc from "Nav.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Nav.html:2
package layout

//line views/layout/Nav.html:2
import (
	"fmt"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/lib/menu"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
)

//line views/layout/Nav.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Nav.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Nav.html:12
func StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:12
	qw422016.N().S(`
<nav id="navbar">
  <a class="logo" href="`)
//line views/layout/Nav.html:14
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:14
	qw422016.N().S(`" title="`)
//line views/layout/Nav.html:14
	qw422016.E().S(util.AppName)
//line views/layout/Nav.html:14
	qw422016.N().S(` `)
//line views/layout/Nav.html:14
	qw422016.E().S(as.BuildInfo.String())
//line views/layout/Nav.html:14
	qw422016.N().S(`">`)
//line views/layout/Nav.html:14
	components.StreamSVGSimple(qw422016, ps.RootIcon, 32, ps)
//line views/layout/Nav.html:14
	qw422016.N().S(`</a>
  <div class="breadcrumbs">
`)
//line views/layout/Nav.html:16
	extra := util.Choose(len(ps.Breadcrumbs) == 0, " simple", "")

//line views/layout/Nav.html:16
	qw422016.N().S(`    <a href="`)
//line views/layout/Nav.html:17
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:17
	qw422016.N().S(`" class="nav-root-icon`)
//line views/layout/Nav.html:17
	qw422016.E().S(extra)
//line views/layout/Nav.html:17
	qw422016.N().S(`" title="`)
//line views/layout/Nav.html:17
	qw422016.E().S(util.AppName)
//line views/layout/Nav.html:17
	qw422016.N().S(`">`)
//line views/layout/Nav.html:17
	components.StreamSVGBreadcrumb(qw422016, ps.RootIcon, ps)
//line views/layout/Nav.html:17
	qw422016.N().S(`</a>
    <a class="link nav-root-item`)
//line views/layout/Nav.html:18
	qw422016.E().S(extra)
//line views/layout/Nav.html:18
	qw422016.N().S(`" href="`)
//line views/layout/Nav.html:18
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:18
	qw422016.N().S(`">`)
//line views/layout/Nav.html:18
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:18
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:18
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:18
	qw422016.N().S(`
  </div>
`)
//line views/layout/Nav.html:20
	if ps.SearchPath != "-" {
//line views/layout/Nav.html:20
		qw422016.N().S(`  <form action="`)
//line views/layout/Nav.html:21
		qw422016.E().S(ps.SearchPath)
//line views/layout/Nav.html:21
		qw422016.N().S(`" class="search" title="search">
    <input id="search-input" type="search" name="q" placeholder=" " />
    <div class="search-image" style="display: none;"><svg><use xlink:href="#svg-searchbox" /></svg></div>
  </form>
`)
//line views/layout/Nav.html:25
	}
//line views/layout/Nav.html:25
	qw422016.N().S(`  `)
//line views/layout/Nav.html:26
	StreamProfileLink(qw422016, as, ps)
//line views/layout/Nav.html:26
	qw422016.N().S(`
`)
//line views/layout/Nav.html:27
	if !ps.HideMenu {
//line views/layout/Nav.html:27
		qw422016.N().S(`  <input type="checkbox" id="menu-toggle-input" style="display: none;" />
  <label class="menu-toggle" for="menu-toggle-input"><div class="spinner diagonal part-1"></div><div class="spinner horizontal"></div><div class="spinner diagonal part-2"></div></label>
  `)
//line views/layout/Nav.html:30
		StreamMenu(qw422016, ps)
//line views/layout/Nav.html:30
		qw422016.N().S(`
`)
//line views/layout/Nav.html:31
	}
//line views/layout/Nav.html:31
	qw422016.N().S(`</nav>`)
//line views/layout/Nav.html:32
}

//line views/layout/Nav.html:32
func WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:32
	StreamNav(qw422016, as, ps)
//line views/layout/Nav.html:32
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:32
}

//line views/layout/Nav.html:32
func Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:32
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:32
	WriteNav(qb422016, as, ps)
//line views/layout/Nav.html:32
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:32
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:32
	return qs422016
//line views/layout/Nav.html:32
}

//line views/layout/Nav.html:34
func StreamNavItem(qw422016 *qt422016.Writer, link string, title string, icon string, description string, last bool, ps *cutil.PageState) {
//line views/layout/Nav.html:35
	if link != "" || last {
//line views/layout/Nav.html:36
		extra := util.Choose(description == "", "", fmt.Sprintf(" title=%q", description))

//line views/layout/Nav.html:36
		qw422016.N().S(`<a class="link`)
//line views/layout/Nav.html:37
		if last {
//line views/layout/Nav.html:37
			qw422016.N().S(` `)
//line views/layout/Nav.html:37
			qw422016.N().S(`last`)
//line views/layout/Nav.html:37
		}
//line views/layout/Nav.html:37
		qw422016.N().S(`" href="`)
//line views/layout/Nav.html:37
		qw422016.E().S(link)
//line views/layout/Nav.html:37
		qw422016.N().S(`"`)
//line views/layout/Nav.html:37
		qw422016.N().S(extra)
//line views/layout/Nav.html:37
		qw422016.N().S(`>`)
//line views/layout/Nav.html:38
	}
//line views/layout/Nav.html:38
	qw422016.N().S(`<span>`)
//line views/layout/Nav.html:39
	components.StreamSVGBreadcrumb(qw422016, icon, ps)
//line views/layout/Nav.html:39
	qw422016.N().S(`</span><span class="nav-item-title">`)
//line views/layout/Nav.html:40
	qw422016.E().S(title)
//line views/layout/Nav.html:40
	qw422016.N().S(`</span>`)
//line views/layout/Nav.html:41
	if link != "" || last {
//line views/layout/Nav.html:41
		qw422016.N().S(`</a>`)
//line views/layout/Nav.html:43
	}
//line views/layout/Nav.html:44
}

//line views/layout/Nav.html:44
func WriteNavItem(qq422016 qtio422016.Writer, link string, title string, icon string, description string, last bool, ps *cutil.PageState) {
//line views/layout/Nav.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:44
	StreamNavItem(qw422016, link, title, icon, description, last, ps)
//line views/layout/Nav.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:44
}

//line views/layout/Nav.html:44
func NavItem(link string, title string, icon string, description string, last bool, ps *cutil.PageState) string {
//line views/layout/Nav.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:44
	WriteNavItem(qb422016, link, title, icon, description, last, ps)
//line views/layout/Nav.html:44
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:44
	return qs422016
//line views/layout/Nav.html:44
}

//line views/layout/Nav.html:46
func StreamNavItems(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:47
	for idx, bc := range ps.Breadcrumbs {
//line views/layout/Nav.html:49
		i := ps.Menu.GetByPath(ps.Breadcrumbs[:idx+1])
		if i == nil {
			i = menu.ItemFromString(bc, ps.DefaultNavIcon)
		}

//line views/layout/Nav.html:54
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:54
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:56
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:57
		StreamNavItem(qw422016, i.Route, i.Title, i.Icon, i.Description, idx == len(ps.Breadcrumbs)-1, ps)
//line views/layout/Nav.html:58
	}
//line views/layout/Nav.html:59
}

//line views/layout/Nav.html:59
func WriteNavItems(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:59
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:59
}

//line views/layout/Nav.html:59
func NavItems(ps *cutil.PageState) string {
//line views/layout/Nav.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:59
	WriteNavItems(qb422016, ps)
//line views/layout/Nav.html:59
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:59
	return qs422016
//line views/layout/Nav.html:59
}

//line views/layout/Nav.html:61
func StreamProfileLink(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:61
	qw422016.N().S(`<a class="profile" title="Settings" href="`)
//line views/layout/Nav.html:62
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:62
	qw422016.N().S(`">`)
//line views/layout/Nav.html:63
	components.StreamSVGSimple(qw422016, `profile`, 24, ps)
//line views/layout/Nav.html:63
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:65
}

//line views/layout/Nav.html:65
func WriteProfileLink(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:65
	StreamProfileLink(qw422016, as, ps)
//line views/layout/Nav.html:65
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:65
}

//line views/layout/Nav.html:65
func ProfileLink(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:65
	WriteProfileLink(qb422016, as, ps)
//line views/layout/Nav.html:65
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:65
	return qs422016
//line views/layout/Nav.html:65
}
