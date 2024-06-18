// Code generated by qtc from "File.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/edit/File.html:1
package edit

//line views/components/edit/File.html:1
import "github.com/kyleu/solitaire/views/components"

//line views/components/edit/File.html:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/File.html:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/File.html:3
func StreamFile(qw422016 *qt422016.Writer, key string, id string, label string, value string, placeholder ...string) {
//line views/components/edit/File.html:4
	if id == "" {
//line views/components/edit/File.html:4
		qw422016.N().S(`<label><input type="file" name="`)
//line views/components/edit/File.html:5
		qw422016.E().S(key)
//line views/components/edit/File.html:5
		qw422016.N().S(`" value="`)
//line views/components/edit/File.html:5
		qw422016.E().S(value)
//line views/components/edit/File.html:5
		qw422016.N().S(`"`)
//line views/components/edit/File.html:5
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/File.html:5
		qw422016.N().S(`/>`)
//line views/components/edit/File.html:5
		qw422016.E().S(label)
//line views/components/edit/File.html:5
		qw422016.N().S(`</label>`)
//line views/components/edit/File.html:6
	} else {
//line views/components/edit/File.html:6
		qw422016.N().S(`<label><input id="`)
//line views/components/edit/File.html:7
		qw422016.E().S(id)
//line views/components/edit/File.html:7
		qw422016.N().S(`" type="file" name="`)
//line views/components/edit/File.html:7
		qw422016.E().S(key)
//line views/components/edit/File.html:7
		qw422016.N().S(`" value="`)
//line views/components/edit/File.html:7
		qw422016.E().S(value)
//line views/components/edit/File.html:7
		qw422016.N().S(`"`)
//line views/components/edit/File.html:7
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/File.html:7
		qw422016.N().S(`/>`)
//line views/components/edit/File.html:7
		qw422016.E().S(label)
//line views/components/edit/File.html:7
		qw422016.N().S(`</label>`)
//line views/components/edit/File.html:8
	}
//line views/components/edit/File.html:9
}

//line views/components/edit/File.html:9
func WriteFile(qq422016 qtio422016.Writer, key string, id string, label string, value string, placeholder ...string) {
//line views/components/edit/File.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/File.html:9
	StreamFile(qw422016, key, id, label, value, placeholder...)
//line views/components/edit/File.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/File.html:9
}

//line views/components/edit/File.html:9
func File(key string, id string, label string, value string, placeholder ...string) string {
//line views/components/edit/File.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/File.html:9
	WriteFile(qb422016, key, id, label, value, placeholder...)
//line views/components/edit/File.html:9
	qs422016 := string(qb422016.B)
//line views/components/edit/File.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/File.html:9
	return qs422016
//line views/components/edit/File.html:9
}

//line views/components/edit/File.html:11
func StreamFileTable(qw422016 *qt422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/edit/File.html:11
	qw422016.N().S(`<tr><th class="shrink"><label for="`)
//line views/components/edit/File.html:13
	qw422016.E().S(id)
//line views/components/edit/File.html:13
	qw422016.N().S(`">`)
//line views/components/edit/File.html:13
	qw422016.E().S(title)
//line views/components/edit/File.html:13
	qw422016.N().S(`</label></th><td>`)
//line views/components/edit/File.html:15
	StreamFile(qw422016, key, id, label, value)
//line views/components/edit/File.html:15
	qw422016.N().S(`</td></tr>`)
//line views/components/edit/File.html:18
}

//line views/components/edit/File.html:18
func WriteFileTable(qq422016 qtio422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/edit/File.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/File.html:18
	StreamFileTable(qw422016, key, id, title, label, value)
//line views/components/edit/File.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/File.html:18
}

//line views/components/edit/File.html:18
func FileTable(key string, id string, title string, label string, value string) string {
//line views/components/edit/File.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/File.html:18
	WriteFileTable(qb422016, key, id, title, label, value)
//line views/components/edit/File.html:18
	qs422016 := string(qb422016.B)
//line views/components/edit/File.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/File.html:18
	return qs422016
//line views/components/edit/File.html:18
}

