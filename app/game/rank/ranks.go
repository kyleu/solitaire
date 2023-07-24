package rank

import (
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

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
		return Unknown
	}
	return r[len(r)-1]
}

var (
	rankMap = map[string]Rank{
		"2": Two,
		"3": Three,
		"4": Four,
		"5": Five,
		"6": Six,
		"7": Seven,
		"8": Eight,
		"9": Nine,
		"x": Ten,
		"j": Jack,
		"q": Queen,
		"k": King,
		"a": Ace,
		"?": Unknown,
	}

	RanksAll = Ranks{
		Two, Three, Four, Five, Six, Seven, Eight,
		Nine, Ten, Jack, Queen, King, Ace,
	}
)

func Parse(key string) (Rank, error) {
	key = strings.TrimSpace(strings.ToLower(key))
	ret, ok := rankMap[key]
	if !ok {
		return Unknown, errors.Errorf("invalid suit [%s]", key)
	}
	return ret, nil
}
