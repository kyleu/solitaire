package rank

import (
	"fmt"

	"github.com/kyleu/solitaire/app/util"
)

func (r Rank) Key() string {
	switch r {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "X"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	case Unknown:
		return "?"
	default:
		return "?"
	}
}

func (r Rank) Name() string {
	switch r {
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	case Unknown:
		return util.KeyUnknown
	default:
		return fmt.Sprintf("Unknown[%d]", r)
	}
}

func (r Rank) Plural() string {
	switch r {
	case Two:
		return "Twos"
	case Three:
		return "Threes"
	case Four:
		return "Fours"
	case Five:
		return "Fives"
	case Six:
		return "Sixes"
	case Seven:
		return "Sevens"
	case Eight:
		return "Eights"
	case Nine:
		return "Nines"
	case Ten:
		return "Tens"
	case Jack:
		return "Jacks"
	case Queen:
		return "Queens"
	case King:
		return "Kings"
	case Ace:
		return "Aces"
	case Unknown:
		return util.KeyUnknown
	default:
		return fmt.Sprintf("Unknown[%d]", r)
	}
}
