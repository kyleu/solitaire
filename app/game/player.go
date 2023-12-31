package game

import "github.com/google/uuid"

type Player struct {
	UserID         uuid.UUID `json:"id"`
	Name           string    `json:"n,omitempty"`
	AutoFlipOption bool      `json:"af,omitempty"`
}

type Players []*Player
