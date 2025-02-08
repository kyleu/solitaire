package parse

import (
	"fmt"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/game/deck"
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/game/rules/cell"
	"github.com/kyleu/solitaire/app/game/rules/foundation"
	"github.com/kyleu/solitaire/app/game/rules/pyramid"
	"github.com/kyleu/solitaire/app/game/rules/reserve"
	"github.com/kyleu/solitaire/app/game/rules/stock"
	"github.com/kyleu/solitaire/app/game/rules/tableau"
	"github.com/kyleu/solitaire/app/game/rules/waste"
)

type JSONRules struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed,omitempty"`
	Title     string `json:"title"`
	Layout    string `json:"layout"`

	AKA     map[string]string `json:"aka,omitempty"`
	Like    string            `json:"like,omitempty"`
	Related []string          `json:"related,omitempty"`
	Links   []*JSONLink       `json:"links,omitempty"`

	MaxPlayers        int           `json:"maxPlayers,omitempty"`
	VictoryCondition  string        `json:"victoryCondition,omitempty"`
	CardRemovalMethod int           `json:"cardRemovalMethod,omitempty"`
	DeckOptions       *deck.Options `json:"deckOptions,omitempty"`

	Stock       *stock.Stock      `json:"stock,omitempty"`
	Waste       *waste.Waste      `json:"waste,omitempty"`
	Reserves    *reserve.Reserve  `json:"reserves,omitempty"`
	Cells       *JSONCell         `json:"cells,omitempty"`
	Foundations []*JSONFoundation `json:"foundations,omitempty"`
	Tableaus    []*JSONTableau    `json:"tableaus,omitempty"`
	Pyramids    []*JSONPyramid    `json:"pyramids,omitempty"`
	Special     *rules.Special    `json:"special,omitempty"`

	Error string `json:"error,omitempty"`
}

func (j *JSONRules) ToRules() *rules.Rules {
	return &rules.Rules{
		Key:               j.ID,
		Completed:         j.Completed,
		Name:              j.Title,
		Layout:            j.Layout,
		AKA:               j.AKA,
		Like:              j.Like,
		Related:           j.Related,
		Links:             lo.Map(j.Links, func(j *JSONLink, _ int) *rules.Link { return j.ToLink() }),
		MaxPlayers:        j.MaxPlayers,
		VictoryCondition:  j.VictoryCondition,
		CardRemovalMethod: fmt.Sprint(j.CardRemovalMethod),
		DeckOptions:       j.DeckOptions,
		Stock:             j.Stock,
		Waste:             j.Waste,
		Reserves:          j.Reserves,
		Cells:             j.Cells.ToCell(),
		Foundations:       lo.Map(j.Foundations, func(j *JSONFoundation, _ int) *foundation.Foundation { return j.ToFoundation() }),
		Tableaus:          lo.Map(j.Tableaus, func(j *JSONTableau, _ int) *tableau.Tableau { return j.ToTableau() }),
		Pyramids:          lo.Map(j.Pyramids, func(j *JSONPyramid, _ int) *pyramid.Pyramid { return j.ToPyramid() }),
	}
}

type JSONLink struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func (j *JSONLink) ToLink() *rules.Link {
	if j == nil {
		return nil
	}
	return &rules.Link{Name: j.Title, URL: j.URL}
}

type JSONFoundation struct {
	Name                      string `json:"name,omitempty"`
	SetNumber                 int    `json:"setNumber,omitempty"`
	NumPiles                  int    `json:"numPiles,omitempty"`
	CardsShown                int    `json:"cardsShown,omitempty"`
	LowRank                   int    `json:"lowRank,omitempty"`
	InitialCardRestriction    int    `json:"initialCardRestriction,omitempty"`
	InitialCards              int    `json:"initialCards,omitempty"`
	SuitMatchRule             string `json:"suitMatchRule,omitempty"`
	RankMatchRule             string `json:"rankMatchRule,omitempty"`
	Wrap                      bool   `json:"wrap,omitempty"`
	MoveCompleteSequencesOnly bool   `json:"moveCompleteSequencesOnly,omitempty"`
	MaxCards                  int    `json:"maxCards,omitempty"`
	CanMoveFrom               int    `json:"canMoveFrom,omitempty"`
	MayMoveToFrom             []any  `json:"mayMoveToFrom,omitempty"`
	Visible                   bool   `json:"visible,omitempty"`
	AutoMoveCards             bool   `json:"autoMoveCards,omitempty"`
	AutoMoveFrom              []any  `json:"autoMoveFrom,omitempty"`
}

