// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vfile/List.html:1
package vfile

//line views/vfile/List.html:1
import (
	"path/filepath"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/views/components"
)

//line views/vfile/List.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/List.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/List.html:10
func StreamList(qw422016 *qt422016.Writer, path []string, files filesystem.FileInfos, fl filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vfile/List.html:10
	qw422016.N().S(`
  <h3><a href="`)
//line views/vfile/List.html:11
	qw422016.E().S(urlPrefix)
//line views/vfile/List.html:11
	qw422016.N().S(`">.</a>`)
//line views/vfile/List.html:11
	for idx, p := range path {
//line views/vfile/List.html:11
		qw422016.N().S(`/<a href="`)
//line views/vfile/List.html:11
		qw422016.E().S(urlPrefix)
//line views/vfile/List.html:11
		qw422016.N().S(`/`)
//line views/vfile/List.html:11
		qw422016.E().S(filepath.Join(path[:idx+1]...))
//line views/vfile/List.html:11
		qw422016.N().S(`">`)
//line views/vfile/List.html:11
		qw422016.E().S(p)
//line views/vfile/List.html:11
		qw422016.N().S(`</a>`)
//line views/vfile/List.html:11
	}
//line views/vfile/List.html:11
	qw422016.N().S(`</h3>
  <div class="mt">
`)
//line views/vfile/List.html:13
	for _, f := range files {
//line views/vfile/List.html:15
		icon := "file"
		if f.IsDir {
			icon = "folder"
		}
		x := []string{urlPrefix}
		x = append(x, path...)
		x = append(x, f.Name)
		u := filepath.Join(x...)

//line views/vfile/List.html:23
		qw422016.N().S(`    <div><a href="`)
//line views/vfile/List.html:24
		qw422016.E().S(u)
//line views/vfile/List.html:24
		qw422016.N().S(`">`)
//line views/vfile/List.html:24
		components.StreamSVGInline(qw422016, icon, 16, ps)
//line views/vfile/List.html:24
		qw422016.N().S(` `)
//line views/vfile/List.html:24
		qw422016.E().S(f.Name)
//line views/vfile/List.html:24
		qw422016.N().S(`</a></div>
`)
//line views/vfile/List.html:25
	}
//line views/vfile/List.html:25
	qw422016.N().S(`  </div>
`)
//line views/vfile/List.html:27
}

//line views/vfile/List.html:27
func WriteList(qq422016 qtio422016.Writer, path []string, files filesystem.FileInfos, fl filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) {
//line views/vfile/List.html:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/List.html:27
	StreamList(qw422016, path, files, fl, urlPrefix, as, ps)
//line views/vfile/List.html:27
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/List.html:27
}

//line views/vfile/List.html:27
func List(path []string, files filesystem.FileInfos, fl filesystem.FileLoader, urlPrefix string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/List.html:27
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/List.html:27
	WriteList(qb422016, path, files, fl, urlPrefix, as, ps)
//line views/vfile/List.html:27
	qs422016 := string(qb422016.B)
//line views/vfile/List.html:27
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/List.html:27
	return qs422016
//line views/vfile/List.html:27
}
