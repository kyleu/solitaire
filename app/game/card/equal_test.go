//go:build test_all || !func_test

package card_test

import (
	"testing"

	"github.com/kyleu/solitaire/app/game/card"
	"github.com/kyleu/solitaire/app/game/rank"
	"github.com/kyleu/solitaire/app/game/suit"
)

type cardEqualTest struct {
	k        string
	l        *card.Card
	r        *card.Card
	simple   bool
	succeeds bool
}

func commonCard(faceUp bool) *card.Card {
	return card.NewCard(0, rank.RankAce, suit.Hearts, faceUp)
}

var cardEqualTests = []*cardEqualTest{
	{k: "same", l: commonCard(true), r: commonCard(true), succeeds: true},
	{k: "diffId", l: commonCard(true), r: card.NewCard(1, rank.RankAce, suit.Hearts, true)},
	{k: "diffRank", l: commonCard(true), r: card.NewCard(0, rank.RankKing, suit.Hearts, true)},
	{k: "diffSuit", l: commonCard(true), r: card.NewCard(0, rank.RankAce, suit.Spades, true)},
	{k: "diffFaceUp", l: commonCard(true), r: card.NewCard(0, rank.RankAce, suit.Hearts, false)},

	{k: "sameSimple", l: commonCard(true), r: commonCard(true), succeeds: true, simple: true},
	{k: "diffFaceUpSimple", l: commonCard(true), r: commonCard(false), succeeds: true, simple: true},
	{k: "diffRank", l: commonCard(true), r: card.NewCard(0, rank.RankKing, suit.Hearts, true), simple: true},
	{k: "diffSuit", l: commonCard(true), r: card.NewCard(0, rank.RankAce, suit.Spades, true), simple: true},
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
