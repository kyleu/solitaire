// Code generated by qtc from "Load.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vpage/Load.html:1
package vpage

//line views/vpage/Load.html:1
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vpage/Load.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vpage/Load.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vpage/Load.html:7
type Load struct {
	layout.Basic
	URL              string
	Title            string
	Message          string
	HideInstructions bool
}

//line views/vpage/Load.html:15
func (p *Load) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Load.html:15
	qw422016.N().S(`
`)
//line views/vpage/Load.html:17
	if p.Message == "" {
		p.Message = "Please wait as your request is processed..."
	}

//line views/vpage/Load.html:20
	qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vpage/Load.html:22
	qw422016.E().S(p.Title)
//line views/vpage/Load.html:22
	qw422016.N().S(`</h3>
    <p>`)
//line views/vpage/Load.html:23
	qw422016.E().S(p.Message)
//line views/vpage/Load.html:23
	qw422016.N().S(`</p>
`)
//line views/vpage/Load.html:24
	if !p.HideInstructions {
//line views/vpage/Load.html:24
		qw422016.N().S(`    <div class="mt"><em>Please avoid refreshing the browser or navigating away, your page is loading</em></div>
`)
//line views/vpage/Load.html:26
	}
//line views/vpage/Load.html:26
	qw422016.N().S(`  </div>
  <meta http-equiv="refresh" content="0; url=`)
//line views/vpage/Load.html:28
	qw422016.E().S(p.URL)
//line views/vpage/Load.html:28
	qw422016.N().S(`">
`)
//line views/vpage/Load.html:29
}

//line views/vpage/Load.html:29
func (p *Load) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Load.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vpage/Load.html:29
	p.StreamBody(qw422016, as, ps)
//line views/vpage/Load.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/vpage/Load.html:29
}

//line views/vpage/Load.html:29
func (p *Load) Body(as *app.State, ps *cutil.PageState) string {
//line views/vpage/Load.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vpage/Load.html:29
	p.WriteBody(qb422016, as, ps)
//line views/vpage/Load.html:29
	qs422016 := string(qb422016.B)
//line views/vpage/Load.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vpage/Load.html:29
	return qs422016
//line views/vpage/Load.html:29
}
