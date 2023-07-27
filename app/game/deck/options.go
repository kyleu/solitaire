package deck

import (
	"github.com/kyleu/solitaire/app/game/rank"
	"github.com/kyleu/solitaire/app/game/suit"
	"github.com/kyleu/solitaire/app/util"
	"github.com/samber/lo"
)

type Options struct {
	NumDecks int        `json:"numDecks,omitempty"`
	Ranks    rank.Ranks `json:"ranks,omitempty"`
	Suits    suit.Suits `json:"suits,omitempty"`
	LowRank  rank.Rank  `json:"lowRank,omitempty"`
}

func (o *Options) HighRank() rank.Rank {
	return o.LowRank.Previous()
}

func (o *Options) CardCount() int {
	return o.NumDecks * len(o.Suits) * len(o.Ranks)
}

func OptionsFromMap(m util.ValueMap, logger util.Logger) *Options {
	if len(m) == 0 {
		return nil
	}
	ret := &Options{}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "numDecks":
			ret.NumDecks = m.GetIntOpt(k)
		case "ranks":
			ret.Ranks = rank.FromSeqString(m.GetStringOpt(k), logger)
		case "suits":
			ret.Suits = suit.FromSeqString(m.GetStringOpt(k), logger)
		case "lowRank":
			ret.LowRank = rank.FromName(m.GetStringOpt(k))
		default:
			logger.Errorf("unhandled stock options key [%s]", k)
		}
	})
	return ret
}
