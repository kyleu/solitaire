package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/app/lib/graphql"
	"github.com/kyleu/solitaire/app/lib/log"
	"github.com/kyleu/solitaire/app/lib/telemetry"
	"github.com/kyleu/solitaire/app/lib/theme"
	"github.com/kyleu/solitaire/app/util"
)

var once sync.Once

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	if b.Date == util.KeyUnknown {
		return b.Version
	}
	d, _ := util.TimeFromJS(b.Date)
	return fmt.Sprintf("%s (%s)", b.Version, util.TimeToYMD(d))
}

type State struct {
	Debug     bool
	BuildInfo *BuildInfo
	Files     filesystem.FileLoader
	GraphQL   *graphql.Service
	Themes    *theme.Service
	Services  *Services
	Started   time.Time
}

func NewState(ctx context.Context, debug bool, bi *BuildInfo, f filesystem.FileLoader, enableTelemetry bool, _ uint16, logger util.Logger) (*State, error) {
	var loadLocationError error
	once.Do(func() {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			loadLocationError = err
			return
		}
		time.Local = loc
	})
	if loadLocationError != nil {
		return nil, loadLocationError
	}

	_ = telemetry.InitializeIfNeeded(ctx, enableTelemetry, bi.Version, logger)

	return &State{
		Debug:     debug,
		BuildInfo: bi,
		Files:     f,
		GraphQL:   graphql.NewService(),
		Themes:    theme.NewService(f),
		Started:   util.TimeCurrent(),
	}, nil
}

func (s State) Close(ctx context.Context, logger util.Logger) error {
	defer func() { _ = telemetry.Close(ctx) }()
	if err := s.GraphQL.Close(); err != nil {
		logger.Errorf("error closing GraphQL service: %+v", err)
	}
	return s.Services.Close(ctx, logger)
}

func Bootstrap(ctx context.Context, bi *BuildInfo, cfgDir string, port uint16, debug bool, logger util.Logger) (*State, error) {
	fs, err := filesystem.NewFileSystem(cfgDir, false, "")
	if err != nil {
		return nil, err
	}

	telemetryDisabled := util.GetEnvBool("disable_telemetry", false)
	st, err := NewState(ctx, debug, bi, fs, !telemetryDisabled, port, logger)
	if err != nil {
		return nil, err
	}

	ctx, span, logger := telemetry.StartSpan(ctx, "app:init", logger)
	defer span.Complete()
	t := util.TimerStart()
	svcs, err := NewServices(ctx, st, logger)
	if err != nil {
		return nil, err
	}
	logger.Debugf("created app state in [%s]", util.MicrosToMillis(t.End()))
	st.Services = svcs

	return st, nil
}

func BootstrapRunDefault[T any](ctx context.Context, bi *BuildInfo, fn func(as *State, logger util.Logger) (T, error)) (T, error) {
	logger, _ := log.InitLogging(false)
	as, err := Bootstrap(ctx, bi, util.ConfigDir, 0, false, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	ret, err := fn(as, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	err = as.Close(ctx, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	return ret, nil
}
