// Code generated by qtc from "Game.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vgame/Game.html:1
package vgame

//line views/vgame/Game.html:1
import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/views/layout"
)

//line views/vgame/Game.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vgame/Game.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vgame/Game.html:7
type Game struct {
	layout.Basic
}

//line views/vgame/Game.html:11
func (p *Game) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgame/Game.html:11
	qw422016.N().S(`
  <div class="card">
    <h3>Game Test</h3>
    <div class="mt">Good luck!</div>
  </div>
  <div id="playfield" class="card">
    ...
  </div>
`)
//line views/vgame/Game.html:19
}

//line views/vgame/Game.html:19
func (p *Game) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vgame/Game.html:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vgame/Game.html:19
	p.StreamBody(qw422016, as, ps)
//line views/vgame/Game.html:19
	qt422016.ReleaseWriter(qw422016)
//line views/vgame/Game.html:19
}

//line views/vgame/Game.html:19
func (p *Game) Body(as *app.State, ps *cutil.PageState) string {
//line views/vgame/Game.html:19
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vgame/Game.html:19
	p.WriteBody(qb422016, as, ps)
//line views/vgame/Game.html:19
	qs422016 := string(qb422016.B)
//line views/vgame/Game.html:19
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vgame/Game.html:19
	return qs422016
//line views/vgame/Game.html:19
}
