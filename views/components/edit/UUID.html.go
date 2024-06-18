// Code generated by qtc from "UUID.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/edit/UUID.html:1
package edit

//line views/components/edit/UUID.html:1
import (
	"github.com/google/uuid"

	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/components"
)

//line views/components/edit/UUID.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/UUID.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/UUID.html:8
func StreamUUID(qw422016 *qt422016.Writer, key string, id string, value *uuid.UUID, placeholder ...string) {
//line views/components/edit/UUID.html:10
	v := ""
	if value != nil {
		v = value.String()
	}

//line views/components/edit/UUID.html:15
	StreamString(qw422016, key, id, v, placeholder...)
//line views/components/edit/UUID.html:16
}

//line views/components/edit/UUID.html:16
func WriteUUID(qq422016 qtio422016.Writer, key string, id string, value *uuid.UUID, placeholder ...string) {
//line views/components/edit/UUID.html:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/UUID.html:16
	StreamUUID(qw422016, key, id, value, placeholder...)
//line views/components/edit/UUID.html:16
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/UUID.html:16
}

//line views/components/edit/UUID.html:16
func UUID(key string, id string, value *uuid.UUID, placeholder ...string) string {
//line views/components/edit/UUID.html:16
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/UUID.html:16
	WriteUUID(qb422016, key, id, value, placeholder...)
//line views/components/edit/UUID.html:16
	qs422016 := string(qb422016.B)
//line views/components/edit/UUID.html:16
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/UUID.html:16
	return qs422016
//line views/components/edit/UUID.html:16
}

//line views/components/edit/UUID.html:18
func StreamUUIDVertical(qw422016 *qt422016.Writer, key string, id string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/edit/UUID.html:19
	id = cutil.CleanID(key, id)

//line views/components/edit/UUID.html:19
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/UUID.html:21
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/UUID.html:21
	qw422016.N().S(`<label for="`)
//line views/components/edit/UUID.html:22
	qw422016.E().S(id)
//line views/components/edit/UUID.html:22
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/UUID.html:22
	qw422016.E().S(title)
//line views/components/edit/UUID.html:22
	qw422016.N().S(`</em></label>`)
//line views/components/edit/UUID.html:23
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/UUID.html:23
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/UUID.html:24
	StreamUUID(qw422016, key, id, value, help...)
//line views/components/edit/UUID.html:24
	qw422016.N().S(`</div>`)
//line views/components/edit/UUID.html:25
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/UUID.html:25
	qw422016.N().S(`</div>`)
//line views/components/edit/UUID.html:27
}

//line views/components/edit/UUID.html:27
func WriteUUIDVertical(qq422016 qtio422016.Writer, key string, id string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/edit/UUID.html:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/UUID.html:27
	StreamUUIDVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/UUID.html:27
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/UUID.html:27
}

//line views/components/edit/UUID.html:27
func UUIDVertical(key string, id string, title string, value *uuid.UUID, indent int, help ...string) string {
//line views/components/edit/UUID.html:27
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/UUID.html:27
	WriteUUIDVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/UUID.html:27
	qs422016 := string(qb422016.B)
//line views/components/edit/UUID.html:27
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/UUID.html:27
	return qs422016
//line views/components/edit/UUID.html:27
}

//line views/components/edit/UUID.html:29
func StreamUUIDTable(qw422016 *qt422016.Writer, key string, id string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/edit/UUID.html:30
	id = cutil.CleanID(key, id)

//line views/components/edit/UUID.html:30
	qw422016.N().S(`<tr>`)
//line views/components/edit/UUID.html:32
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/UUID.html:32
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/UUID.html:33
	qw422016.E().S(id)
//line views/components/edit/UUID.html:33
	qw422016.N().S(`"`)
//line views/components/edit/UUID.html:33
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/UUID.html:33
	qw422016.N().S(`>`)
//line views/components/edit/UUID.html:33
	qw422016.E().S(title)
//line views/components/edit/UUID.html:33
	qw422016.N().S(`</label></th>`)
//line views/components/edit/UUID.html:34
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/UUID.html:34
	qw422016.N().S(`<td>`)
//line views/components/edit/UUID.html:35
	StreamUUID(qw422016, key, id, value, help...)
//line views/components/edit/UUID.html:35
	qw422016.N().S(`</td>`)
//line views/components/edit/UUID.html:36
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/UUID.html:36
	qw422016.N().S(`</tr>`)
//line views/components/edit/UUID.html:38
}

//line views/components/edit/UUID.html:38
func WriteUUIDTable(qq422016 qtio422016.Writer, key string, id string, title string, value *uuid.UUID, indent int, help ...string) {
//line views/components/edit/UUID.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/UUID.html:38
	StreamUUIDTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/UUID.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/UUID.html:38
}

//line views/components/edit/UUID.html:38
func UUIDTable(key string, id string, title string, value *uuid.UUID, indent int, help ...string) string {
//line views/components/edit/UUID.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/UUID.html:38
	WriteUUIDTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/UUID.html:38
	qs422016 := string(qb422016.B)
//line views/components/edit/UUID.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/UUID.html:38
	return qs422016
//line views/components/edit/UUID.html:38
}
