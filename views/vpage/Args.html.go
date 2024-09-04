// Code generated by qtc from "Args.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vpage/Args.html:1
package vpage

//line views/vpage/Args.html:1
import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components/edit"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vpage/Args.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vpage/Args.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vpage/Args.html:13
type Args struct {
	layout.Basic
	URL        string
	Directions string
	Results    *util.FieldDescResults
	Hidden     map[string]string
	Warning    string
}

//line views/vpage/Args.html:22
func (p *Args) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Args.html:22
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vpage/Args.html:24
	if p.Directions == "" {
//line views/vpage/Args.html:24
		qw422016.N().S(`Enter Data`)
//line views/vpage/Args.html:24
	} else {
//line views/vpage/Args.html:24
		qw422016.E().S(p.Directions)
//line views/vpage/Args.html:24
	}
//line views/vpage/Args.html:24
	qw422016.N().S(`</h3>
`)
//line views/vpage/Args.html:26
	onsubmit := ""
	if p.Warning != "" {
		onsubmit = fmt.Sprintf(` onsubmit="return confirm('%s')"`, strings.ReplaceAll(strings.ReplaceAll(p.Warning, "'", "\\'"), "\"", ""))
	}

//line views/vpage/Args.html:30
	qw422016.N().S(`    <form action="`)
//line views/vpage/Args.html:31
	qw422016.E().S(p.URL)
//line views/vpage/Args.html:31
	qw422016.N().S(`" method="get"`)
//line views/vpage/Args.html:31
	qw422016.N().S(onsubmit)
//line views/vpage/Args.html:31
	qw422016.N().S(`">
`)
//line views/vpage/Args.html:32
	for k, v := range p.Hidden {
//line views/vpage/Args.html:32
		qw422016.N().S(`      <input type="hidden" name="`)
//line views/vpage/Args.html:33
		qw422016.E().S(k)
//line views/vpage/Args.html:33
		qw422016.N().S(`" value="`)
//line views/vpage/Args.html:33
		qw422016.E().S(v)
//line views/vpage/Args.html:33
		qw422016.N().S(`" />
`)
//line views/vpage/Args.html:34
	}
//line views/vpage/Args.html:34
	qw422016.N().S(`      <div class="overflow full-width">
        <table class="mt min-200 expanded">
          <tbody>
`)
//line views/vpage/Args.html:38
	for _, arg := range p.Results.FieldDescs {
//line views/vpage/Args.html:40
		v := util.OrDefault(p.Results.Values.GetStringOpt(arg.Key), arg.Default)
		title := arg.Title
		if len(title) > 50 {
			title = title[:50] + "..."
		}

//line views/vpage/Args.html:46
		switch arg.Type {
//line views/vpage/Args.html:47
		case "bool":
//line views/vpage/Args.html:47
			qw422016.N().S(`            `)
//line views/vpage/Args.html:48
			edit.StreamBoolTable(qw422016, arg.Key, title, v == "true", 5, arg.Description)
//line views/vpage/Args.html:48
			qw422016.N().S(`
`)
//line views/vpage/Args.html:49
		case "textarea":
//line views/vpage/Args.html:49
			qw422016.N().S(`            `)
//line views/vpage/Args.html:50
			edit.StreamTextareaTable(qw422016, arg.Key, "", title, 12, v, 5, arg.Description)
//line views/vpage/Args.html:50
			qw422016.N().S(`
`)
//line views/vpage/Args.html:51
		case "number", "int":
//line views/vpage/Args.html:52
			i, _ := strconv.ParseInt(v, 10, 32)

//line views/vpage/Args.html:52
			qw422016.N().S(`            `)
//line views/vpage/Args.html:53
			edit.StreamIntTable(qw422016, arg.Key, "", title, int(i), 5, arg.Description)
//line views/vpage/Args.html:53
			qw422016.N().S(`
`)
//line views/vpage/Args.html:54
		case "float":
//line views/vpage/Args.html:55
			f, _ := strconv.ParseFloat(v, 64)

//line views/vpage/Args.html:55
			qw422016.N().S(`            `)
//line views/vpage/Args.html:56
			edit.StreamFloatTable(qw422016, arg.Key, "", title, f, 5, arg.Description)
//line views/vpage/Args.html:56
			qw422016.N().S(`
`)
//line views/vpage/Args.html:57
		default:
//line views/vpage/Args.html:57
			qw422016.N().S(`            `)
//line views/vpage/Args.html:58
			edit.StreamDatalistTable(qw422016, arg.Key, "", title, v, arg.Choices, nil, 5, arg.Description)
//line views/vpage/Args.html:58
			qw422016.N().S(`
`)
//line views/vpage/Args.html:59
		}
//line views/vpage/Args.html:60
	}
//line views/vpage/Args.html:60
	qw422016.N().S(`          </tbody>
        </table>
      </div>
      <button class="mt" type="submit">Submit</button>
    </form>
  </div>
`)
//line views/vpage/Args.html:67
}

//line views/vpage/Args.html:67
func (p *Args) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vpage/Args.html:67
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vpage/Args.html:67
	p.StreamBody(qw422016, as, ps)
//line views/vpage/Args.html:67
	qt422016.ReleaseWriter(qw422016)
//line views/vpage/Args.html:67
}

//line views/vpage/Args.html:67
func (p *Args) Body(as *app.State, ps *cutil.PageState) string {
//line views/vpage/Args.html:67
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vpage/Args.html:67
	p.WriteBody(qb422016, as, ps)
//line views/vpage/Args.html:67
	qs422016 := string(qb422016.B)
//line views/vpage/Args.html:67
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vpage/Args.html:67
	return qs422016
//line views/vpage/Args.html:67
}
