package rules

import (
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

type Special struct {
	Name                string `json:"name,omitempty"`
	RedealsAllowed      int    `json:"redealsAllowed,omitempty"`
	RotationsAllowed    int    `json:"rotationsAllowed,omitempty"`
	RotationTopToBottom bool   `json:"rotationTopToBottom,omitempty"`
	DrawsAllowed        int    `json:"drawsAllowed,omitempty"`
	DrawsAfterRedeals   bool   `json:"drawsAfterRedeals,omitempty"`
}

func SpecialFromMap(m util.ValueMap, logger util.Logger) *Special {
	if len(m) == 0 {
		return nil
	}
	ret := &Special{}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "name":
			ret.Name = m.GetStringOpt(k)
		case "redealsAllowed":
			ret.RedealsAllowed = m.GetIntOpt(k)
		case "rotationsAllowed":
			ret.RotationsAllowed = m.GetIntOpt(k)
		case "rotationTopToBottom":
			ret.RotationTopToBottom = m.GetBoolOpt(k)
		case "drawsAllowed":
			ret.DrawsAllowed = m.GetIntOpt(k)
		case "drawsAfterRedeals":
			ret.DrawsAfterRedeals = m.GetBoolOpt(k)
		default:
			logger.Errorf("unhandled special options key [%s]", k)
		}
	})
	return ret
}
