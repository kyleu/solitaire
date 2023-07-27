package suit

import (
	"github.com/kyleu/solitaire/app/parse/extract"
	"github.com/samber/lo"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/kyleu/solitaire/app/util"
)

type Suits []Suit

func (s Suits) Sorted() Suits {
	ret := slices.Clone(s)
	slices.SortFunc(ret, func(l Suit, r Suit) bool {
		return uint8(l) < uint8(r)
	})
	return ret
}

func (s Suits) Random() Suit {
	return s[util.RandomInt(len(s))]
}

var (
	SuitsCommon   = Suits{Hearts, Spades, Diamonds, Clubs}
	SuitsExpanded = Suits{Hearts, Spades, Diamonds, Clubs, Horseshoes, Stars, Tridents, Moons}
	SuitsAll      = Suits{Hearts, Spades, Diamonds, Clubs, Horseshoes, Stars, Tridents, Moons, Suitless, Unknown}
	suitMap       = map[string]Suit{
		"s": Spades, "spades": Spades,
		"h": Hearts, "hearts": Hearts,
		"d": Diamonds, "diamonds": Diamonds,
		"c": Clubs, "clubs": Clubs,
		"o": Horseshoes, "horseshoes": Horseshoes,
		"r": Stars, "stars": Stars,
		"t": Tridents, "tridents": Tridents,
		"m": Moons, "moons": Moons,
		"x": Suitless, "suitless": Suitless,
		"?": Unknown, "unknown": Unknown,
	}
)

func Parse(key string) (Suit, error) {
	key = strings.TrimSpace(strings.ToLower(key))
	ret, ok := suitMap[key]
	if !ok {
		return Suit(0), errors.Errorf("invalid suit [%s]", key)
	}
	return ret, nil
}

func FromName(name string) Suit {
	name = strings.TrimSpace(strings.TrimPrefix(name, "Suit."))
	return lo.FindOrElse(SuitsAll, Unknown, func(r Suit) bool {
		return r.Name() == name || r.Key() == name
	})
}

func FromSeqString(s string, logger util.Logger) Suits {
	if strings.TrimPrefix(s, "Suit.") == "all" {
		return SuitsExpanded
	}
	return extract.Extract(s, func(x string) Suit {
		return FromName(strings.TrimPrefix(x, "Suit."))
	}, logger)
}
