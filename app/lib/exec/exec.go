// Content managed by Project Forge, see [projectforge.md] for details.
package exec

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"github.com/kyleu/solitaire/app/util"
)

type writer struct {
	Key string
	fn  func(key string, b []byte) error
}

func (w *writer) Write(p []byte) (int, error) {
	err := w.fn(w.Key, p)
	return len(p), err
}

type Exec struct {
	Key       string        `json:"key"`
	Idx       int           `json:"idx"`
	Cmd       string        `json:"cmd"`
	Env       []string      `json:"env,omitempty"`
	Path      string        `json:"path"`
	Started   *time.Time    `json:"started"`
	PID       int           `json:"pid"`
	Completed *time.Time    `json:"completed"`
	ExitCode  int           `json:"exitCode"`
	Link      string        `json:"link,omitempty"`
	Buffer    *bytes.Buffer `json:"-"`
	execCmd   *exec.Cmd
}

func NewExec(key string, idx int, cmd string, path string, envvars ...string) *Exec {
	return &Exec{Key: key, Idx: idx, Cmd: cmd, Env: envvars, Path: path, Buffer: &bytes.Buffer{}}
}

func (e *Exec) WebPath() string {
	return fmt.Sprintf("/admin/exec/%s/%d", e.Key, e.Idx)
}

func (e *Exec) Start(fns ...func(key string, b []byte) error) error {
	if e.Started != nil {
		return errors.New("process already started")
	}
	var w io.Writer = e.Buffer
	lo.ForEach(fns, func(fn func(key string, b []byte) error, _ int) {
		w = io.MultiWriter(w, &writer{Key: e.String(), fn: fn})
	})
	e.Started = util.TimeCurrentP()
	cmd, err := util.StartProcess(e.Cmd, e.Path, nil, w, w, e.Env...)
	if err != nil {
		return err
	}
	e.execCmd = cmd
	e.PID = cmd.Process.Pid
	defer func() {
		go func() {
			_ = e.Wait()
		}()
	}()
	return nil
}

func (e *Exec) Kill() error {
	if e.execCmd == nil {
		return errors.New("not started")
	}
	return e.execCmd.Process.Kill()
}

func (e *Exec) Wait() error {
	if e.execCmd == nil {
		return errors.New("not started")
	}
	exit, err := e.execCmd.Process.Wait()
	if err != nil {
		_ = e.execCmd.Wait()
		for (e.execCmd.ProcessState == nil || !e.execCmd.ProcessState.Exited()) && time.Since(util.TimeCurrent()) < (4*time.Second) {
			time.Sleep(500 * time.Millisecond)
		}
		exit = e.execCmd.ProcessState
	}

	e.Completed = util.TimeCurrentP()
	if exit != nil {
		e.ExitCode = exit.ExitCode()
	}
	return nil
}

func (e *Exec) String() string {
	return fmt.Sprintf("%s:%d", e.Key, e.Idx)
}

type Execs []*Exec

func (m Execs) Get(key string, idx int) *Exec {
	return lo.FindOrElse(m, nil, func(x *Exec) bool {
		return x.Key == key && x.Idx == idx
	})
}

func (m Execs) GetByKey(key string) Execs {
	return lo.Filter(m, func(x *Exec, _ int) bool {
		return x.Key == key
	})
}

func (m Execs) Running() int {
	return lo.CountBy(m, func(x *Exec) bool {
		return x.Completed == nil
	})
}

func (m Execs) Sort() {
	slices.SortFunc(m, func(l *Exec, r *Exec) bool {
		lk, rk := strings.ToLower(l.Key), strings.ToLower(r.Key)
		if lk != rk {
			return lk < rk
		}
		return l.Idx < r.Idx
	})
}
