//go:build test_all || !func_test

package card_test

import (
	"fmt"
	"testing"

	"github.com/kyleu/solitaire/app/game/card"
	"github.com/kyleu/solitaire/app/game/rank"
	"github.com/kyleu/solitaire/app/game/suit"
	"github.com/kyleu/solitaire/app/util"
)

type cardSerializeTest struct {
	k string
	c *card.Card
}

var cardSerializeTests = []*cardSerializeTest{
	{k: "AH", c: card.NewCard(0, rank.Ace, suit.Hearts, false)},
	{k: "AH+", c: card.NewCard(0, rank.Ace, suit.Hearts, true)},
	{k: "1:AH", c: card.NewCard(1, rank.Ace, suit.Hearts, false)},
	{k: "1:AH+", c: card.NewCard(1, rank.Ace, suit.Hearts, true)},
	{k: "KS", c: card.NewCard(0, rank.King, suit.Spades, false)},
}

func TestSerialize(t *testing.T) {
	t.Parallel()

	for _, tt := range cardSerializeTests {
		s := tt.c.String()
		if s != tt.k {
			t.Errorf("card [%s] expected [%s]", tt.c.String(), s)
		}
		j := util.ToJSONBytes(tt.c, false)
		if string(j) != fmt.Sprintf("%q", tt.k) {
			t.Errorf("card [%s] serialized to [%s], expected [%s]", tt.c.String(), j, tt.k)
		}
	}
}
