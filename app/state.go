// Package app - Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/app/lib/graphql"
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

func NewState(debug bool, bi *BuildInfo, f filesystem.FileLoader, enableTelemetry bool, _ uint16, logger util.Logger) (*State, error) {
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

	_ = telemetry.InitializeIfNeeded(enableTelemetry, bi.Version, logger)

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
