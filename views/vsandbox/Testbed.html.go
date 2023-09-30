// Code generated by qtc from "Testbed.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsandbox/Testbed.html:1
package vsandbox

//line views/vsandbox/Testbed.html:1
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vsandbox/Testbed.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsandbox/Testbed.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsandbox/Testbed.html:8
type Testbed struct{ layout.Basic }

//line views/vsandbox/Testbed.html:10
func (p *Testbed) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:10
	qw422016.N().S(`
`)
//line views/vsandbox/Testbed.html:11
	streamicons(qw422016, as, ps)
//line views/vsandbox/Testbed.html:11
	qw422016.N().S(`
`)
//line views/vsandbox/Testbed.html:12
	streamtabs(qw422016, as, ps)
//line views/vsandbox/Testbed.html:12
	qw422016.N().S(`
`)
//line views/vsandbox/Testbed.html:13
	streamaccordion(qw422016, as, ps)
//line views/vsandbox/Testbed.html:13
	qw422016.N().S(`
`)
//line views/vsandbox/Testbed.html:14
	streammodal(qw422016, as, ps)
//line views/vsandbox/Testbed.html:14
	qw422016.N().S(`
`)
//line views/vsandbox/Testbed.html:15
}

//line views/vsandbox/Testbed.html:15
func (p *Testbed) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsandbox/Testbed.html:15
	p.StreamBody(qw422016, as, ps)
//line views/vsandbox/Testbed.html:15
	qt422016.ReleaseWriter(qw422016)
//line views/vsandbox/Testbed.html:15
}

//line views/vsandbox/Testbed.html:15
func (p *Testbed) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsandbox/Testbed.html:15
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsandbox/Testbed.html:15
	p.WriteBody(qb422016, as, ps)
//line views/vsandbox/Testbed.html:15
	qs422016 := string(qb422016.B)
//line views/vsandbox/Testbed.html:15
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsandbox/Testbed.html:15
	return qs422016
//line views/vsandbox/Testbed.html:15
}

//line views/vsandbox/Testbed.html:17
func streamicons(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:17
	qw422016.N().S(`  <div class="card">
    <h3>SVG Icons</h3>
    `)
//line views/vsandbox/Testbed.html:20
	components.StreamIconGallery(qw422016, as, ps)
//line views/vsandbox/Testbed.html:20
	qw422016.N().S(`
  </div>
`)
//line views/vsandbox/Testbed.html:22
}

//line views/vsandbox/Testbed.html:22
func writeicons(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsandbox/Testbed.html:22
	streamicons(qw422016, as, ps)
//line views/vsandbox/Testbed.html:22
	qt422016.ReleaseWriter(qw422016)
//line views/vsandbox/Testbed.html:22
}

//line views/vsandbox/Testbed.html:22
func icons(as *app.State, ps *cutil.PageState) string {
//line views/vsandbox/Testbed.html:22
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsandbox/Testbed.html:22
	writeicons(qb422016, as, ps)
//line views/vsandbox/Testbed.html:22
	qs422016 := string(qb422016.B)
//line views/vsandbox/Testbed.html:22
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsandbox/Testbed.html:22
	return qs422016
//line views/vsandbox/Testbed.html:22
}

//line views/vsandbox/Testbed.html:24
func streamtabs(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:24
	qw422016.N().S(`  <div class="card">
    <h3>Tabs</h3>
    <div class="tabs">
`)
//line views/vsandbox/Testbed.html:28
	for _, o := range []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"} {
//line views/vsandbox/Testbed.html:28
		qw422016.N().S(`      <input name="type" type="radio" id="tab-`)
//line views/vsandbox/Testbed.html:29
		qw422016.E().S(o)
//line views/vsandbox/Testbed.html:29
		qw422016.N().S(`" class="input"/>
      <label for="tab-`)
//line views/vsandbox/Testbed.html:30
		qw422016.E().S(o)
//line views/vsandbox/Testbed.html:30
		qw422016.N().S(`" class="label">`)
//line views/vsandbox/Testbed.html:30
		qw422016.E().S(o)
//line views/vsandbox/Testbed.html:30
		qw422016.N().S(`</label>
      <div class="panel"><p>This is a tab named `)
//line views/vsandbox/Testbed.html:31
		qw422016.E().S(o)
//line views/vsandbox/Testbed.html:31
		qw422016.N().S(`</p></div>
`)
//line views/vsandbox/Testbed.html:32
	}
