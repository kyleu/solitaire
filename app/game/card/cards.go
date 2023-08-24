package card

import (
	"cmp"
	"slices"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/game/rank"
	"github.com/kyleu/solitaire/app/game/suit"
	"github.com/kyleu/solitaire/app/util"
)

type Cards []*Card

func (c Cards) Ranks() rank.Ranks {
	return lo.Map(c, func(crd *Card, _ int) rank.Rank {
		return crd.Rank
	})
}

func (c Cards) Suits() suit.Suits {
	return lo.Map(c, func(crd *Card, _ int) suit.Suit {
		return crd.Suit
	})
}

func (c Cards) WithRank(r rank.Rank, max int) Cards {
	ret := lo.Filter(c, func(crd *Card, _ int) bool {
		return crd.Rank == r
	})
	if max == 0 || len(ret) <= max {
		return ret
	}
	return ret[:max]
}

func (c Cards) WithSuit(s suit.Suit, max int) Cards {
	ret := lo.Filter(c, func(crd *Card, _ int) bool {
		return crd.Suit == s
	})
	if max == 0 || len(ret) <= max {
		return ret
	}
	return ret[:max]
}

func (c Cards) Sorted() Cards {
	ret := slices.Clone(c)
	slices.SortFunc(ret, func(l *Card, r *Card) int {
		if l.Rank == r.Rank {
			return cmp.Compare(uint8(l.Suit), uint8(r.Suit))
		}
		return cmp.Compare(uint8(l.Rank), uint8(r.Rank))
	})
	return ret
}

func (c Cards) Last() *Card {
	if len(c) == 0 {
		return nil
	}
	return c[len(c)-1]
}

func (c Cards) MaxRank() rank.Rank {
	var ret rank.Rank
	lo.ForEach(c, func(crd *Card, _ int) {
		if uint8(crd.Rank) > uint8(ret) {
			ret = crd.Rank
		}
	})
	return ret
}

func (c Cards) MaxSuit() suit.Suit {
	var ret suit.Suit
	lo.ForEach(c, func(crd *Card, _ int) {
		if uint8(crd.Suit) > uint8(ret) {
			ret = crd.Suit
		}
	})
	return ret
}

func (c Cards) Equals(x Cards) bool {
	if len(c) != len(x) {
		return false
	}
	for idx, l := range c {
		r := x[idx]
		if !l.Equals(r) {
			return false
		}
	}
	return true
}

func (c Cards) EqualsSimple(x Cards) bool {
	if len(c) != len(x) {
		return false
	}
	for idx, l := range c {
		r := x[idx]
		if !l.EqualsSimple(r) {
			return false
		}
	}
	return true
}

func RandomCards(amount int, suits suit.Suits, offset int) Cards {
	return lo.Map(lo.Range(amount), func(i int, _ int) *Card {
		return RandomCard(i+offset, suits)
	})
}

func RandomCard(id int, suits suit.Suits) *Card {
	return &Card{ID: id, Rank: rank.RanksCommon.Random(), Suit: suits.Random(), FaceUp: util.RandomBool()}
}
