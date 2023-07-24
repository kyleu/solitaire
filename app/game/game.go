package game

import (
	"github.com/google/uuid"

	"github.com/kyleu/solitaire/app/game/deck"
	"github.com/kyleu/solitaire/app/game/pile"
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/util"
)

type Game struct {
	ID    uuid.UUID    `json:"id"`
	Rules *rules.Rules `json:"rules,omitempty"`
	Seed  int          `json:"seed"`

	Deck     *deck.Deck `json:"deck,omitempty"`
	PileSets pile.Sets  `json:"pileSets"`

	Players Players `json:"players,omitempty"`

	StockCounter int `json:"stockCounter,omitempty"`
}

func New(rules *rules.Rules) *Game {
	d := deck.FreshDeck(1, nil, nil, false, 0)
	return &Game{ID: util.UUID(), Rules: rules, Deck: d, PileSets: nil}
}
