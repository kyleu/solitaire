// Code generated by qtc from "Index.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vgame/Index.html:1
package vgame

//line views/vgame/Index.html:1
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vgame/Index.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vgame/Index.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vgame/Index.html:7
type Index struct {
	layout.Basic
}

//line views/vgame/Index.html:11
func (p *Index) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgame/Index.html:11
	qw422016.N().S(`
  <div class="card">
    <h3>Game Tests</h3>
    <ul class="mt">
      <li><a href="/game/test/html">HTML</a></li>
      <li><a href="/game/test/json">JSON</a></li>
      <li><a href="/game/test/wasm">WASM</a></li>
    </ul>
  </div>
`)
//line views/vgame/Index.html:20
}

//line views/vgame/Index.html:20
func (p *Index) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgame/Index.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgame/Index.html:20
	p.StreamBody(qw422016, as, ps)
//line views/vgame/Index.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/vgame/Index.html:20
}

//line views/vgame/Index.html:20
func (p *Index) Body(as *app.State, ps *cutil.PageState) string {
//line views/vgame/Index.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgame/Index.html:20
	p.WriteBody(qb422016, as, ps)
//line views/vgame/Index.html:20
	qs422016 := string(qb422016.B)
//line views/vgame/Index.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgame/Index.html:20
	return qs422016
//line views/vgame/Index.html:20
}
