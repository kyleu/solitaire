package pyramid

type Pyramid struct {
	Name          string `json:"name,omitempty"`
	SetNumber     int    `json:"setNumber,omitempty"`
	PyramidType   int    `json:"pyramidType,omitempty"`
	Height        int    `json:"height,omitempty"`
	CardsFaceDown int    `json:"cardsFaceDown,omitempty"`

	SuitMatchRuleForBuilding     string `json:"suitMatchRuleForBuilding,omitempty"`
	RankMatchRuleForBuilding     string `json:"rankMatchRuleForBuilding,omitempty"`
	Wrap                         bool   `json:"wrap,omitempty"`
	SuitMatchRuleForMovingStacks string `json:"suitMatchRuleForMovingStacks,omitempty"`
	RankMatchRuleForMovingStacks string `json:"rankMatchRuleForMovingStacks,omitempty"`

	MayMoveToNonEmptyFrom []string `json:"mayMoveToNonEmptyFrom,omitempty"`
	MayMoveToEmptyFrom    []string `json:"mayMoveToEmptyFrom,omitempty"`
	EmptyFilledWith       int      `json:"emptyFilledWith,omitempty"`
}
