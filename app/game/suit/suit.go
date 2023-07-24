package suit

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/util"
)

type Suit uint8

const (
	Spades Suit = iota + 1
	Hearts
	Diamonds
	Clubs

	Horseshoes
	Stars
	Tridents
	Moons

	Suitless
	Unknown
)

func (s Suit) Key() string {
	switch s {
	case Spades:
		return "S"
	case Hearts:
		return "H"
	case Diamonds:
		return "D"
	case Clubs:
		return "C"
	case Horseshoes:
		return "O"
	case Stars:
		return "R"
	case Tridents:
		return "T"
	case Moons:
		return "M"
	case Suitless:
		return "X"
	case Unknown:
		return "?"
	default:
		return fmt.Sprintf("invalid suit key [%d]", s)
	}
}

func (s Suit) Name() string {
	switch s {
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Horseshoes:
		return "Horseshoes"
	case Stars:
		return "Stars"
	case Tridents:
		return "Tridents"
	case Moons:
		return "Moons"
	case Suitless:
		return "Suitless"
	case Unknown:
		return util.KeyUnknown
	default:
		return fmt.Sprintf("Unknown[%d]", s)
	}
}

func (s Suit) Symbol() string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣ "
	case Horseshoes:
		return "?"
	case Stars:
		return "?"
	case Tridents:
		return "?"
	case Moons:
		return "?"
	case Suitless:
		return "?"
	case Unknown:
		return "?"
	default:
		return "?"
	}
}

func (s Suit) Color() Color {
	switch s {
	case Spades:
		return ColorBlack
	case Hearts:
		return ColorRed
	case Diamonds:
		return ColorRed
	case Clubs:
		return ColorBlack
	case Horseshoes:
		return ColorRed
	case Stars:
		return ColorBlack
	case Tridents:
		return ColorGreen
	case Moons:
		return ColorBlue
	case Suitless:
		return ColorColorless
	case Unknown:
		return ColorUnknown
	default:
		return ColorUnknown
	}
}

func (s Suit) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(s.Key(), false), nil
}

func (s *Suit) UnmarshalJSON(data []byte) error {
	var key string
	err := util.FromJSON(data, &key)
	if err != nil {
		return errors.Wrapf(err, "Invalid suit key [%s]", string(data))
	}
	*s, err = Parse(key)
	return err
}