func (j *JSONFoundation) ToFoundation() *foundation.Foundation {
	if j == nil {
		return nil
	}
	return &foundation.Foundation{
		Name:                      j.Name,
		SetNumber:                 j.SetNumber,
		NumPiles:                  j.NumPiles,
		CardsShown:                j.CardsShown,
		LowRank:                   fmt.Sprint(j.LowRank),
		InitialCardRestriction:    fmt.Sprint(j.InitialCardRestriction),
		InitialCards:              j.InitialCards,
		SuitMatchRule:             j.SuitMatchRule,
		RankMatchRule:             j.RankMatchRule,
		Wrap:                      j.Wrap,
		MoveCompleteSequencesOnly: j.MoveCompleteSequencesOnly,
		MaxCards:                  j.MaxCards,
		CanMoveFrom:               fmt.Sprint(j.CanMoveFrom),
		MayMoveToFrom:             j.MayMoveToFrom,
		Visible:                   j.Visible,
		AutoMoveCards:             j.AutoMoveCards,
		AutoMoveFrom:              j.AutoMoveFrom,
	}
}

type JSONTableau struct {
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

func (j *JSONTableau) ToTableau() *tableau.Tableau {
	if j == nil {
		return nil
	}
	return &tableau.Tableau{
		Name:                         j.Name,
		SetNumber:                    j.SetNumber,
		NumPiles:                     j.NumPiles,
		CardsShown:                   j.CardsShown,
		InitialCards:                 j.InitialCards,
		CustomInitialCards:           j.CustomInitialCards,
		UniqueRanks:                  j.UniqueRanks,
		CardsFaceDown:                j.CardsFaceDown,
		SuitMatchRuleForBuilding:     j.SuitMatchRuleForBuilding,
		RankMatchRuleForBuilding:     j.RankMatchRuleForBuilding,
		Wrap:                         j.Wrap,
		SuitMatchRuleForMovingStacks: j.SuitMatchRuleForMovingStacks,
		RankMatchRuleForMovingStacks: j.RankMatchRuleForMovingStacks,
		AutoFillEmptyFrom:            j.AutoFillEmptyFrom,
		EmptyFilledWith:              j.EmptyFilledWith,
		MayMoveToNonEmptyFrom:        j.MayMoveToNonEmptyFrom,
		MayMoveToEmptyFrom:           j.MayMoveToEmptyFrom,
		MaxCards:                     j.MaxCards,
		ActionDuringDeal:             j.ActionDuringDeal,
		ActionAfterDeal:              j.ActionAfterDeal,
		PilesWithLowCardsAtBottom:    j.PilesWithLowCardsAtBottom,
	}
}

type JSONCell struct {
	Name          string   `json:"name,omitempty"`
	PluralName    string   `json:"pluralName,omitempty"`
	NumPiles      int      `json:"numPiles,omitempty"`
	MayMoveToFrom []string `json:"mayMoveToFrom,omitempty"`
	InitialCards  int      `json:"initialCards,omitempty"`
	NumEphemeral  int      `json:"numEphemeral,omitempty"`
}

func (j *JSONCell) ToCell() *cell.Cell {
	if j == nil {
		return nil
	}
	return &cell.Cell{
		Name:          j.Name,
		PluralName:    j.PluralName,
		NumPiles:      j.NumPiles,
		MayMoveToFrom: j.MayMoveToFrom,
		InitialCards:  j.InitialCards,
		NumEphemeral:  j.NumEphemeral,
	}
}

type JSONPyramid struct {
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

func (j *JSONPyramid) ToPyramid() *pyramid.Pyramid {
	if j == nil {
		return nil
	}
	return &pyramid.Pyramid{
		Name:                         j.Name,
		SetNumber:                    j.SetNumber,
		PyramidType:                  j.PyramidType,
		Height:                       j.Height,
		CardsFaceDown:                j.CardsFaceDown,
		SuitMatchRuleForBuilding:     j.SuitMatchRuleForBuilding,
		RankMatchRuleForBuilding:     j.RankMatchRuleForBuilding,
		Wrap:                         j.Wrap,
		SuitMatchRuleForMovingStacks: j.SuitMatchRuleForMovingStacks,
		RankMatchRuleForMovingStacks: j.RankMatchRuleForMovingStacks,
		MayMoveToNonEmptyFrom:        j.MayMoveToNonEmptyFrom,
		MayMoveToEmptyFrom:           j.MayMoveToEmptyFrom,
		EmptyFilledWith:              j.EmptyFilledWith,
	}
}
