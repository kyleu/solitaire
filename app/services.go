package app

import (
	"context"
	"encoding/json"

	"github.com/kyleu/solitaire/app/lib/exec"
	"github.com/kyleu/solitaire/app/lib/websocket"
	"github.com/kyleu/solitaire/app/util"
)

type Services struct {
	Exec   *exec.Service
	Socket *websocket.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	return &Services{
		Exec:   exec.NewService(),
		Socket: websocket.NewService(nil, socketHandler, nil),
	}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}

func socketHandler(_ context.Context, s *websocket.Service, c *websocket.Connection, _ string, cmd string, _ json.RawMessage, logger util.Logger) error {
	switch cmd {
	case "connect":
		_, err := s.Join(c.ID, "tap", logger)
		if err != nil {
			return err
		}
	default:
		logger.Error("unhandled command [" + cmd + "]")
	}
	return nil
}
