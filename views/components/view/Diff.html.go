// Code generated by qtc from "Diff.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/Diff.html:1
package view

//line views/components/view/Diff.html:1
import (
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/components"
)

//line views/components/view/Diff.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Diff.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Diff.html:9
func StreamDiffs(qw422016 *qt422016.Writer, value util.Diffs) {
//line views/components/view/Diff.html:10
	if len(value) == 0 {
//line views/components/view/Diff.html:10
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/view/Diff.html:12
	} else {
//line views/components/view/Diff.html:12
		qw422016.N().S(`<div class="overflow full-width"><table class="expanded"><thead><tr><th>Path</th><th>Old</th><th></th><th>New</th></tr></thead><tbody>`)
//line views/components/view/Diff.html:24
		for _, d := range value {
//line views/components/view/Diff.html:24
			qw422016.N().S(`<tr><td style="width: 30%;"><code>`)
//line views/components/view/Diff.html:26
			qw422016.E().S(d.Path)
//line views/components/view/Diff.html:26
			qw422016.N().S(`</code></td><td style="width: 30%;"><code><em>`)
//line views/components/view/Diff.html:27
			qw422016.E().S(d.Old)
//line views/components/view/Diff.html:27
			qw422016.N().S(`</em></code></td><td style="width: 10%;">→</td><td style="width: 30%;"><code class="success">`)
//line views/components/view/Diff.html:29
			qw422016.E().S(d.New)
//line views/components/view/Diff.html:29
			qw422016.N().S(`</code></td></tr>`)
//line views/components/view/Diff.html:31
		}
//line views/components/view/Diff.html:31
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Diff.html:35
	}
//line views/components/view/Diff.html:36
}

//line views/components/view/Diff.html:36
func WriteDiffs(qq422016 qtio422016.Writer, value util.Diffs) {
//line views/components/view/Diff.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Diff.html:36
	StreamDiffs(qw422016, value)
//line views/components/view/Diff.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Diff.html:36
}

//line views/components/view/Diff.html:36
func Diffs(value util.Diffs) string {
//line views/components/view/Diff.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Diff.html:36
	WriteDiffs(qb422016, value)
//line views/components/view/Diff.html:36
	qs422016 := string(qb422016.B)
//line views/components/view/Diff.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Diff.html:36
	return qs422016
//line views/components/view/Diff.html:36
}

//line views/components/view/Diff.html:38
func StreamDiffsSet(qw422016 *qt422016.Writer, key string, value util.DiffsSet, ps *cutil.PageState) {
//line views/components/view/Diff.html:39
	if len(value) == 0 {
//line views/components/view/Diff.html:39
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/view/Diff.html:41
	} else {
//line views/components/view/Diff.html:41
		qw422016.N().S(`<ul class="accordion">`)
//line views/components/view/Diff.html:43
		for idx, k := range util.ArraySorted[string](lo.Keys(value)) {
//line views/components/view/Diff.html:44
			dk, u := util.StringSplitLast(k, '^', true)

//line views/components/view/Diff.html:45
			v := value[k]

//line views/components/view/Diff.html:46
			if idx < 100 {
//line views/components/view/Diff.html:46
				qw422016.N().S(`<li><input id="accordion-`)
//line views/components/view/Diff.html:48
				qw422016.E().S(k)
//line views/components/view/Diff.html:48
				qw422016.N().S(`-`)
//line views/components/view/Diff.html:48
				qw422016.N().D(idx)
//line views/components/view/Diff.html:48
				qw422016.N().S(`" type="checkbox" hidden="hidden" /><label for="accordion-`)
//line views/components/view/Diff.html:49
				qw422016.E().S(k)
//line views/components/view/Diff.html:49
				qw422016.N().S(`-`)
//line views/components/view/Diff.html:49
				qw422016.N().D(idx)
//line views/components/view/Diff.html:49
				qw422016.N().S(`"><div class="right">`)
//line views/components/view/Diff.html:51
				if len(v) == 1 {
//line views/components/view/Diff.html:51
					qw422016.N().S(`<em>(`)
//line views/components/view/Diff.html:52
					qw422016.E().S(v[0].String())
//line views/components/view/Diff.html:52
					qw422016.N().S(`)</em>`)
//line views/components/view/Diff.html:52
					qw422016.N().S(` `)
//line views/components/view/Diff.html:53
				}
//line views/components/view/Diff.html:54
				qw422016.E().S(util.StringPlural(len(v), "diff"))
//line views/components/view/Diff.html:54
				qw422016.N().S(`</div>`)
//line views/components/view/Diff.html:56
				components.StreamExpandCollapse(qw422016, 3, ps)
//line views/components/view/Diff.html:57
				if u != "" {
//line views/components/view/Diff.html:57
					qw422016.N().S(`<a href="`)
//line views/components/view/Diff.html:57
					qw422016.E().S(u)
//line views/components/view/Diff.html:57
					qw422016.N().S(`">`)
//line views/components/view/Diff.html:57
					qw422016.E().S(dk)
//line views/components/view/Diff.html:57
					qw422016.N().S(`</a>`)
//line views/components/view/Diff.html:57
				} else {
//line views/components/view/Diff.html:57
					qw422016.E().S(dk)
//line views/components/view/Diff.html:57
				}
//line views/components/view/Diff.html:57
				qw422016.N().S(`</label><div class="bd"><div><div>`)
//line views/components/view/Diff.html:60
				StreamDiffs(qw422016, v)
//line views/components/view/Diff.html:60
				qw422016.N().S(`</div></div></div></li>`)
//line views/components/view/Diff.html:63
			}
//line views/components/view/Diff.html:64
			if idx == 100 {
//line views/components/view/Diff.html:64
				qw422016.N().S(`<li><input id="accordion-`)
//line views/components/view/Diff.html:66
				qw422016.E().S(k)
//line views/components/view/Diff.html:66
				qw422016.N().S(`-extras" type="checkbox" hidden="hidden" /><label for="accordion-`)
//line views/components/view/Diff.html:67
				qw422016.E().S(k)
//line views/components/view/Diff.html:67
				qw422016.N().S(`-extras">...and`)
//line views/components/view/Diff.html:67
				qw422016.N().S(` `)
//line views/components/view/Diff.html:67
				qw422016.N().D(len(value) - 100)
//line views/components/view/Diff.html:67
				qw422016.N().S(` `)
//line views/components/view/Diff.html:67
				qw422016.N().S(`extra</label></li>`)
//line views/components/view/Diff.html:69
			}
//line views/components/view/Diff.html:70
		}
//line views/components/view/Diff.html:70
		qw422016.N().S(`</ul>`)
//line views/components/view/Diff.html:72
	}
//line views/components/view/Diff.html:73
}

//line views/components/view/Diff.html:73
func WriteDiffsSet(qq422016 qtio422016.Writer, key string, value util.DiffsSet, ps *cutil.PageState) {
//line views/components/view/Diff.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Diff.html:73
	StreamDiffsSet(qw422016, key, value, ps)
//line views/components/view/Diff.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Diff.html:73
}

//line views/components/view/Diff.html:73
func DiffsSet(key string, value util.DiffsSet, ps *cutil.PageState) string {
//line views/components/view/Diff.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Diff.html:73
	WriteDiffsSet(qb422016, key, value, ps)
//line views/components/view/Diff.html:73
	qs422016 := string(qb422016.B)
//line views/components/view/Diff.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Diff.html:73
	return qs422016
//line views/components/view/Diff.html:73
}
