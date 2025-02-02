package rules

import (
	"github.com/kyleu/solitaire/app/game/deck"
	"github.com/kyleu/solitaire/app/game/rules/cell"
	"github.com/kyleu/solitaire/app/game/rules/foundation"
	"github.com/kyleu/solitaire/app/game/rules/pyramid"
	"github.com/kyleu/solitaire/app/game/rules/reserve"
	"github.com/kyleu/solitaire/app/game/rules/stock"
	"github.com/kyleu/solitaire/app/game/rules/tableau"
	"github.com/kyleu/solitaire/app/game/rules/waste"
)

type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Links []*Link

type Rules struct {
	Key       string `json:"key"`
	Completed bool   `json:"completed,omitempty"`
	Name      string `json:"name"`
	Layout    string `json:"layout"`

	AKA     map[string]string `json:"aka,omitempty"`
	Like    string            `json:"like,omitempty"`
	Related []string          `json:"related,omitempty"`
	Links   Links             `json:"links,omitempty"`

	MaxPlayers        int           `json:"maxPlayers,omitempty"`
	VictoryCondition  string        `json:"victoryCondition,omitempty"`
	CardRemovalMethod string        `json:"cardRemovalMethod,omitempty"`
	DeckOptions       *deck.Options `json:"deckOptions,omitempty"`

	Stock       *stock.Stock             `json:"stock,omitempty"`
	Waste       *waste.Waste             `json:"waste,omitempty"`
	Reserves    *reserve.Reserve         `json:"reserves,omitempty"`
	Cells       *cell.Cell               `json:"cells,omitempty"`
	Foundations []*foundation.Foundation `json:"foundations,omitempty"`
	Tableaus    []*tableau.Tableau       `json:"tableaus,omitempty"`
	Pyramids    []*pyramid.Pyramid       `json:"pyramids,omitempty"`
	Special     *Special                 `json:"special,omitempty"`

	Error string `json:"error,omitempty"`
}
