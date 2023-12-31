package card

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/game/rank"
	"github.com/kyleu/solitaire/app/game/suit"
)

type Card struct {
	ID     int       `json:"id"`
	Rank   rank.Rank `json:"r"`
	Suit   suit.Suit `json:"s"`
	FaceUp bool      `json:"u,omitempty"`
}

func NewCardOld(r rank.Rank, s suit.Suit) *Card {
	return &Card{Rank: r, Suit: s}
}

func NewCard(id int, r rank.Rank, s suit.Suit, u bool) *Card {
	return &Card{ID: id, Rank: r, Suit: s, FaceUp: u}
}

func FromCardString(str string) (*Card, error) {
	ret := &Card{}

	if idx := strings.Index(str, ":"); idx > -1 {
		idStr := str[:idx]
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid card id [%s]", idStr)
		}
		ret.ID = int(id)
		str = str[idx+1:]
	}
	if str == "" {
		return nil, errors.New("Empty card string")
	}
	l := len(str)
	if l == 1 || l > 3 {
		return nil, errors.Errorf("Invalid card [%s]", str)
	}

	r, err := rank.Parse(str[0:1])
	if err != nil {
		return nil, errors.Wrapf(err, "Invalid card rank [%s]", str[0:1])
	}
	ret.Rank = r

	s, err := suit.Parse(str[1:2])
	if err != nil {
		return nil, errors.Wrapf(err, "Invalid card suit [%s]", str[1:2])
	}
	ret.Suit = s

	if l == 3 {
		ret.FaceUp = str[2] == '+'
	}

	return ret, nil
}

func (c *Card) Equals(x *Card) bool {
	return c.ID == x.ID && c.Rank == x.Rank && c.Suit == x.Suit && c.FaceUp == x.FaceUp
}

func (c *Card) EqualsSimple(x *Card) bool {
	return c.Rank == x.Rank && c.Suit == x.Suit
}

func (c *Card) String() string {
	var ret string
	if c.ID != 0 {
		ret += fmt.Sprintf("%d:", c.ID)
	}
	ret += fmt.Sprintf("%s%s", c.Rank.Key(), c.Suit.Key())
	if c.FaceUp {
		ret += "+"
	}
	return ret
}

func (c *Card) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *Card) UnmarshalText(data []byte) error {
	curr, err := FromCardString(string(data))
	if err != nil {
		return err
	}
	*c = *curr
	return nil
}

func (c *Card) ReplaceFrom(x *Card) {
	c.ID = x.ID
	c.Rank = x.Rank
	c.Suit = x.Suit
	c.FaceUp = x.FaceUp
}

func (c *Card) Unicode() string {
	return UnicodeCards[c.Rank.Key()+c.Suit.Key()]
}
