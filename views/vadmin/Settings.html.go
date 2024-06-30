// Code generated by qtc from "Settings.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vadmin/Settings.html:1
package vadmin

//line views/vadmin/Settings.html:1
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/components/view"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vadmin/Settings.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Settings.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Settings.html:10
type Settings struct {
	layout.Basic
}

//line views/vadmin/Settings.html:14
func (p *Settings) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:14
	qw422016.N().S(`
  <div class="card">
`)
//line views/vadmin/Settings.html:16
	if util.AppSource != "" {
//line views/vadmin/Settings.html:16
		qw422016.N().S(`    <div class="right"><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:17
		qw422016.E().S(util.AppSource)
//line views/vadmin/Settings.html:17
		qw422016.N().S(`"><button>GitHub</button></a></div>
`)
//line views/vadmin/Settings.html:18
	}
//line views/vadmin/Settings.html:18
	qw422016.N().S(`    <h3 title="github:`)
//line views/vadmin/Settings.html:19
	qw422016.E().S(as.BuildInfo.Commit)
//line views/vadmin/Settings.html:19
	qw422016.N().S(`">`)
//line views/vadmin/Settings.html:19
	components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:19
	qw422016.N().S(` `)
//line views/vadmin/Settings.html:19
	qw422016.E().S(util.AppName)
//line views/vadmin/Settings.html:19
	qw422016.N().S(` `)
//line views/vadmin/Settings.html:19
	qw422016.E().S(as.BuildInfo.String())
//line views/vadmin/Settings.html:19
	qw422016.N().S(`</h3>
`)
//line views/vadmin/Settings.html:20
	if util.AppLegal != "" {
//line views/vadmin/Settings.html:20
		qw422016.N().S(`    <div class="mt">`)
//line views/vadmin/Settings.html:21
		qw422016.N().S(util.AppLegal)
//line views/vadmin/Settings.html:21
		qw422016.N().S(`</div>
`)
//line views/vadmin/Settings.html:22
	}
//line views/vadmin/Settings.html:23
	if util.AppURL != "" {
//line views/vadmin/Settings.html:23
		qw422016.N().S(`    <p>`)
//line views/vadmin/Settings.html:24
		view.StreamURL(qw422016, util.AppURL, "", true, ps)
//line views/vadmin/Settings.html:24
		qw422016.N().S(`</p>
`)
//line views/vadmin/Settings.html:25
	}
//line views/vadmin/Settings.html:25
	qw422016.N().S(`    <em>This page is for the settings of the application. To change your user preferences, such as themes, <a href="/profile">edit your profile</a>.</em>
  </div>

  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:30
	components.StreamSVGIcon(qw422016, `archive`, ps)
//line views/vadmin/Settings.html:30
	qw422016.N().S(` Admin Functions</h3>
    `)
//line views/vadmin/Settings.html:31
	streamsettingsLink(qw422016, "/admin/server", "archive", "App Information", "All sorts of info about the server and runtime", ps)
//line views/vadmin/Settings.html:31
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:32
	streamsettingsLink(qw422016, "/admin/modules", "robot", "Go Modules", "The Go modules used by "+util.AppName, ps)
//line views/vadmin/Settings.html:32
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:33
	streamsettingsLink(qw422016, "/theme", "archive", "Edit Themes", "Configure the design themes available to end users", ps)
//line views/vadmin/Settings.html:33
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:34
	streamsettingsLink(qw422016, "/admin/logs", "folder", "Recent Logs", "Displays the 100 most recent app logs", ps)
//line views/vadmin/Settings.html:34
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:38
	components.StreamSVGIcon(qw422016, `bolt`, ps)
//line views/vadmin/Settings.html:38
	qw422016.N().S(` HTTP Methods</h3>
    `)
//line views/vadmin/Settings.html:39
	streamsettingsLink(qw422016, "/admin/sitemap", "graph", "Sitemap", "Displays the HTTP actions that are available, with documentation", ps)
//line views/vadmin/Settings.html:39
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:40
	streamsettingsLink(qw422016, "/admin/routes", "folder", "HTTP routes", "Enumerates all registered HTTP routes, by method", ps)
//line views/vadmin/Settings.html:40
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:41
	streamsettingsLink(qw422016, "/admin/session", "play", "User Session", "View the user session, including all cookies and settings", ps)
//line views/vadmin/Settings.html:41
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:42
	streamsettingsLink(qw422016, "/admin/request", "download", "Debug HTTP Request", "Full debug view of an HTTP request from your browser", ps)
//line views/vadmin/Settings.html:42
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:43
	streamsettingsLink(qw422016, "/admin/sockets", "cog", "Active WebSockets", "Manage the active WebSockets in this server", ps)
//line views/vadmin/Settings.html:43
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:47
	components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:47
	qw422016.N().S(` App Profiling</h3>
    `)
//line views/vadmin/Settings.html:48
	streamsettingsLink(qw422016, "/admin/memusage", "desktop", "Memory Usage", "Detailed memory usage statistics for this application", ps)
//line views/vadmin/Settings.html:48
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:49
	streamsettingsLink(qw422016, "/admin/gc", "cog", "Collect Garbage", "Runs the Go garbage collector", ps)
//line views/vadmin/Settings.html:49
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:50
	streamsettingsLink(qw422016, "/admin/heap", "cog", "Write Memory Dump", "Writes a memory dump to <em>./tmp/mem.pprof</em>, use script to view", ps)
//line views/vadmin/Settings.html:50
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:51
	streamsettingsLink(qw422016, "/admin/cpu/start", "cog", "Start CPU Profile", "Profiles the CPU using <em>./tmp/cpu.pprof</em>, use script to view", ps)
//line views/vadmin/Settings.html:51
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:52
	streamsettingsLink(qw422016, "/admin/cpu/stop", "cog", "Stop CPU Profile", "Stops the active CPU profile", ps)
//line views/vadmin/Settings.html:52
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
`)
//line views/vadmin/Settings.html:55
}

//line views/vadmin/Settings.html:55
func (p *Settings) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:55
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Settings.html:55
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:55
}

