package waste

type Waste struct {
	Name          string `json:"name,omitempty"`
	NumPiles      int    `json:"numPiles,omitempty"`
	CardsShown    int    `json:"cardsShown,omitempty"`
	MaxCards      int    `json:"maxCards,omitempty"`
	PlayableCards string `json:"playableCards,omitempty"`
}
