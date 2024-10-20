//go:build js
// +build js

package wasm

import (
	"time"

	"github.com/kyleu/solitaire/app/lib/log"
	"github.com/kyleu/solitaire/app/util"
)

type WASM struct {
	logger  util.Logger
	CloseCh chan struct{}
}

func NewWASM() *WASM {
	ret := &WASM{CloseCh: make(chan struct{})}

	logFn := func(level string, occurred time.Time, loggerName string, message string, caller util.ValueMap, stack string, fields util.ValueMap) {
		ret.Log(level, occurred, loggerName, message, caller, stack, fields)
	}
	l, err := log.InitLogging(true, logFn)
	if err != nil {
		println(err)
	}
	ret.logger = l

	t := util.TimerStart()
	ret.wireFunctions()
	l.Infof("[%s] started in [%s]", util.AppName, t.EndString())

	return ret
}
