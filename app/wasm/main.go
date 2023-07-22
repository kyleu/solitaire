// Content managed by Project Forge, see [projectforge.md] for details.
//go:build js

package main

import (
	lg "github.com/kyleu/solitaire/app/lib/log"
	"github.com/kyleu/solitaire/app/util"
)

var _rootLogger util.Logger

func main() {
	l, err := lg.InitLogging(true)
	if err != nil {
		println(err)
	}
	_rootLogger = l

	t := util.TimerStart()
	wireFunctions()

	initWASM(l)

	l.Infof("[%s] started in [%s]", util.AppName, t.EndString())
	<-make(chan struct{})
}
