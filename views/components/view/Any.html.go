// Code generated by qtc from "Any.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/Any.html:1
package view

//line views/components/view/Any.html:1
import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
)

//line views/components/view/Any.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Any.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Any.html:12
func StreamAny(qw422016 *qt422016.Writer, x any, ps *cutil.PageState) {
//line views/components/view/Any.html:13
	if x == nil {
//line views/components/view/Any.html:13
		qw422016.N().S(`<em>nil</em>`)
//line views/components/view/Any.html:15
	} else {
//line views/components/view/Any.html:16
		switch t := x.(type) {
//line views/components/view/Any.html:17
		case bool:
//line views/components/view/Any.html:18
			StreamBool(qw422016, t)
//line views/components/view/Any.html:19
		case util.Diffs:
//line views/components/view/Any.html:20
			StreamDiffs(qw422016, t)
//line views/components/view/Any.html:21
		case float32:
//line views/components/view/Any.html:22
			StreamFloat(qw422016, t)
//line views/components/view/Any.html:23
		case float64:
//line views/components/view/Any.html:24
			StreamFloat(qw422016, t)
//line views/components/view/Any.html:25
		case int:
//line views/components/view/Any.html:26
			StreamInt(qw422016, t)
//line views/components/view/Any.html:27
		case int32:
//line views/components/view/Any.html:28
			StreamInt(qw422016, t)
//line views/components/view/Any.html:29
		case int64:
//line views/components/view/Any.html:30
			StreamInt(qw422016, t)
//line views/components/view/Any.html:31
		case util.ToMap:
//line views/components/view/Any.html:32
			StreamMap(qw422016, false, t.ToMap(), ps)
//line views/components/view/Any.html:33
		case util.ValueMap:
//line views/components/view/Any.html:34
			StreamMap(qw422016, false, t, ps)
//line views/components/view/Any.html:35
		case map[string]any:
//line views/components/view/Any.html:36
			StreamMap(qw422016, false, t, ps)
//line views/components/view/Any.html:37
		case util.ToMaps:
//line views/components/view/Any.html:38
			StreamMapArray(qw422016, false, ps, t.ToMaps()...)
//line views/components/view/Any.html:39
		case []util.ValueMap:
//line views/components/view/Any.html:40
			StreamMapArray(qw422016, false, ps, t...)
//line views/components/view/Any.html:41
		case util.Pkg:
//line views/components/view/Any.html:42
			StreamPackage(qw422016, t)
//line views/components/view/Any.html:43
		case string:
//line views/components/view/Any.html:44
			StreamString(qw422016, t)
//line views/components/view/Any.html:45
		case []string:
//line views/components/view/Any.html:46
			StreamStringArray(qw422016, t)
//line views/components/view/Any.html:47
		case time.Time:
//line views/components/view/Any.html:48
			StreamTimestamp(qw422016, &t)
//line views/components/view/Any.html:49
		case *time.Time:
//line views/components/view/Any.html:50
			StreamTimestamp(qw422016, t)
//line views/components/view/Any.html:51
		case url.URL:
//line views/components/view/Any.html:52
			StreamURL(qw422016, t, "", true, ps)
//line views/components/view/Any.html:53
		case *url.URL:
//line views/components/view/Any.html:54
			StreamURL(qw422016, t, "", true, ps)
//line views/components/view/Any.html:55
		case uuid.UUID:
//line views/components/view/Any.html:56
			StreamUUID(qw422016, &t)
//line views/components/view/Any.html:57
		case *uuid.UUID:
//line views/components/view/Any.html:58
			StreamUUID(qw422016, t)
//line views/components/view/Any.html:59
		case []any:
//line views/components/view/Any.html:60
			if len(t) == 0 {
//line views/components/view/Any.html:60
				qw422016.N().S(`<em>empty array</em>`)
//line views/components/view/Any.html:62
			} else {
//line views/components/view/Any.html:63
				arr, extra := util.ArrayLimit(t, 8)

//line views/components/view/Any.html:64
				for idx, e := range arr {
//line views/components/view/Any.html:64
					qw422016.N().S(`<div class="flex bt"><div class="mrs mts"><em>`)
//line views/components/view/Any.html:66
					qw422016.N().D(idx + 1)
//line views/components/view/Any.html:66
					qw422016.N().S(`</em></div><div class="">`)
//line views/components/view/Any.html:67
					StreamAny(qw422016, e, ps)
//line views/components/view/Any.html:67
					qw422016.N().S(`</div></div>`)
//line views/components/view/Any.html:69
				}
//line views/components/view/Any.html:70
				if extra > 0 {
//line views/components/view/Any.html:70
					qw422016.N().S(`<div class="bb"><em>...and`)
//line views/components/view/Any.html:71
					qw422016.N().S(` `)
//line views/components/view/Any.html:71
					qw422016.N().D(extra)
//line views/components/view/Any.html:71
					qw422016.N().S(` `)
//line views/components/view/Any.html:71
					qw422016.N().S(`more</em></div>`)
//line views/components/view/Any.html:72
				}
//line views/components/view/Any.html:73
			}
//line views/components/view/Any.html:74
		default:
//line views/components/view/Any.html:74
			qw422016.N().S(`unhandled type [`)
//line views/components/view/Any.html:75
			qw422016.E().S(fmt.Sprintf("%T", x))
//line views/components/view/Any.html:75
			qw422016.N().S(`]`)
//line views/components/view/Any.html:76
		}
//line views/components/view/Any.html:77
	}
//line views/components/view/Any.html:78
}

//line views/components/view/Any.html:78
func WriteAny(qq422016 qtio422016.Writer, x any, ps *cutil.PageState) {
//line views/components/view/Any.html:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Any.html:78
	StreamAny(qw422016, x, ps)
//line views/components/view/Any.html:78
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Any.html:78
}

//line views/components/view/Any.html:78
func Any(x any, ps *cutil.PageState) string {
//line views/components/view/Any.html:78
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Any.html:78
	WriteAny(qb422016, x, ps)
//line views/components/view/Any.html:78
	qs422016 := string(qb422016.B)
//line views/components/view/Any.html:78
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Any.html:78
	return qs422016
//line views/components/view/Any.html:78
}
