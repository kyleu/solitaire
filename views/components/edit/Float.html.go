// Code generated by qtc from "Float.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/edit/Float.html:2
package edit

//line views/components/edit/Float.html:2
import (
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/components"
)

//line views/components/edit/Float.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Float.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Float.html:7
func StreamFloat(qw422016 *qt422016.Writer, key string, id string, value any, placeholder ...string) {
//line views/components/edit/Float.html:8
	if id == "" {
//line views/components/edit/Float.html:8
		qw422016.N().S(`<input name="`)
//line views/components/edit/Float.html:9
		qw422016.E().S(key)
//line views/components/edit/Float.html:9
		qw422016.N().S(`" type="number" step="any" value="`)
//line views/components/edit/Float.html:9
		qw422016.E().V(value)
//line views/components/edit/Float.html:9
		qw422016.N().S(`"`)
//line views/components/edit/Float.html:9
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Float.html:9
		qw422016.N().S(`/>`)
//line views/components/edit/Float.html:10
	} else {
//line views/components/edit/Float.html:10
		qw422016.N().S(`<input id="`)
//line views/components/edit/Float.html:11
		qw422016.E().S(id)
//line views/components/edit/Float.html:11
		qw422016.N().S(`" name="`)
//line views/components/edit/Float.html:11
		qw422016.E().S(key)
//line views/components/edit/Float.html:11
		qw422016.N().S(`" type="number" step="any" value="`)
//line views/components/edit/Float.html:11
		qw422016.E().V(value)
//line views/components/edit/Float.html:11
		qw422016.N().S(`"`)
//line views/components/edit/Float.html:11
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Float.html:11
		qw422016.N().S(`/>`)
//line views/components/edit/Float.html:12
	}
//line views/components/edit/Float.html:13
}

//line views/components/edit/Float.html:13
func WriteFloat(qq422016 qtio422016.Writer, key string, id string, value any, placeholder ...string) {
//line views/components/edit/Float.html:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Float.html:13
	StreamFloat(qw422016, key, id, value, placeholder...)
//line views/components/edit/Float.html:13
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Float.html:13
}

//line views/components/edit/Float.html:13
func Float(key string, id string, value any, placeholder ...string) string {
//line views/components/edit/Float.html:13
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Float.html:13
	WriteFloat(qb422016, key, id, value, placeholder...)
//line views/components/edit/Float.html:13
	qs422016 := string(qb422016.B)
//line views/components/edit/Float.html:13
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Float.html:13
	return qs422016
//line views/components/edit/Float.html:13
}

//line views/components/edit/Float.html:15
func StreamFloatVertical(qw422016 *qt422016.Writer, key string, id string, title string, value float64, indent int, help ...string) {
//line views/components/edit/Float.html:16
	id = cutil.CleanID(key, id)

//line views/components/edit/Float.html:16
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Float.html:18
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Float.html:18
	qw422016.N().S(`<label for="`)
//line views/components/edit/Float.html:19
	qw422016.E().S(id)
//line views/components/edit/Float.html:19
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/Float.html:19
	qw422016.E().S(title)
//line views/components/edit/Float.html:19
	qw422016.N().S(`</em></label>`)
//line views/components/edit/Float.html:20
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Float.html:20
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Float.html:21
	StreamFloat(qw422016, key, id, value, help...)
//line views/components/edit/Float.html:21
	qw422016.N().S(`</div>`)
//line views/components/edit/Float.html:22
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Float.html:22
	qw422016.N().S(`</div>`)
//line views/components/edit/Float.html:24
}

//line views/components/edit/Float.html:24
func WriteFloatVertical(qq422016 qtio422016.Writer, key string, id string, title string, value float64, indent int, help ...string) {
//line views/components/edit/Float.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Float.html:24
	StreamFloatVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/Float.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Float.html:24
}

//line views/components/edit/Float.html:24
func FloatVertical(key string, id string, title string, value float64, indent int, help ...string) string {
//line views/components/edit/Float.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Float.html:24
	WriteFloatVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/Float.html:24
	qs422016 := string(qb422016.B)
//line views/components/edit/Float.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Float.html:24
	return qs422016
//line views/components/edit/Float.html:24
}

//line views/components/edit/Float.html:26
func StreamFloatTable(qw422016 *qt422016.Writer, key string, id string, title string, value float64, indent int, help ...string) {
//line views/components/edit/Float.html:27
	id = cutil.CleanID(key, id)

//line views/components/edit/Float.html:27
	qw422016.N().S(`<tr>`)
//line views/components/edit/Float.html:29
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Float.html:29
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/Float.html:30
	qw422016.E().S(id)
//line views/components/edit/Float.html:30
	qw422016.N().S(`"`)
//line views/components/edit/Float.html:30
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Float.html:30
	qw422016.N().S(`>`)
//line views/components/edit/Float.html:30
	qw422016.E().S(title)
//line views/components/edit/Float.html:30
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Float.html:31
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Float.html:31
	qw422016.N().S(`<td>`)
//line views/components/edit/Float.html:32
	StreamFloat(qw422016, key, id, value, help...)
//line views/components/edit/Float.html:32
	qw422016.N().S(`</td>`)
//line views/components/edit/Float.html:33
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Float.html:33
	qw422016.N().S(`</tr>`)
//line views/components/edit/Float.html:35
}

//line views/components/edit/Float.html:35
func WriteFloatTable(qq422016 qtio422016.Writer, key string, id string, title string, value float64, indent int, help ...string) {
//line views/components/edit/Float.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Float.html:35
	StreamFloatTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/Float.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Float.html:35
}

//line views/components/edit/Float.html:35
func FloatTable(key string, id string, title string, value float64, indent int, help ...string) string {
//line views/components/edit/Float.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Float.html:35
	WriteFloatTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/Float.html:35
	qs422016 := string(qb422016.B)
//line views/components/edit/Float.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Float.html:35
	return qs422016
//line views/components/edit/Float.html:35
}
