// Package cmd - Content managed by Project Forge, see [projectforge.md] for details.
package cmd

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/app/lib/telemetry"
	"github.com/kyleu/solitaire/app/util"
)

func buildDefaultAppState(flags *Flags, logger util.Logger) (*app.State, error) {
	fs, err := filesystem.NewFileSystem(flags.ConfigDir, false, "")
	if err != nil {
		return nil, err
	}

	telemetryDisabled := util.GetEnvBool("disable_telemetry", false)
	st, err := app.NewState(flags.Debug, _buildInfo, fs, !telemetryDisabled, flags.Port, logger)
	if err != nil {
		return nil, err
	}

	ctx, span, logger := telemetry.StartSpan(context.Background(), "app:init", logger)
	defer span.Complete()
	t := util.TimerStart()
	svcs, err := app.NewServices(ctx, st, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error creating services")
	}
	logger.Debugf("created app state in [%s]", util.MicrosToMillis(t.End()))
	st.Services = svcs

	return st, nil
}
