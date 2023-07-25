// Code generated by qtc from "ServerInfo.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/ServerInfo.html:2
package vadmin

//line views/vadmin/ServerInfo.html:2
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vadmin/ServerInfo.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/ServerInfo.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/ServerInfo.html:9
type ServerInfo struct {
	layout.Basic
	Info *util.DebugInfo
}

//line views/vadmin/ServerInfo.html:14
func (p *ServerInfo) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/ServerInfo.html:14
	qw422016.N().S(`
  `)
//line views/vadmin/ServerInfo.html:15
	streamrenderTags(qw422016, "Server Information", p.Info.ServerTags)
//line views/vadmin/ServerInfo.html:15
	qw422016.N().S(`
  `)
//line views/vadmin/ServerInfo.html:16
	streamrenderTags(qw422016, "Runtime Information", p.Info.RuntimeTags)
//line views/vadmin/ServerInfo.html:16
	qw422016.N().S(`
  `)
//line views/vadmin/ServerInfo.html:17
	streamrenderTags(qw422016, "App Information", p.Info.AppTags)
//line views/vadmin/ServerInfo.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>Go Modules</h3>
    `)
//line views/vadmin/ServerInfo.html:20
	streammoduleTable(qw422016, p.Info.Mods)
//line views/vadmin/ServerInfo.html:20
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/ServerInfo.html:22
}

//line views/vadmin/ServerInfo.html:22
func (p *ServerInfo) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/ServerInfo.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/ServerInfo.html:22
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/ServerInfo.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/ServerInfo.html:22
}

//line views/vadmin/ServerInfo.html:22
func (p *ServerInfo) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/ServerInfo.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/ServerInfo.html:22
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/ServerInfo.html:22
	qs422016 := string(qb422016.B)
//line views/vadmin/ServerInfo.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/ServerInfo.html:22
	return qs422016
//line views/vadmin/ServerInfo.html:22
}

//line views/vadmin/ServerInfo.html:24
func streamrenderTags(qw422016 *qt422016.Writer, title string, tags *util.OrderedMap[string]) {
//line views/vadmin/ServerInfo.html:24
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/ServerInfo.html:26
	qw422016.E().S(title)
//line views/vadmin/ServerInfo.html:26
	qw422016.N().S(`</h3>
    <table class="mt min-200">
      <tbody>
`)
//line views/vadmin/ServerInfo.html:29
	for _, k := range tags.Order {
//line views/vadmin/ServerInfo.html:29
		qw422016.N().S(`      <tr>
        <th class="shrink">`)
//line views/vadmin/ServerInfo.html:31
		qw422016.E().S(k)
//line views/vadmin/ServerInfo.html:31
		qw422016.N().S(`</th>
        <td>`)
//line views/vadmin/ServerInfo.html:32
		qw422016.E().S(tags.GetSimple(k))
//line views/vadmin/ServerInfo.html:32
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vadmin/ServerInfo.html:34
	}
//line views/vadmin/ServerInfo.html:34
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vadmin/ServerInfo.html:38
}

//line views/vadmin/ServerInfo.html:38
func writerenderTags(qq422016 qtio422016.Writer, title string, tags *util.OrderedMap[string]) {
//line views/vadmin/ServerInfo.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/ServerInfo.html:38
	streamrenderTags(qw422016, title, tags)
//line views/vadmin/ServerInfo.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/ServerInfo.html:38
}

//line views/vadmin/ServerInfo.html:38
func renderTags(title string, tags *util.OrderedMap[string]) string {
//line views/vadmin/ServerInfo.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/ServerInfo.html:38
	writerenderTags(qb422016, title, tags)
//line views/vadmin/ServerInfo.html:38
	qs422016 := string(qb422016.B)
//line views/vadmin/ServerInfo.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/ServerInfo.html:38
	return qs422016
//line views/vadmin/ServerInfo.html:38
}
