package cell

type Cell struct {
	Name          string   `json:"name,omitempty"`
	PluralName    string   `json:"pluralName,omitempty"`
	NumPiles      int      `json:"numPiles,omitempty"`
	MayMoveToFrom []string `json:"mayMoveToFrom,omitempty"`
	InitialCards  int      `json:"initialCards,omitempty"`
	NumEphemeral  int      `json:"numEphemeral,omitempty"`
}
