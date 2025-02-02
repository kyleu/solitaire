package deck

import (
	"github.com/kyleu/solitaire/app/game/rank"
	"github.com/kyleu/solitaire/app/game/suit"
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
