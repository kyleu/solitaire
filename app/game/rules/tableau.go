package rules

import (
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

type Tableau struct {
	Name  string        `json:"name,omitempty"`
	Extra util.ValueMap `json:"extra"`
}

func TableauFromMap(m util.ValueMap, logger util.Logger) *Tableau {
	if len(m) == 0 {
		return nil
	}
	ret := &Tableau{Extra: m}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "name":
			ret.Name = m.GetStringOpt(k)
		default:
			// logger.Errorf("unhandled tableau options key [%s]", k)
		}
	})
	return ret
}

func TableausFromMaps(ms []util.ValueMap, logger util.Logger) []*Tableau {
	return lo.Map(ms, func(m util.ValueMap, index int) *Tableau {
		return TableauFromMap(m, logger)
	})
}
