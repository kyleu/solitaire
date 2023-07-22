package card

import (
	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/util"
)

type Suit struct {
	Index int
	Key   string
	Name  string
	Color *Color
}

func (s *Suit) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(s.Key, false), nil
}

func (s *Suit) UnmarshalJSON(data []byte) error {
	var key string
	err := util.FromJSON(data, &key)
	if err != nil {
		return errors.Wrapf(err, "Invalid suit key [%s]", string(data))
	}
	curr, ok := suitsByKey[key]
	if !ok {
		return errors.Wrapf(err, "Incorrect suit key [%s]", string(data))
	}
	s.ReplaceFrom(curr)
	return nil
}

func (s *Suit) ReplaceFrom(x *Suit) {
	s.Index = x.Index
	s.Key = x.Key
	s.Name = x.Name
	s.Color = x.Color
}

var (
	SuitHearts   = &Suit{Index: 0, Key: "H", Name: "Hearts", Color: ColorRed}
	SuitSpades   = &Suit{Index: 1, Key: "S", Name: "Spades", Color: ColorBlack}
	SuitDiamonds = &Suit{Index: 2, Key: "D", Name: "Diamond", Color: ColorRed}
	SuitClubs    = &Suit{Index: 3, Key: "C", Name: "Clubs", Color: ColorBlack}

	SuitHorseshoes = &Suit{Index: 4, Key: "O", Name: "Horseshoes", Color: ColorRed}
	SuitStars      = &Suit{Index: 5, Key: "R", Name: "Stars", Color: ColorBlack}
	SuitTridents   = &Suit{Index: 6, Key: "T", Name: "Tridents", Color: ColorGreen}
	SuitMoons      = &Suit{Index: 7, Key: "M", Name: "Moons", Color: ColorBlue}

	SuitSuitless = &Suit{Index: 8, Key: "X", Name: "Suitless", Color: ColorColorless}
	SuitUnknown  = &Suit{Index: -1, Key: "?", Name: "Unknown", Color: ColorUnknown}

	SuitsCommon   = Suits{SuitHearts, SuitSpades, SuitDiamonds, SuitClubs}
	SuitsExpanded = Suits{SuitHearts, SuitSpades, SuitDiamonds, SuitClubs, SuitHorseshoes, SuitStars, SuitTridents, SuitMoons}
)
