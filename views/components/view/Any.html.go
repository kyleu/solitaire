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

	"github.com/kyleu/solitaire/app/util"
)

//line views/components/view/Any.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Any.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Any.html:11
func StreamAny(qw422016 *qt422016.Writer, x any) {
//line views/components/view/Any.html:12
	if x == nil {
//line views/components/view/Any.html:12
		qw422016.N().S(`<em>nil</em>`)
//line views/components/view/Any.html:14
	} else {
//line views/components/view/Any.html:15
		switch t := x.(type) {
//line views/components/view/Any.html:16
		case bool:
//line views/components/view/Any.html:17
			StreamBool(qw422016, t)
//line views/components/view/Any.html:18
		case util.Diffs:
//line views/components/view/Any.html:19
			StreamDiffs(qw422016, t)
//line views/components/view/Any.html:20
		case float32:
//line views/components/view/Any.html:21
			StreamFloat(qw422016, t)
//line views/components/view/Any.html:22
		case float64:
//line views/components/view/Any.html:23
			StreamFloat(qw422016, t)
//line views/components/view/Any.html:24
		case int:
//line views/components/view/Any.html:25
			StreamInt(qw422016, t)
//line views/components/view/Any.html:26
		case int32:
//line views/components/view/Any.html:27
			StreamInt(qw422016, t)
//line views/components/view/Any.html:28
		case int64:
//line views/components/view/Any.html:29
			StreamInt(qw422016, t)
//line views/components/view/Any.html:30
		case util.ValueMap:
//line views/components/view/Any.html:31
			StreamMap(qw422016, false, t)
//line views/components/view/Any.html:32
		case []util.ValueMap:
//line views/components/view/Any.html:33
			StreamMapArray(qw422016, false, t...)
//line views/components/view/Any.html:34
		case util.Pkg:
//line views/components/view/Any.html:35
			StreamPackage(qw422016, t)
//line views/components/view/Any.html:36
		case string:
//line views/components/view/Any.html:37
			StreamString(qw422016, t)
//line views/components/view/Any.html:38
		case time.Time:
//line views/components/view/Any.html:39
			StreamTimestamp(qw422016, &t)
//line views/components/view/Any.html:40
		case *time.Time:
//line views/components/view/Any.html:41
			StreamTimestamp(qw422016, t)
//line views/components/view/Any.html:42
		case url.URL:
//line views/components/view/Any.html:43
			StreamURL(qw422016, t)
//line views/components/view/Any.html:44
		case *url.URL:
//line views/components/view/Any.html:45
			StreamURL(qw422016, t)
//line views/components/view/Any.html:46
		case uuid.UUID:
//line views/components/view/Any.html:47
			StreamUUID(qw422016, &t)
//line views/components/view/Any.html:48
		case *uuid.UUID:
//line views/components/view/Any.html:49
			StreamUUID(qw422016, t)
//line views/components/view/Any.html:50
		default:
//line views/components/view/Any.html:50
			qw422016.N().S(`unhandled type [`)
//line views/components/view/Any.html:51
			qw422016.E().S(fmt.Sprintf("%T", x))
//line views/components/view/Any.html:51
			qw422016.N().S(`]`)
//line views/components/view/Any.html:52
		}
//line views/components/view/Any.html:53
	}
//line views/components/view/Any.html:54
}

//line views/components/view/Any.html:54
func WriteAny(qq422016 qtio422016.Writer, x any) {
//line views/components/view/Any.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Any.html:54
	StreamAny(qw422016, x)
//line views/components/view/Any.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Any.html:54
}

//line views/components/view/Any.html:54
func Any(x any) string {
//line views/components/view/Any.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Any.html:54
	WriteAny(qb422016, x)
//line views/components/view/Any.html:54
	qs422016 := string(qb422016.B)
//line views/components/view/Any.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Any.html:54
	return qs422016
//line views/components/view/Any.html:54
}
