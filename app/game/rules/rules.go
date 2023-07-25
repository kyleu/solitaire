package rules

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

	MaxPlayers        int `json:"maxPlayers,omitempty"`
	VictoryCondition  any `json:"victoryCondition,omitempty"`
	CardRemovalMethod any `json:"cardRemovalMethod,omitempty"`
	DeckOptions       any `json:"deckOptions,omitempty"`

	Stock       any `json:"stock,omitempty"`
	Waste       any `json:"waste,omitempty"`
	Reserves    any `json:"reserves,omitempty"`
	Cells       any `json:"cells,omitempty"`
	Foundations any `json:"foundations,omitempty"`
	Tableaus    any `json:"tableaus,omitempty"`
	Pyramids    any `json:"pyramids,omitempty"`
	Special     any `json:"special,omitempty"`

	Error string `json:"error,omitempty"`
}
