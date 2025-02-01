package rules

import (
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

type Cell struct {
	Name          string `json:"name,omitempty"`
	PluralName    string `json:"pluralName,omitempty"`
	NumPiles      int    `json:"numPiles,omitempty"`
	MayMoveToFrom string `json:"mayMoveToFrom,omitempty"`
	InitialCards  int    `json:"initialCards,omitempty"`
	NumEphemeral  int    `json:"numEphemeral,omitempty"`
}

func CellFromMap(m util.ValueMap, logger util.Logger) *Cell {
	if len(m) == 0 {
		return nil
	}
	ret := &Cell{}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "name":
			ret.Name = m.GetStringOpt(k)
		case "pluralName":
			ret.PluralName = m.GetStringOpt(k)
		case "numPiles":
			ret.NumPiles = m.GetIntOpt(k)
		case "mayMoveToFrom":
			ret.MayMoveToFrom = m.GetStringOpt(k)
		case "initialCards":
			ret.InitialCards = m.GetIntOpt(k)
		case "numEphemeral":
			ret.NumEphemeral = m.GetIntOpt(k)
		default:
			logger.Errorf("unhandled cell options key [%s]", k)
		}
	})
	return ret
}
