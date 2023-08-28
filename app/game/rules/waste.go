package rules

import (
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

type Waste struct {
	Name          string `json:"name,omitempty"`
	NumPiles      int    `json:"numPiles,omitempty"`
	CardsShown    int    `json:"cardsShown,omitempty"`
	MaxCards      int    `json:"maxCards,omitempty"`
	PlayableCards string `json:"playableCards,omitempty"`
}

func WasteFromMap(m util.ValueMap, logger util.Logger) *Waste {
	if len(m) == 0 {
		return nil
	}
	ret := &Waste{}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "name":
			ret.Name = m.GetStringOpt(k)
		case "numPiles":
			ret.NumPiles = m.GetIntOpt(k)
		case "cardsShown":
			ret.CardsShown = m.GetIntOpt(k)
		case "maxCards":
			ret.MaxCards = m.GetIntOpt(k)
		case "playableCards":
			ret.PlayableCards = strings.TrimPrefix(m.GetStringOpt(k), "WastePlayableCards.")
		default:
			logger.Errorf("unhandled waste options key [%s]", k)
		}
	})
	return ret
}
