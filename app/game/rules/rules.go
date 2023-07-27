package rules

import (
	"github.com/kyleu/solitaire/app/game/deck"
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

	Stock       *Stock        `json:"stock,omitempty"`
	Waste       *Waste        `json:"waste,omitempty"`
	Reserves    *Reserve      `json:"reserves,omitempty"`
	Cells       *Cell         `json:"cells,omitempty"`
	Foundations []*Foundation `json:"foundations,omitempty"`
	Tableaus    []*Tableau    `json:"tableaus,omitempty"`
	Pyramids    []*Pyramid    `json:"pyramids,omitempty"`
	Special     *Special      `json:"special,omitempty"`

	Error string `json:"error,omitempty"`
}
