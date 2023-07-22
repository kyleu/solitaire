package card

import (
	"github.com/kyleu/solitaire/app/util"
	"golang.org/x/exp/slices"
)

type Ranks []*Rank

func (r Ranks) Sorted() Ranks {
	ret := slices.Clone(r)
	slices.SortFunc(ret, func(l *Rank, r *Rank) bool {
		return l.Index < r.Index
	})
	return ret
}

func (r Ranks) Random() *Rank {
	return r[util.RandomInt(len(r))]
}

func (r Ranks) Last() *Rank {
	if len(r) == 0 {
		return nil
	}
	return r[len(r)-1]
}

var (
	ranksByKey = map[string]*Rank{
		RankTwo.Key:     RankTwo,
		RankThree.Key:   RankThree,
		RankFour.Key:    RankFour,
		RankFive.Key:    RankFive,
		RankSix.Key:     RankSix,
		RankSeven.Key:   RankSeven,
		RankEight.Key:   RankEight,
		RankNine.Key:    RankNine,
		RankTen.Key:     RankTen,
		RankJack.Key:    RankJack,
		RankQueen.Key:   RankQueen,
		RankKing.Key:    RankKing,
		RankAce.Key:     RankAce,
		RankUnknown.Key: RankUnknown,
	}

	RanksAll = Ranks{
		RankTwo, RankThree, RankFour, RankFive, RankSix, RankSeven, RankEight,
		RankNine, RankTen, RankJack, RankQueen, RankKing, RankAce,
	}
)
