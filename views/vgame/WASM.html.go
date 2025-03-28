// Code generated by qtc from "WASM.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vgame/WASM.html:1
package vgame

//line views/vgame/WASM.html:1
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/game"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vgame/WASM.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vgame/WASM.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vgame/WASM.html:10
type WASM struct {
	layout.Basic
	Game *game.Game
}

//line views/vgame/WASM.html:15
func (p *WASM) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgame/WASM.html:15
	qw422016.N().S(`
<div class="card">
  <h3>`)
//line views/vgame/WASM.html:17
	components.StreamSVGIcon(qw422016, `gift`, ps)
//line views/vgame/WASM.html:17
	qw422016.N().S(` `)
//line views/vgame/WASM.html:17
	qw422016.E().S(util.AppName)
//line views/vgame/WASM.html:17
	qw422016.N().S(` WASM</h3>
  <em id="load-status">Loading...</em>
</div>
<div class="card">
  <h3>Audits</h3>
  <ul id="audit-log" class="mt"></ul>
</div>
<script>
  function wasmInit(ms) {
    document.getElementById("load-status").innerText = "Loaded in [" + ms + "ms]";
  }
</script>
`)
//line views/vgame/WASM.html:29
	components.StreamWASMScript(qw422016, p.Game)
//line views/vgame/WASM.html:29
	qw422016.N().S(`
`)
//line views/vgame/WASM.html:30
}

//line views/vgame/WASM.html:30
func (p *WASM) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgame/WASM.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgame/WASM.html:30
	p.StreamBody(qw422016, as, ps)
//line views/vgame/WASM.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/vgame/WASM.html:30
}

//line views/vgame/WASM.html:30
func (p *WASM) Body(as *app.State, ps *cutil.PageState) string {
//line views/vgame/WASM.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgame/WASM.html:30
	p.WriteBody(qb422016, as, ps)
//line views/vgame/WASM.html:30
	qs422016 := string(qb422016.B)
//line views/vgame/WASM.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgame/WASM.html:30
	return qs422016
//line views/vgame/WASM.html:30
}
