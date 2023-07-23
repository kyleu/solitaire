package deck

import (
	"github.com/kyleu/solitaire/app/rank"
	"github.com/kyleu/solitaire/app/suit"
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/card"
)

type Deck struct {
	Cards card.Cards `json:"cards"`
}

func (d *Deck) Equals(x *Deck) bool {
	return d.Cards.Equals(x.Cards)
}

func (d *Deck) EqualsSimple(x *Deck) bool {
	return d.Cards.EqualsSimple(x.Cards)
}

func (d *Deck) Empty() bool {
	return len(d.Cards) == 0
}

func FreshDeck(numLoops int, ranks rank.Ranks, suits suit.Suits, faceUp bool, idOffset int) *Deck {
	if len(ranks) == 0 {
		ranks = rank.RanksAll
	}
	if len(suits) == 0 {
		suits = suit.SuitsCommon
	}
	id := 0
	cards := lo.FlatMap(lo.Range(numLoops), func(_ int, loopIdx int) []*card.Card {
		return lo.FlatMap(suits, func(s suit.Suit, suitIdx int) []*card.Card {
			return lo.Map(ranks, func(r rank.Rank, rankIdx int) *card.Card {
				id++
				return card.NewCard(idOffset+id, r, s, faceUp)
			})
		})
	})
	return &Deck{Cards: cards}
}
