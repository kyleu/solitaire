package reserve

type Reserve struct {
	Name          string `json:"name,omitempty"`
	NumPiles      int    `json:"numPiles,omitempty"`
	InitialCards  int    `json:"initialCards,omitempty"`
	CardsFaceDown int    `json:"cardsFaceDown,omitempty"`
}
