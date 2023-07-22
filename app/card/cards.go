package card

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"github.com/kyleu/solitaire/app/util"
)

type Cards []*Card

func (c Cards) Ranks() Ranks {
	return lo.Map(c, func(crd *Card, _ int) *Rank {
		return crd.Rank
	})
}

func (c Cards) Suits() Suits {
	return lo.Map(c, func(crd *Card, _ int) *Suit {
		return crd.Suit
	})
}

func (c Cards) WithRank(r *Rank, max int) Cards {
	ret := lo.Filter(c, func(crd *Card, _ int) bool {
		return crd.Rank == r
	})
	if max == 0 || len(ret) <= max {
		return ret
	}
	return ret[:max]
}

func (c Cards) WithSuit(s *Suit, max int) Cards {
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
	slices.SortFunc(ret, func(l *Card, r *Card) bool {
		if l.Rank == r.Rank {
			return l.Suit.Index < r.Suit.Index
		}
		return l.Rank.Index < r.Rank.Index
	})
	return ret
}

func (c Cards) Last() *Card {
	if len(c) == 0 {
		return nil
	}
	return c[len(c)-1]
}

func (c Cards) MaxRank() *Rank {
	var ret *Rank
	lo.ForEach(c, func(crd *Card, _ int) {
		if ret == nil || crd.Rank.Index > ret.Index {
			ret = crd.Rank
		}
	})
	return ret
}

func (c Cards) MaxSuit() *Suit {
	var ret *Suit
	lo.ForEach(c, func(crd *Card, _ int) {
		if ret == nil || crd.Suit.Index > ret.Index {
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

func RandomCards(amount int, suits Suits, offset int) Cards {
	return lo.Map(lo.Range(amount), func(i int, _ int) *Card {
		return RandomCard(i+offset, suits)
	})
}

func RandomCard(id int, suits Suits) *Card {
	return &Card{ID: id, Rank: RanksAll.Random(), Suit: suits.Random(), FaceUp: util.RandomBool()}
}
