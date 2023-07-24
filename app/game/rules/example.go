package rules

var Example = &Rules{
	Key:       "example",
	Completed: true,
	Name:      "Example",
	Layout:    "s",

	AKA:     map[string]string{"testx": "Test X", "testy": "Test Y"},
	Like:    "testz",
	Related: []string{"a", "b", "c"},
	Links:   Links{{Name: "Solitaire", URL: "https://solitaire.kyleu.com"}},

	MaxPlayers:        1,
	VictoryCondition:  nil,
	CardRemovalMethod: nil,
	DeckOptions:       nil,

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
