// Code generated by qtc from "Nav.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Nav.html:2
package layout

//line views/layout/Nav.html:2
import (
	"strings"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/vutil"
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
	components.StreamSVGRef(qw422016, ps.RootIcon, 32, 32, ``, ps)
//line views/layout/Nav.html:14
	qw422016.N().S(`</a>
  <div class="breadcrumbs">
    <a class="link" href="`)
//line views/layout/Nav.html:16
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:16
	qw422016.N().S(`">`)
//line views/layout/Nav.html:16
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:16
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:16
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:16
	qw422016.N().S(`
  </div>
`)
//line views/layout/Nav.html:18
	if ps.SearchPath != "-" {
//line views/layout/Nav.html:18
		qw422016.N().S(`  <form action="`)
//line views/layout/Nav.html:19
		qw422016.E().S(ps.SearchPath)
//line views/layout/Nav.html:19
		qw422016.N().S(`" class="search" title="search">
    <input type="search" name="q" placeholder=" " />
    <div class="search-image" style="display: none;"><svg><use xlink:href="#svg-searchbox" /></svg></div>
  </form>
`)
//line views/layout/Nav.html:23
	}
//line views/layout/Nav.html:23
	qw422016.N().S(`  <a class="profile" title="Settings" href="`)
//line views/layout/Nav.html:24
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:24
	qw422016.N().S(`">
    `)
//line views/layout/Nav.html:25
	components.StreamSVGRef(qw422016, `profile`, 24, 24, ``, ps)
//line views/layout/Nav.html:25
	qw422016.N().S(`
  </a>
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
func StreamNavItems(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:35
	for idx, bc := range ps.Breadcrumbs {
//line views/layout/Nav.html:37
		i := ps.Menu.GetByPath(ps.Breadcrumbs[:idx+1])

//line views/layout/Nav.html:39
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:39
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:41
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:42
		if i == nil {
//line views/layout/Nav.html:44
			bcLink := ""
			if strings.Contains(bc, "||") {
				bci := strings.Index(bc, "||")
				bcLink = bc[bci+2:]
				bc = bc[:bci]
			}

//line views/layout/Nav.html:50
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:51
			qw422016.E().S(bcLink)
//line views/layout/Nav.html:51
			qw422016.N().S(`">`)
//line views/layout/Nav.html:51
			qw422016.E().S(bc)
//line views/layout/Nav.html:51
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:52
		} else {
//line views/layout/Nav.html:52
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:53
			qw422016.E().S(i.Route)
//line views/layout/Nav.html:53
			qw422016.N().S(`">`)
//line views/layout/Nav.html:53
			components.StreamSVGRef(qw422016, i.Icon, 28, 28, "breadcrumb-icon", ps)
//line views/layout/Nav.html:53
			qw422016.E().S(i.Title)
//line views/layout/Nav.html:53
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:54
		}
//line views/layout/Nav.html:55
	}
//line views/layout/Nav.html:56
}

//line views/layout/Nav.html:56
func WriteNavItems(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:56
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:56
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:56
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:56
}

//line views/layout/Nav.html:56
func NavItems(ps *cutil.PageState) string {
//line views/layout/Nav.html:56
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:56
	WriteNavItems(qb422016, ps)
//line views/layout/Nav.html:56
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:56
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:56
	return qs422016
//line views/layout/Nav.html:56
}
