// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vgraphql/List.html:2
package vgraphql

//line views/vgraphql/List.html:2
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/components"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vgraphql/List.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vgraphql/List.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vgraphql/List.html:9
type List struct {
	layout.Basic
	Keys   []string
	Counts []int
}

//line views/vgraphql/List.html:15
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgraphql/List.html:15
	qw422016.N().S(`
`)
//line views/vgraphql/List.html:16
	titles := as.GraphQL.Titles()

//line views/vgraphql/List.html:16
	qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vgraphql/List.html:18
	components.StreamSVGIcon(qw422016, `graph`, ps)
//line views/vgraphql/List.html:18
	qw422016.N().S(`GraphQL</h3>
`)
//line views/vgraphql/List.html:19
	if len(p.Keys) == 0 {
//line views/vgraphql/List.html:19
		qw422016.N().S(`    <p><em>no schemata available</em></p>
`)
//line views/vgraphql/List.html:21
	} else {
//line views/vgraphql/List.html:21
		qw422016.N().S(`    <div class="overflow full-width">
      <table class="mt">
        <thead>
          <tr>
            <th>Name</th>
            <th class="shrink">Exec Count</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vgraphql/List.html:31
		for idx, k := range p.Keys {
//line views/vgraphql/List.html:31
			qw422016.N().S(`          <tr>
            <td><a href="/graphql/`)
//line views/vgraphql/List.html:33
			qw422016.E().S(k)
//line views/vgraphql/List.html:33
			qw422016.N().S(`">`)
//line views/vgraphql/List.html:33
			qw422016.E().S(titles[k])
//line views/vgraphql/List.html:33
			qw422016.N().S(`</a></td>
            <td>`)
//line views/vgraphql/List.html:34
			qw422016.N().D(p.Counts[idx])
//line views/vgraphql/List.html:34
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vgraphql/List.html:36
		}
//line views/vgraphql/List.html:36
		qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vgraphql/List.html:40
	}
//line views/vgraphql/List.html:40
	qw422016.N().S(`  </div>
`)
//line views/vgraphql/List.html:42
}

//line views/vgraphql/List.html:42
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgraphql/List.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgraphql/List.html:42
	p.StreamBody(qw422016, as, ps)
//line views/vgraphql/List.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/vgraphql/List.html:42
}

//line views/vgraphql/List.html:42
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vgraphql/List.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgraphql/List.html:42
	p.WriteBody(qb422016, as, ps)
//line views/vgraphql/List.html:42
	qs422016 := string(qb422016.B)
//line views/vgraphql/List.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgraphql/List.html:42
	return qs422016
//line views/vgraphql/List.html:42
}
