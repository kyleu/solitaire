package rank

import (
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"strings"

	"github.com/kyleu/solitaire/app/util"
)

type Ranks []Rank

func (r Ranks) Sorted() Ranks {
	ret := slices.Clone(r)
	slices.SortFunc(ret, func(l Rank, r Rank) bool {
		return uint8(l) < uint8(r)
	})
	return ret
}

func (r Ranks) Random() Rank {
	return r[util.RandomInt(len(r))]
}

func (r Ranks) Last() Rank {
	if len(r) == 0 {
		return RankUnknown
	}
	return r[len(r)-1]
}

var (
	rankMap = map[string]Rank{
		"2": RankTwo,
		"3": RankThree,
		"4": RankFour,
		"5": RankFive,
		"6": RankSix,
		"7": RankSeven,
		"8": RankEight,
		"9": RankNine,
		"x": RankTen,
		"j": RankJack,
		"q": RankQueen,
		"k": RankKing,
		"a": RankAce,
		"?": RankUnknown,
	}

	RanksAll = Ranks{
		RankTwo, RankThree, RankFour, RankFive, RankSix, RankSeven, RankEight,
		RankNine, RankTen, RankJack, RankQueen, RankKing, RankAce,
	}
)

func Parse(key string) (Rank, error) {
	key = strings.TrimSpace(strings.ToLower(key))
	ret, ok := rankMap[key]
	if !ok {
		return RankUnknown, errors.Errorf("invalid suit [%s]", key)
	}
	return ret, nil
}
