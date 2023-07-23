package rank

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/util"
)

type Rank uint8

const (
	RankTwo Rank = iota + 1
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
	RankUnknown
)

func (r Rank) Key() string {
	switch r {
	case RankTwo:
		return "2"
	case RankThree:
		return "3"
	case RankFour:
		return "4"
	case RankFive:
		return "5"
	case RankSix:
		return "6"
	case RankSeven:
		return "7"
	case RankEight:
		return "8"
	case RankNine:
		return "9"
	case RankTen:
		return "X"
	case RankJack:
		return "J"
	case RankQueen:
		return "Q"
	case RankKing:
		return "K"
	case RankAce:
		return "A"
	case RankUnknown:
		return "?"
	default:
		return "?"
	}
}

func (r Rank) Name() string {
	switch r {
	case RankTwo:
		return "Two"
	case RankThree:
		return "Three"
	case RankFour:
		return "Four"
	case RankFive:
		return "Five"
	case RankSix:
		return "Six"
	case RankSeven:
		return "Seven"
	case RankEight:
		return "Eight"
	case RankNine:
		return "Nine"
	case RankTen:
		return "Ten"
	case RankJack:
		return "Jack"
	case RankQueen:
		return "Queen"
	case RankKing:
		return "King"
	case RankAce:
		return "Ace"
	case RankUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("Unknown[%d]", r)
	}
}

func (r Rank) Plural() string {
	switch r {
	case RankTwo:
		return "Twos"
	case RankThree:
		return "Threes"
	case RankFour:
		return "Fours"
	case RankFive:
		return "Fives"
	case RankSix:
		return "Sixes"
	case RankSeven:
		return "Sevens"
	case RankEight:
		return "Eights"
	case RankNine:
		return "Nines"
	case RankTen:
		return "Tens"
	case RankJack:
		return "Jacks"
	case RankQueen:
		return "Queens"
	case RankKing:
		return "Kings"
	case RankAce:
		return "Aces"
	case RankUnknown:
		return "Unknown"
	default:
		return fmt.Sprintf("Unknown[%d]", r)
	}
}

func (r Rank) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(r.Key, false), nil
}

func (r *Rank) UnmarshalJSON(data []byte) error {
	var key string
	err := util.FromJSON(data, &key)
	if err != nil {
		return errors.Wrapf(err, "Invalid suit key [%s]", string(data))
	}
	*r, err = Parse(key)
	return nil
}