//line views/vadmin/Settings.html:55
func (p *Settings) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:55
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:55
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Settings.html:55
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:55
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:55
	return qs422016
//line views/vadmin/Settings.html:55
}

//line views/vadmin/Settings.html:57
func streamsettingsLink(qw422016 *qt422016.Writer, href string, icon string, title string, description string, ps *cutil.PageState) {
//line views/vadmin/Settings.html:57
	qw422016.N().S(`<hr class="clear" /><div class="mts ml"><a href="`)
//line views/vadmin/Settings.html:60
	qw422016.E().S(href)
//line views/vadmin/Settings.html:60
	qw422016.N().S(`"><strong>`)
//line views/vadmin/Settings.html:60
	qw422016.E().S(title)
//line views/vadmin/Settings.html:60
	qw422016.N().S(`</strong></a><div><em>`)
//line views/vadmin/Settings.html:61
	qw422016.N().S(description)
//line views/vadmin/Settings.html:61
	qw422016.N().S(`</em></div></div>`)
//line views/vadmin/Settings.html:63
}

//line views/vadmin/Settings.html:63
func writesettingsLink(qq422016 qtio422016.Writer, href string, icon string, title string, description string, ps *cutil.PageState) {
//line views/vadmin/Settings.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:63
	streamsettingsLink(qw422016, href, icon, title, description, ps)
//line views/vadmin/Settings.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:63
}

//line views/vadmin/Settings.html:63
func settingsLink(href string, icon string, title string, description string, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:63
	writesettingsLink(qb422016, href, icon, title, description, ps)
//line views/vadmin/Settings.html:63
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:63
	return qs422016
//line views/vadmin/Settings.html:63
}
