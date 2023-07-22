package card

import (
	"golang.org/x/exp/slices"

	"github.com/kyleu/solitaire/app/util"
)

type Suits []*Suit

func (s Suits) Sorted() Suits {
	ret := slices.Clone(s)
	slices.SortFunc(ret, func(l *Suit, r *Suit) bool {
		return l.Index < r.Index
	})
	return ret
}

func (s Suits) Random() *Suit {
	return s[util.RandomInt(len(s))]
}

var suitsByKey = map[string]*Suit{
	SuitHearts.Key:   SuitHearts,
	SuitSpades.Key:   SuitSpades,
	SuitDiamonds.Key: SuitDiamonds,
	SuitClubs.Key:    SuitClubs,

	SuitHorseshoes.Key: SuitHorseshoes,
	SuitStars.Key:      SuitStars,
	SuitTridents.Key:   SuitTridents,
	SuitMoons.Key:      SuitMoons,

	SuitSuitless.Key: SuitSuitless,
	SuitUnknown.Key:  SuitUnknown,
}
