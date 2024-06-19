package app

import (
	"context"

	"github.com/kyleu/solitaire/app/lib/websocket"
	"github.com/kyleu/solitaire/app/util"
)

type CoreServices struct {
	Socket *websocket.Service
}

func initCoreServices(ctx context.Context, st *State, logger util.Logger) CoreServices {
	return CoreServices{
		Socket: websocket.NewService(nil, nil),
	}
}