//line views/components/edit/File.html:20
func StreamFileMultiple(qw422016 *qt422016.Writer, key string, id string, label string, value string, placeholder ...string) {
//line views/components/edit/File.html:21
	if id == "" {
//line views/components/edit/File.html:21
		qw422016.N().S(`<label><input type="file" name="`)
//line views/components/edit/File.html:22
		qw422016.E().S(key)
//line views/components/edit/File.html:22
		qw422016.N().S(`" value="`)
//line views/components/edit/File.html:22
		qw422016.E().S(value)
//line views/components/edit/File.html:22
		qw422016.N().S(`" multiple="multiple"`)
//line views/components/edit/File.html:22
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/File.html:22
		qw422016.N().S(`/>`)
//line views/components/edit/File.html:22
		qw422016.E().S(label)
//line views/components/edit/File.html:22
		qw422016.N().S(`</label>`)
//line views/components/edit/File.html:23
	} else {
//line views/components/edit/File.html:23
		qw422016.N().S(`<label><input id="`)
//line views/components/edit/File.html:24
		qw422016.E().S(id)
//line views/components/edit/File.html:24
		qw422016.N().S(`" type="file" name="`)
//line views/components/edit/File.html:24
		qw422016.E().S(key)
//line views/components/edit/File.html:24
		qw422016.N().S(`" value="`)
//line views/components/edit/File.html:24
		qw422016.E().S(value)
//line views/components/edit/File.html:24
		qw422016.N().S(`" multiple="multiple"`)
//line views/components/edit/File.html:24
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/File.html:24
		qw422016.N().S(`/>`)
//line views/components/edit/File.html:24
		qw422016.E().S(label)
//line views/components/edit/File.html:24
		qw422016.N().S(`</label>`)
//line views/components/edit/File.html:25
	}
//line views/components/edit/File.html:26
}

//line views/components/edit/File.html:26
func WriteFileMultiple(qq422016 qtio422016.Writer, key string, id string, label string, value string, placeholder ...string) {
//line views/components/edit/File.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/File.html:26
	StreamFileMultiple(qw422016, key, id, label, value, placeholder...)
//line views/components/edit/File.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/File.html:26
}

//line views/components/edit/File.html:26
func FileMultiple(key string, id string, label string, value string, placeholder ...string) string {
//line views/components/edit/File.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/File.html:26
	WriteFileMultiple(qb422016, key, id, label, value, placeholder...)
//line views/components/edit/File.html:26
	qs422016 := string(qb422016.B)
//line views/components/edit/File.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/File.html:26
	return qs422016
//line views/components/edit/File.html:26
}

//line views/components/edit/File.html:28
func StreamFileMultipleTable(qw422016 *qt422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/edit/File.html:28
	qw422016.N().S(`<tr><th class="shrink"><label for="`)
//line views/components/edit/File.html:30
	qw422016.E().S(id)
//line views/components/edit/File.html:30
	qw422016.N().S(`">`)
//line views/components/edit/File.html:30
	qw422016.E().S(title)
//line views/components/edit/File.html:30
	qw422016.N().S(`</label></th><td>`)
//line views/components/edit/File.html:32
	StreamFileMultiple(qw422016, key, id, label, value)
//line views/components/edit/File.html:32
	qw422016.N().S(`</td></tr>`)
//line views/components/edit/File.html:35
}

//line views/components/edit/File.html:35
func WriteFileMultipleTable(qq422016 qtio422016.Writer, key string, id string, title string, label string, value string) {
//line views/components/edit/File.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/File.html:35
	StreamFileMultipleTable(qw422016, key, id, title, label, value)
//line views/components/edit/File.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/File.html:35
}

//line views/components/edit/File.html:35
func FileMultipleTable(key string, id string, title string, label string, value string) string {
//line views/components/edit/File.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/File.html:35
	WriteFileMultipleTable(qb422016, key, id, title, label, value)
//line views/components/edit/File.html:35
	qs422016 := string(qb422016.B)
//line views/components/edit/File.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/File.html:35
	return qs422016
//line views/components/edit/File.html:35
}
