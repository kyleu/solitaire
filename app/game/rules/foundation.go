package rules

import (
	"github.com/kyleu/solitaire/app/util"
	"github.com/samber/lo"
	"strings"
)

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
	MayMoveToFrom             string `json:"mayMoveToFrom,omitempty"`
	Visible                   bool   `json:"visible,omitempty"`
	AutoMoveCards             bool   `json:"autoMoveCards,omitempty"`
	AutoMoveFrom              string `json:"autoMoveFrom,omitempty"`
}

func FoundationFromMap(m util.ValueMap, logger util.Logger) *Foundation {
	if len(m) == 0 {
		return nil
	}
	ret := &Foundation{}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "name":
			ret.Name = m.GetStringOpt(k)
		case "setNumber":
			ret.SetNumber = m.GetIntOpt(k)
		case "numPiles":
			ret.NumPiles = m.GetIntOpt(k)
		case "cardsShown":
			ret.CardsShown = m.GetIntOpt(k)
		case "lowRank":
			ret.LowRank = strings.TrimPrefix(m.GetStringOpt(k), "FoundationLowRank.")
		case "initialCardRestriction":
			ret.InitialCardRestriction = strings.TrimPrefix(m.GetStringOpt(k), "FoundationInitialCardRestriction.")
		case "initialCards":
			ret.InitialCards = m.GetIntOpt(k)
		case "suitMatchRule":
			ret.SuitMatchRule = strings.TrimPrefix(m.GetStringOpt(k), "SuitMatchRule.")
		case "rankMatchRule":
			ret.RankMatchRule = strings.TrimPrefix(m.GetStringOpt(k), "RankMatchRule.")
		case "wrap":
			ret.Wrap = m.GetBoolOpt(k)
		case "moveCompleteSequencesOnly":
			ret.MoveCompleteSequencesOnly = m.GetBoolOpt(k)
		case "maxCards":
			ret.MaxCards = m.GetIntOpt(k)
		case "canMoveFrom":
			ret.CanMoveFrom = m.GetStringOpt(k)
		case "mayMoveToFrom":
			ret.MayMoveToFrom = m.GetStringOpt(k)
		case "visible":
			ret.Visible = m.GetBoolOpt(k)
		case "autoMoveCards":
			ret.AutoMoveCards = m.GetBoolOpt(k)
		case "autoMoveFrom":
			ret.AutoMoveFrom = m.GetStringOpt(k)
		default:
			logger.Errorf("unhandled foundation options key [%s]", k)
		}
	})
	return ret
}

func FoundationsFromMaps(ms []util.ValueMap, logger util.Logger) []*Foundation {
	return lo.Map(ms, func(m util.ValueMap, index int) *Foundation {
		return FoundationFromMap(m, logger)
	})
}
