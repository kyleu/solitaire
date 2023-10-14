package rules

import (
	"github.com/kyleu/solitaire/app/util"
	"github.com/samber/lo"
)

type Pyramid struct {
	Name  string        `json:"name,omitempty"`
	Extra util.ValueMap `json:"extra"`
}

func PyramidFromMap(m util.ValueMap, logger util.Logger) *Pyramid {
	if len(m) == 0 {
		return nil
	}
	ret := &Pyramid{Extra: m}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "name":
			ret.Name = m.GetStringOpt(k)
		default:
			// logger.Errorf("unhandled pyramid options key [%s]", k)
		}
	})
	return ret
}

func PyramidsFromMaps(ms []util.ValueMap, logger util.Logger) []*Pyramid {
	return lo.Map(ms, func(m util.ValueMap, index int) *Pyramid {
		return PyramidFromMap(m, logger)
	})
}
