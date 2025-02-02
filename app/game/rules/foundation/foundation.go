package foundation

type Foundation struct {
	Name                      string `json:"name,omitempty"`
	SetNumber                 int    `json:"setNumber,omitempty"`
	NumPiles                  int    `json:"numPiles,omitempty"`
	CardsShown                int    `json:"cardsShown,omitempty"`
	LowRank                   string `json:"lowRank,omitempty"`
	InitialCardRestriction    string `json:"initialCardRestriction,omitempty"`
	InitialCards              int    `json:"initialCards,omitempty"`
	SuitMatchRule             string `json:"suitMatchRule,omitempty"`
	RankMatchRule             string `json:"rankMatchRule,omitempty"`
	Wrap                      bool   `json:"wrap,omitempty"`
	MoveCompleteSequencesOnly bool   `json:"moveCompleteSequencesOnly,omitempty"`
	MaxCards                  int    `json:"maxCards,omitempty"`
	CanMoveFrom               string `json:"canMoveFrom,omitempty"`
	MayMoveToFrom             []any  `json:"mayMoveToFrom,omitempty"`
	Visible                   bool   `json:"visible,omitempty"`
	AutoMoveCards             bool   `json:"autoMoveCards,omitempty"`
	AutoMoveFrom              []any  `json:"autoMoveFrom,omitempty"`
}
