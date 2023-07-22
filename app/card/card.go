package card

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/util"
)

type Card struct {
	ID     int   `json:"id"`
	Rank   *Rank `json:"r"`
	Suit   *Suit `json:"s"`
	FaceUp bool  `json:"u,omitempty"`
}

func NewCardOld(r *Rank, s *Suit) *Card {
	return &Card{Rank: r, Suit: s}
}

func NewCard(id int, r *Rank, s *Suit, u bool) *Card {
	return &Card{ID: id, Rank: r, Suit: s, FaceUp: u}
}

func FromCardString(str string) (*Card, error) {
	ret := &Card{}

	if idx := strings.Index(str, ":"); idx > -1 {
		idStr := str[:idx]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid card id [%s]", idStr)
		}
		ret.ID = id
		str = str[idx+1:]
	}
	if str == "" {
		return nil, errors.New("Empty card string")
	}
	l := len(str)
	if l == 1 || l > 3 {
		return nil, errors.Errorf("Invalid card [%s]", str)
	}

	r, ok := ranksByKey[str[0:1]]
	if !ok {
		return nil, errors.Errorf("Invalid card rank [%s]", str[0:1])
	}
	ret.Rank = r

	s, ok := suitsByKey[str[1:2]]
	if !ok {
		return nil, errors.Errorf("Invalid card suit [%s]", str[1:2])
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
	ret += fmt.Sprintf("%s%s", c.Rank.Key, c.Suit.Key)
	if c.FaceUp {
		ret += "+"
	}
	return ret
}

func (c *Card) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(c.String(), false), nil
}

func (c *Card) UnmarshalJSON(data []byte) error {
	var str string
	err := util.FromJSON(data, &str)
	if err != nil {
		return errors.Wrapf(err, "Invalid card [%s]", string(data))
	}
	curr, err := FromCardString(str)
	if err != nil {
		return err
	}
	c.ReplaceFrom(curr)
	return nil
}

func (c *Card) ReplaceFrom(x *Card) {
	c.ID = x.ID
	c.Rank = x.Rank
	c.Suit = x.Suit
	c.FaceUp = x.FaceUp
}
