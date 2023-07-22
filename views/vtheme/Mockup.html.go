// Code generated by qtc from "Mockup.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/Mockup.html:2
package vtheme

//line views/vtheme/Mockup.html:2
import (
	"fmt"

	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/lib/theme"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/vutil"
)

//line views/vtheme/Mockup.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Mockup.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Mockup.html:12
func StreamMockupTheme(qw422016 *qt422016.Writer, t *theme.Theme, pointer bool, icon string, indent int, ps *cutil.PageState) {
//line views/vtheme/Mockup.html:12
	qw422016.N().S(`<div class="title small-text">`)
//line views/vtheme/Mockup.html:13
	qw422016.E().S(t.Key)
//line views/vtheme/Mockup.html:13
	qw422016.N().S(`</div>`)
//line views/vtheme/Mockup.html:14
	StreamMockupColors(qw422016, util.AppName, "light", t.Light, pointer, icon, indent, ps)
//line views/vtheme/Mockup.html:15
	StreamMockupColors(qw422016, util.AppName, "dark", t.Dark, pointer, icon, indent, ps)
//line views/vtheme/Mockup.html:16
}

//line views/vtheme/Mockup.html:16
func WriteMockupTheme(qq422016 qtio422016.Writer, t *theme.Theme, pointer bool, icon string, indent int, ps *cutil.PageState) {
//line views/vtheme/Mockup.html:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Mockup.html:16
	StreamMockupTheme(qw422016, t, pointer, icon, indent, ps)
//line views/vtheme/Mockup.html:16
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Mockup.html:16
}

//line views/vtheme/Mockup.html:16
func MockupTheme(t *theme.Theme, pointer bool, icon string, indent int, ps *cutil.PageState) string {
//line views/vtheme/Mockup.html:16
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Mockup.html:16
	WriteMockupTheme(qb422016, t, pointer, icon, indent, ps)
//line views/vtheme/Mockup.html:16
	qs422016 := string(qb422016.B)
//line views/vtheme/Mockup.html:16
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Mockup.html:16
	return qs422016
//line views/vtheme/Mockup.html:16
}

//line views/vtheme/Mockup.html:18
func StreamMockupColors(qw422016 *qt422016.Writer, navTitle string, mode string, c *theme.Colors, pointer bool, icon string, indent int, ps *cutil.PageState) {
//line views/vtheme/Mockup.html:20
	navStyle := fmt.Sprintf("color: %s; background-color: %s;", c.NavForeground, c.NavBackground)

	menuStyle := fmt.Sprintf("color: %s; background-color: %s;", c.MenuForeground, c.MenuBackground)
	menuLinkStyle := fmt.Sprintf("color: %s;", c.MenuForeground)
	menuLinkSelectedStyle := fmt.Sprintf("color: %s; background-color: %s;", c.MenuSelectedForeground, c.MenuSelectedBackground)

	mainStyle := fmt.Sprintf("color: %s; background-color: %s;", c.Foreground, c.Background)
	mutedStyle := fmt.Sprintf("color: %s; background-color: %s;", c.ForegroundMuted, c.BackgroundMuted)
	linkStyle := fmt.Sprintf("color: %s;", c.LinkForeground)
	linkVisitedStyle := fmt.Sprintf("color: %s;", c.LinkVisitedForeground)

	cls := "mockup"
	if mode != "" {
		cls += " only-" + mode
	}
	if pointer {
		cls += " pointer"
	}

//line views/vtheme/Mockup.html:39
	vutil.StreamIndent(qw422016, true, indent)
//line views/vtheme/Mockup.html:39
	qw422016.N().S(`<div class="`)
//line views/vtheme/Mockup.html:40
	qw422016.E().S(cls)
//line views/vtheme/Mockup.html:40
	qw422016.N().S(`"><div class="mock-nav" style="`)
//line views/vtheme/Mockup.html:41
	qw422016.E().S(navStyle)
//line views/vtheme/Mockup.html:41
	qw422016.N().S(`">`)
//line views/vtheme/Mockup.html:41
	components.StreamSVGRef(qw422016, icon, 12, 12, `icon`, ps)
//line views/vtheme/Mockup.html:41
	qw422016.E().S(navTitle)
//line views/vtheme/Mockup.html:41
	qw422016.N().S(`</div><div class="mock-menu" style="`)
//line views/vtheme/Mockup.html:42
	qw422016.E().S(menuStyle)
//line views/vtheme/Mockup.html:42
	qw422016.N().S(`"><div class="mock-link" style="`)
//line views/vtheme/Mockup.html:43
	qw422016.E().S(menuLinkStyle)
//line views/vtheme/Mockup.html:43
	qw422016.N().S(`">A</div><div class="mock-link-selected" style="`)
//line views/vtheme/Mockup.html:44
	qw422016.E().S(menuLinkSelectedStyle)
//line views/vtheme/Mockup.html:44
	qw422016.N().S(`">B</div><div class="mock-link" style="`)
//line views/vtheme/Mockup.html:45
	qw422016.E().S(menuLinkStyle)
//line views/vtheme/Mockup.html:45
	qw422016.N().S(`">C</div><div class="mock-link" style="`)
//line views/vtheme/Mockup.html:46
	qw422016.E().S(menuLinkStyle)
//line views/vtheme/Mockup.html:46
	qw422016.N().S(`">D</div></div><div class="mock-main" style="`)
//line views/vtheme/Mockup.html:48
	qw422016.E().S(mainStyle)
//line views/vtheme/Mockup.html:48
	qw422016.N().S(`"><div class="mock-muted" style="`)
//line views/vtheme/Mockup.html:49
	qw422016.E().S(mutedStyle)
//line views/vtheme/Mockup.html:49
	qw422016.N().S(`">Welcome!</div><div><div class="mock-list">Here's some links:</div><ul><li class="mock-link" style="`)
//line views/vtheme/Mockup.html:53
	qw422016.E().S(linkStyle)
//line views/vtheme/Mockup.html:53
	qw422016.N().S(`">New</li><li class="mock-link" style="`)
//line views/vtheme/Mockup.html:54
	qw422016.E().S(linkStyle)
//line views/vtheme/Mockup.html:54
	qw422016.N().S(`">Also New</li><li class="mock-link-visited" style="`)
//line views/vtheme/Mockup.html:55
	qw422016.E().S(linkVisitedStyle)
//line views/vtheme/Mockup.html:55
	qw422016.N().S(`">Visited</li></ul></div></div></div>`)
//line views/vtheme/Mockup.html:60
}

//line views/vtheme/Mockup.html:60
func WriteMockupColors(qq422016 qtio422016.Writer, navTitle string, mode string, c *theme.Colors, pointer bool, icon string, indent int, ps *cutil.PageState) {
//line views/vtheme/Mockup.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Mockup.html:60
	StreamMockupColors(qw422016, navTitle, mode, c, pointer, icon, indent, ps)
//line views/vtheme/Mockup.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Mockup.html:60
}

//line views/vtheme/Mockup.html:60
func MockupColors(navTitle string, mode string, c *theme.Colors, pointer bool, icon string, indent int, ps *cutil.PageState) string {
//line views/vtheme/Mockup.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Mockup.html:60
	WriteMockupColors(qb422016, navTitle, mode, c, pointer, icon, indent, ps)
//line views/vtheme/Mockup.html:60
	qs422016 := string(qb422016.B)
//line views/vtheme/Mockup.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Mockup.html:60
	return qs422016
//line views/vtheme/Mockup.html:60
}
