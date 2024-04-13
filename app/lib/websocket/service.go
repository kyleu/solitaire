// Package websocket - Content managed by Project Forge, see [projectforge.md] for details.
package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/lib/user"
	"github.com/kyleu/solitaire/app/util"
)

type Handler func(ctx context.Context, s *Service, conn *Connection, svc string, cmd string, param json.RawMessage, logger util.Logger) error

type ConnectEvent func(s *Service, conn *Connection, logger util.Logger) error

type Service struct {
	connections   map[uuid.UUID]*Connection
	connectionsMu sync.Mutex
	channels      map[string]*Channel
	channelsMu    sync.Mutex
	taps          map[uuid.UUID]*websocket.Conn
	tapsMu        sync.Mutex
	onOpen        ConnectEvent
	handler       Handler
	onClose       ConnectEvent
}

func NewService(onOpen ConnectEvent, handler Handler, onClose ConnectEvent) *Service {
	return &Service{
		connections: make(map[uuid.UUID]*Connection),
		channels:    make(map[string]*Channel),
		taps:        make(map[uuid.UUID]*websocket.Conn),
		onOpen:      onOpen,
		handler:     handler,
		onClose:     onClose,
	}
}

var (
	systemID     = *util.UUIDFromString("FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF")
	systemStatus = &Status{
		ID: systemID, Username: util.GetEnv("system_username", "System Broadcast"), Channels: []string{systemID.String()},
	}
)

func (s *Service) ReplaceHandlers(onOpen ConnectEvent, handler Handler, onClose ConnectEvent) {
	s.onOpen = onOpen
	s.handler = handler
	s.onClose = onClose
}

func (s *Service) Close() {
	s.connectionsMu.Lock()
	defer s.connectionsMu.Unlock()
	lo.ForEach(lo.Values(s.connections), func(v *Connection, _ int) {
		_ = v.Close()
	})
}

var upgrader = websocket.Upgrader{EnableCompression: true}

func (s *Service) Upgrade(
	ctx context.Context, w http.ResponseWriter, r *http.Request, channel string, profile *user.Profile, logger util.Logger,
) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	cx, err := s.Register(profile, conn, logger)
	if err != nil {
		logger.Warnf("unable to register websocket connection: %+v", err)
		return nil
	}
	joined, err := s.Join(cx.ID, channel, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("error processing socket join (%v): %+v", joined, err))
		return nil
	}
	err = s.ReadLoop(ctx, cx.ID, logger)
	if err != nil {
		if !strings.Contains(err.Error(), "1001") {
			logger.Error(fmt.Sprintf("error processing socket read loop for connection [%s]: %+v", cx.ID.String(), err))
		}
		return nil
	}
	return nil
}
