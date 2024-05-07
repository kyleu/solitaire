package app

import (
	"context"

	"github.com/kyleu/solitaire/app/lib/websocket"
	"github.com/kyleu/solitaire/app/util"
)

type Services struct {
	CoreServices
	Socket *websocket.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	return &Services{CoreServices: initCoreServices(ctx, st, logger)}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
