package rules

import (
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

type Reserve struct {
	Name          string `json:"name,omitempty"`
	NumPiles      int    `json:"numPiles,omitempty"`
	InitialCards  int    `json:"initialCards,omitempty"`
	CardsFaceDown int    `json:"cardsFaceDown,omitempty"`
}

func ReserveFromMap(m util.ValueMap, logger util.Logger) *Reserve {
	if len(m) == 0 {
		return nil
	}
	ret := &Reserve{}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "name":
			ret.Name = m.GetStringOpt(k)
		case "numPiles":
			ret.NumPiles = m.GetIntOpt(k)
		case "initialCards":
			ret.InitialCards = m.GetIntOpt(k)
		case "cardsFaceDown":
			ret.CardsFaceDown = m.GetIntOpt(k)
		default:
			logger.Errorf("unhandled reserve options key [%s]", k)
		}
	})
	return ret
}
