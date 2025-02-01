// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vexec/List.html:1
package vexec

//line views/vexec/List.html:1
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/lib/exec"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vexec/List.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vexec/List.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vexec/List.html:10
type List struct {
	layout.Basic
	Execs exec.Execs
}

//line views/vexec/List.html:15
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexec/List.html:15
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/exec/new" title="start a new process">`)
//line views/vexec/List.html:17
	components.StreamSVGIcon(qw422016, `plus`, ps)
//line views/vexec/List.html:17
	qw422016.N().S(`</a></div>
    <h3>`)
//line views/vexec/List.html:18
	components.StreamSVGIcon(qw422016, `desktop`, ps)
//line views/vexec/List.html:18
	qw422016.N().S(` `)
//line views/vexec/List.html:18
	qw422016.E().S(util.StringPlural(len(p.Execs), "Process"))
//line views/vexec/List.html:18
	qw422016.N().S(`</h3>
    `)
//line views/vexec/List.html:19
	StreamTable(qw422016, p.Execs, as, ps)
//line views/vexec/List.html:19
	qw422016.N().S(`
  </div>
`)
//line views/vexec/List.html:21
}

//line views/vexec/List.html:21
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vexec/List.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vexec/List.html:21
	p.StreamBody(qw422016, as, ps)
//line views/vexec/List.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/vexec/List.html:21
}

//line views/vexec/List.html:21
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vexec/List.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vexec/List.html:21
	p.WriteBody(qb422016, as, ps)
//line views/vexec/List.html:21
	qs422016 := string(qb422016.B)
//line views/vexec/List.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vexec/List.html:21
	return qs422016
//line views/vexec/List.html:21
}
