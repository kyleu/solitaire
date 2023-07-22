//go:build test_all || !func_test

package card_test

import (
	"testing"

	"github.com/kyleu/solitaire/app/card"
)

type cardEqualTest struct {
	k        string
	l        *card.Card
	r        *card.Card
	simple   bool
	succeeds bool
}

func commonCard(faceUp bool) *card.Card {
	return card.NewCard(0, card.RankAce, card.SuitHearts, faceUp)
}

var cardEqualTests = []*cardEqualTest{
	{k: "same", l: commonCard(true), r: commonCard(true), succeeds: true},
	{k: "diffId", l: commonCard(true), r: card.NewCard(1, card.RankAce, card.SuitHearts, true)},
	{k: "diffRank", l: commonCard(true), r: card.NewCard(0, card.RankKing, card.SuitHearts, true)},
	{k: "diffSuit", l: commonCard(true), r: card.NewCard(0, card.RankAce, card.SuitSpades, true)},
	{k: "diffFaceUp", l: commonCard(true), r: card.NewCard(0, card.RankAce, card.SuitHearts, false)},

	{k: "sameSimple", l: commonCard(true), r: commonCard(true), succeeds: true, simple: true},
	{k: "diffFaceUpSimple", l: commonCard(true), r: commonCard(false), succeeds: true, simple: true},
	{k: "diffRank", l: commonCard(true), r: card.NewCard(0, card.RankKing, card.SuitHearts, true), simple: true},
	{k: "diffSuit", l: commonCard(true), r: card.NewCard(0, card.RankAce, card.SuitSpades, true), simple: true},
}

func TestEqual(t *testing.T) {
	t.Parallel()

	for _, tt := range cardEqualTests {
		var result bool
		if tt.simple {
			result = tt.l.EqualsSimple(tt.r)
		} else {
			result = tt.l.Equals(tt.r)
		}
		if tt.succeeds != result {
			if tt.succeeds {
				t.Errorf("%s: cards [%s] and [%s] are not equal", tt.k, tt.l.String(), tt.r.String())
			} else {
				t.Errorf("%s: cards [%s] and [%s] are incorrectly equal", tt.k, tt.l.String(), tt.r.String())
			}
		}
	}
}
