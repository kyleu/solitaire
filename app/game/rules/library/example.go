package library

import (
	"github.com/kyleu/solitaire/app/game/deck"
	"github.com/kyleu/solitaire/app/game/rules"
)

var Example = &rules.Rules{
	Key:       "example",
	Completed: true,
	Name:      "Example",
	Layout:    "s",

	AKA:     map[string]string{"testx": "Test X", "testy": "Test Y"},
	Like:    "testz",
	Related: []string{"a", "b", "c"},
	Links:   rules.Links{{Name: "Solitaire", URL: "https://solitaire.kyleu.com"}},

	MaxPlayers:        1,
	VictoryCondition:  "?",
	CardRemovalMethod: "?",
	DeckOptions:       &deck.Options{},

	Stock:       nil,
	Waste:       nil,
	Reserves:    nil,
	Cells:       nil,
	Foundations: nil,
	Tableaus:    nil,
	Pyramids:    nil,
	Special:     nil,

	Error: "",
}
