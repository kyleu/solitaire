// Code generated by qtc from "Map.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Map.html:2
package view

//line views/components/view/Map.html:2
import "github.com/kyleu/solitaire/app/util"

//line views/components/view/Map.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Map.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Map.html:4
func StreamMap(qw422016 *qt422016.Writer, preserveWhitespace bool, m util.ValueMap) {
//line views/components/view/Map.html:5
	if m == nil {
//line views/components/view/Map.html:5
		qw422016.N().S(`<em>no result</em>`)
//line views/components/view/Map.html:7
	} else if len(m) == 0 {
//line views/components/view/Map.html:7
		qw422016.N().S(`<em>empty result</em>`)
//line views/components/view/Map.html:9
	} else {
//line views/components/view/Map.html:9
		qw422016.N().S(`<div class="overflow full-width"><table><tbody>`)
//line views/components/view/Map.html:13
		for _, k := range m.Keys() {
//line views/components/view/Map.html:13
			qw422016.N().S(`<tr><th>`)
//line views/components/view/Map.html:15
			qw422016.E().S(k)
//line views/components/view/Map.html:15
			qw422016.N().S(`</th>`)
//line views/components/view/Map.html:16
			if preserveWhitespace {
//line views/components/view/Map.html:16
				qw422016.N().S(`<td class="prews">`)
//line views/components/view/Map.html:17
				StreamAny(qw422016, m[k])
//line views/components/view/Map.html:17
				qw422016.N().S(`</td>`)
//line views/components/view/Map.html:18
			} else {
//line views/components/view/Map.html:18
				qw422016.N().S(`<td>`)
//line views/components/view/Map.html:19
				StreamAny(qw422016, m[k])
//line views/components/view/Map.html:19
				qw422016.N().S(`</td>`)
//line views/components/view/Map.html:20
			}
//line views/components/view/Map.html:20
			qw422016.N().S(`</tr>`)
//line views/components/view/Map.html:22
		}
//line views/components/view/Map.html:22
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Map.html:26
	}
//line views/components/view/Map.html:27
}

//line views/components/view/Map.html:27
func WriteMap(qq422016 qtio422016.Writer, preserveWhitespace bool, m util.ValueMap) {
//line views/components/view/Map.html:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Map.html:27
	StreamMap(qw422016, preserveWhitespace, m)
//line views/components/view/Map.html:27
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Map.html:27
}

//line views/components/view/Map.html:27
func Map(preserveWhitespace bool, m util.ValueMap) string {
//line views/components/view/Map.html:27
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Map.html:27
	WriteMap(qb422016, preserveWhitespace, m)
//line views/components/view/Map.html:27
	qs422016 := string(qb422016.B)
//line views/components/view/Map.html:27
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Map.html:27
	return qs422016
//line views/components/view/Map.html:27
}

//line views/components/view/Map.html:29
func StreamMapArray(qw422016 *qt422016.Writer, preserveWhitespace bool, maps ...util.ValueMap) {
//line views/components/view/Map.html:30
	if len(maps) == 0 {
//line views/components/view/Map.html:30
		qw422016.N().S(`<em>no results</em>`)
//line views/components/view/Map.html:32
	} else {
//line views/components/view/Map.html:32
		qw422016.N().S(`<div class="overflow full-width"><table><thead><tr>`)
//line views/components/view/Map.html:37
		for _, k := range maps[0].Keys() {
//line views/components/view/Map.html:37
			qw422016.N().S(`<th>`)
//line views/components/view/Map.html:38
			qw422016.E().S(k)
//line views/components/view/Map.html:38
			qw422016.N().S(`</th>`)
//line views/components/view/Map.html:39
		}
//line views/components/view/Map.html:39
		qw422016.N().S(`</tr></thead><tbody>`)
//line views/components/view/Map.html:43
		for _, m := range maps {
//line views/components/view/Map.html:43
			qw422016.N().S(`<tr>`)
//line views/components/view/Map.html:45
			for _, k := range m.Keys() {
//line views/components/view/Map.html:46
				if preserveWhitespace {
//line views/components/view/Map.html:46
					qw422016.N().S(`<td class="prews">`)
//line views/components/view/Map.html:47
					StreamAny(qw422016, m[k])
//line views/components/view/Map.html:47
					qw422016.N().S(`</td>`)
//line views/components/view/Map.html:48
				} else {
//line views/components/view/Map.html:48
					qw422016.N().S(`<td>`)
//line views/components/view/Map.html:49
					StreamAny(qw422016, m[k])
//line views/components/view/Map.html:49
					qw422016.N().S(`</td>`)
//line views/components/view/Map.html:50
				}
//line views/components/view/Map.html:51
			}
//line views/components/view/Map.html:51
			qw422016.N().S(`</tr>`)
//line views/components/view/Map.html:53
		}
//line views/components/view/Map.html:53
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/view/Map.html:57
	}
//line views/components/view/Map.html:58
}

//line views/components/view/Map.html:58
func WriteMapArray(qq422016 qtio422016.Writer, preserveWhitespace bool, maps ...util.ValueMap) {
//line views/components/view/Map.html:58
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Map.html:58
	StreamMapArray(qw422016, preserveWhitespace, maps...)
//line views/components/view/Map.html:58
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Map.html:58
}

//line views/components/view/Map.html:58
func MapArray(preserveWhitespace bool, maps ...util.ValueMap) string {
//line views/components/view/Map.html:58
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Map.html:58
	WriteMapArray(qb422016, preserveWhitespace, maps...)
//line views/components/view/Map.html:58
	qs422016 := string(qb422016.B)
//line views/components/view/Map.html:58
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Map.html:58
	return qs422016
//line views/components/view/Map.html:58
}
