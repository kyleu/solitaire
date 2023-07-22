//go:build test_all || !func_test

package card_test

import (
	"github.com/kyleu/solitaire/app/card"
	"testing"
)

type cardEqualTest struct {
	k string
	l *card.Card
	r *card.Card
	x bool
}

var cardEqualTests = []*cardEqualTest{
	{k: "same", l: card.NewCard(card.RankAce, card.SuitHearts), r: card.NewCard(card.RankAce, card.SuitHearts), x: true},
	{k: "diffId", l: card.NewCardID(0, card.RankAce, card.SuitHearts, true), r: card.NewCardID(1, card.RankAce, card.SuitHearts, true)},
	{k: "diffRank", l: card.NewCardID(0, card.RankAce, card.SuitHearts, true), r: card.NewCardID(0, card.RankKing, card.SuitHearts, true)},
	{k: "diffSuit", l: card.NewCardID(0, card.RankAce, card.SuitHearts, true), r: card.NewCardID(0, card.RankAce, card.SuitSpades, true)},
	{k: "diffFaceUp", l: card.NewCardID(0, card.RankAce, card.SuitHearts, true), r: card.NewCardID(0, card.RankAce, card.SuitHearts, false)},
}

func TestEqual(t *testing.T) {
	t.Parallel()

	for _, tt := range cardEqualTests {
		if tt.x != tt.l.Equals(tt.r) {
			if tt.x {
				t.Errorf("%s: cards [%s] and [%s] are not equal", tt.k, tt.l.String(), tt.r.String())
			} else {
				t.Errorf("%s: cards [%s] and [%s] are incorrectly equal", tt.k, tt.l.String(), tt.r.String())
			}
		}
	}
}
