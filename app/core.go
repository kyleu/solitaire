// Package app - Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"

	"github.com/kyleu/solitaire/app/lib/websocket"
	"github.com/kyleu/solitaire/app/util"
)

type CoreServices struct {
	Socket *websocket.Service
}

//nolint:revive
func initCoreServices(ctx context.Context, st *State, logger util.Logger) CoreServices {
	return CoreServices{
		Socket: websocket.NewService(nil, nil),
	}
}
