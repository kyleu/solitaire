// Code generated by qtc from "HTML.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vgame/HTML.html:1
package vgame

//line views/vgame/HTML.html:1
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vgame/HTML.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vgame/HTML.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vgame/HTML.html:9
type HTML struct {
	layout.Basic
}

//line views/vgame/HTML.html:13
func (p *HTML) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgame/HTML.html:13
	qw422016.N().S(`
<div class="card">
  <h3>`)
//line views/vgame/HTML.html:15
	components.StreamSVGRefIcon(qw422016, `gift`, ps)
//line views/vgame/HTML.html:15
	qw422016.E().S(util.AppName)
//line views/vgame/HTML.html:15
	qw422016.N().S(` HTML</h3>
</div>
<div class="card">
  <h3>Log</h3>
  <ul id="audit-log"></ul>
</div>
`)
//line views/vgame/HTML.html:21
}

//line views/vgame/HTML.html:21
func (p *HTML) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgame/HTML.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgame/HTML.html:21
	p.StreamBody(qw422016, as, ps)
//line views/vgame/HTML.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/vgame/HTML.html:21
}

//line views/vgame/HTML.html:21
func (p *HTML) Body(as *app.State, ps *cutil.PageState) string {
//line views/vgame/HTML.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgame/HTML.html:21
	p.WriteBody(qb422016, as, ps)
//line views/vgame/HTML.html:21
	qs422016 := string(qb422016.B)
//line views/vgame/HTML.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgame/HTML.html:21
	return qs422016
//line views/vgame/HTML.html:21
}