//line views/vsandbox/Testbed.html:32
	qw422016.N().S(`    </div>
  </div>
`)
//line views/vsandbox/Testbed.html:35
}

//line views/vsandbox/Testbed.html:35
func writetabs(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsandbox/Testbed.html:35
	streamtabs(qw422016, as, ps)
//line views/vsandbox/Testbed.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/vsandbox/Testbed.html:35
}

//line views/vsandbox/Testbed.html:35
func tabs(as *app.State, ps *cutil.PageState) string {
//line views/vsandbox/Testbed.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsandbox/Testbed.html:35
	writetabs(qb422016, as, ps)
//line views/vsandbox/Testbed.html:35
	qs422016 := string(qb422016.B)
//line views/vsandbox/Testbed.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsandbox/Testbed.html:35
	return qs422016
//line views/vsandbox/Testbed.html:35
}

//line views/vsandbox/Testbed.html:37
func streamaccordion(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:37
	qw422016.N().S(`  <div class="card">
    <h3>Accordion</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-a" type="checkbox" hidden />
        <label for="accordion-a">`)
//line views/vsandbox/Testbed.html:43
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsandbox/Testbed.html:43
	qw422016.N().S(` Option A</label>
        <div class="bd"><div><div>Option A!</div></div></div>
      </li>
      <li>
        <input id="accordion-b" type="checkbox" hidden />
        <label for="accordion-b">`)
//line views/vsandbox/Testbed.html:48
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsandbox/Testbed.html:48
	qw422016.N().S(` Option B</label>
        <div class="bd"><div><div>Option B!</div></div></div>
      </li>
      <li>
        <input id="accordion-c" type="checkbox" hidden />
        <label for="accordion-c">`)
//line views/vsandbox/Testbed.html:53
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsandbox/Testbed.html:53
	qw422016.N().S(` Option C (not animated)</label>
        <div class="bd-no-animation">Option C!</div>
      </li>
    </ul>
  </div>
`)
//line views/vsandbox/Testbed.html:58
}

//line views/vsandbox/Testbed.html:58
func writeaccordion(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:58
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsandbox/Testbed.html:58
	streamaccordion(qw422016, as, ps)
//line views/vsandbox/Testbed.html:58
	qt422016.ReleaseWriter(qw422016)
//line views/vsandbox/Testbed.html:58
}

//line views/vsandbox/Testbed.html:58
func accordion(as *app.State, ps *cutil.PageState) string {
//line views/vsandbox/Testbed.html:58
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsandbox/Testbed.html:58
	writeaccordion(qb422016, as, ps)
//line views/vsandbox/Testbed.html:58
	qs422016 := string(qb422016.B)
//line views/vsandbox/Testbed.html:58
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsandbox/Testbed.html:58
	return qs422016
//line views/vsandbox/Testbed.html:58
}

//line views/vsandbox/Testbed.html:60
func streammodal(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:60
	qw422016.N().S(`  <div class="card">
    <h3>Modal</h3>
    <div class="mt"><a href="#modal-x"><button>Open modal</button></a></div>
  </div>
  <div id="modal-x" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Modal</h2>
      </div>
      <div class="modal-body">
        Here's a modal body!
      </div>
    </div>
  </div>
`)
//line views/vsandbox/Testbed.html:77
}

//line views/vsandbox/Testbed.html:77
func writemodal(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsandbox/Testbed.html:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsandbox/Testbed.html:77
	streammodal(qw422016, as, ps)
//line views/vsandbox/Testbed.html:77
	qt422016.ReleaseWriter(qw422016)
//line views/vsandbox/Testbed.html:77
}

//line views/vsandbox/Testbed.html:77
func modal(as *app.State, ps *cutil.PageState) string {
//line views/vsandbox/Testbed.html:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsandbox/Testbed.html:77
	writemodal(qb422016, as, ps)
//line views/vsandbox/Testbed.html:77
	qs422016 := string(qb422016.B)
//line views/vsandbox/Testbed.html:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsandbox/Testbed.html:77
	return qs422016
//line views/vsandbox/Testbed.html:77
}
