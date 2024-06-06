// Code generated by qtc from "WASM.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsandbox/WASM.html:2
package vsandbox

//line views/vsandbox/WASM.html:2
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vsandbox/WASM.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsandbox/WASM.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsandbox/WASM.html:10
type WASM struct {
	layout.Basic
}

//line views/vsandbox/WASM.html:14
func (p *WASM) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/WASM.html:14
	qw422016.N().S(`
<div class="card">
  <h3>`)
//line views/vsandbox/WASM.html:16
	components.StreamSVGIcon(qw422016, `gift`, ps)
//line views/vsandbox/WASM.html:16
	qw422016.E().S(util.AppName)
//line views/vsandbox/WASM.html:16
	qw422016.N().S(` WASM</h3>
  <em id="load-status">Loading...</em>
</div>
<div class="card">
  <h3>Audits</h3>
  <ul id="audit-log"></ul>
</div>
<script>
  function wasmInit(ms) {
    document.getElementById("load-status").innerText = "Loaded in [" + ms + "ms]";
  }
</script>
`)
//line views/vsandbox/WASM.html:28
	components.StreamWASMScript(qw422016)
//line views/vsandbox/WASM.html:28
	qw422016.N().S(`
`)
//line views/vsandbox/WASM.html:29
}

//line views/vsandbox/WASM.html:29
func (p *WASM) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/WASM.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsandbox/WASM.html:29
	p.StreamBody(qw422016, as, ps)
//line views/vsandbox/WASM.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/vsandbox/WASM.html:29
}

//line views/vsandbox/WASM.html:29
func (p *WASM) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsandbox/WASM.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsandbox/WASM.html:29
	p.WriteBody(qb422016, as, ps)
//line views/vsandbox/WASM.html:29
	qs422016 := string(qb422016.B)
//line views/vsandbox/WASM.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsandbox/WASM.html:29
	return qs422016
//line views/vsandbox/WASM.html:29
}
