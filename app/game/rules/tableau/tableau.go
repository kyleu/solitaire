package tableau

type Tableau struct {
	Name                         string   `json:"name,omitempty"`
	SetNumber                    int      `json:"setNumber,omitempty"`
	NumPiles                     int      `json:"numPiles,omitempty"`
	CardsShown                   int      `json:"cardsShown,omitempty"`
	InitialCards                 any      `json:"initialCards,omitempty"`
	CustomInitialCards           []string `json:"customInitialCards,omitempty"`
	UniqueRanks                  bool     `json:"uniqueRanks,omitempty"`
	CardsFaceDown                string   `json:"cardsFaceDown,omitempty"`
	SuitMatchRuleForBuilding     string   `json:"suitMatchRuleForBuilding,omitempty"`
	RankMatchRuleForBuilding     string   `json:"rankMatchRuleForBuilding,omitempty"`
	Wrap                         bool     `json:"wrap,omitempty"`
	SuitMatchRuleForMovingStacks string   `json:"suitMatchRuleForMovingStacks,omitempty"`
	RankMatchRuleForMovingStacks string   `json:"rankMatchRuleForMovingStacks,omitempty"`

	AutoFillEmptyFrom string `json:"autoFillEmptyFrom,omitempty"`
	EmptyFilledWith   int    `json:"emptyFilledWith,omitempty"`

	MayMoveToNonEmptyFrom []any `json:"mayMoveToNonEmptyFrom,omitempty"`
	MayMoveToEmptyFrom    []any `json:"mayMoveToEmptyFrom,omitempty"`

	MaxCards                  int    `json:"maxCards,omitempty"`
	ActionDuringDeal          string `json:"actionDuringDeal,omitempty"`
	ActionAfterDeal           string `json:"actionAfterDeal,omitempty"`
	PilesWithLowCardsAtBottom int    `json:"pilesWithLowCardsAtBottom,omitempty"`
}
