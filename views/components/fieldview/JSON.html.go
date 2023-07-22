// Code generated by qtc from "JSON.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/fieldview/JSON.html:2
package fieldview

//line views/components/fieldview/JSON.html:2
import (
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
)

//line views/components/fieldview/JSON.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldview/JSON.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldview/JSON.html:7
func StreamJSON(qw422016 *qt422016.Writer, v any) {
//line views/components/fieldview/JSON.html:9
	b, ok := v.([]byte)
	if ok {
		_ = util.FromJSON(b, &v)
	}

//line views/components/fieldview/JSON.html:14
	out, err := cutil.FormatJSON(v)

//line views/components/fieldview/JSON.html:15
	if err == nil {
//line views/components/fieldview/JSON.html:16
		qw422016.N().S(out)
//line views/components/fieldview/JSON.html:17
	} else {
//line views/components/fieldview/JSON.html:18
		qw422016.E().S(err.Error())
//line views/components/fieldview/JSON.html:19
	}
//line views/components/fieldview/JSON.html:20
}

//line views/components/fieldview/JSON.html:20
func WriteJSON(qq422016 qtio422016.Writer, v any) {
//line views/components/fieldview/JSON.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/JSON.html:20
	StreamJSON(qw422016, v)
//line views/components/fieldview/JSON.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/JSON.html:20
}

//line views/components/fieldview/JSON.html:20
func JSON(v any) string {
//line views/components/fieldview/JSON.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/JSON.html:20
	WriteJSON(qb422016, v)
//line views/components/fieldview/JSON.html:20
	qs422016 := string(qb422016.B)
//line views/components/fieldview/JSON.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/JSON.html:20
	return qs422016
//line views/components/fieldview/JSON.html:20
}

//line views/components/fieldview/JSON.html:22
func StreamJSONInline(qw422016 *qt422016.Writer, v any) {
//line views/components/fieldview/JSON.html:24
	b, ok := v.([]byte)
	if ok {
		_ = util.FromJSON(b, &v)
	}

//line views/components/fieldview/JSON.html:29
	qw422016.E().S(util.ToJSON(v))
//line views/components/fieldview/JSON.html:30
}

//line views/components/fieldview/JSON.html:30
func WriteJSONInline(qq422016 qtio422016.Writer, v any) {
//line views/components/fieldview/JSON.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/JSON.html:30
	StreamJSONInline(qw422016, v)
//line views/components/fieldview/JSON.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/JSON.html:30
}

//line views/components/fieldview/JSON.html:30
func JSONInline(v any) string {
//line views/components/fieldview/JSON.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/JSON.html:30
	WriteJSONInline(qb422016, v)
//line views/components/fieldview/JSON.html:30
	qs422016 := string(qb422016.B)
//line views/components/fieldview/JSON.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/JSON.html:30
	return qs422016
//line views/components/fieldview/JSON.html:30
}
