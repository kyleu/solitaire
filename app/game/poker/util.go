package poker

import (
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"github.com/kyleu/solitaire/app/game/card"
	"github.com/kyleu/solitaire/app/game/rank"
	"github.com/kyleu/solitaire/app/game/suit"
)

func hasSuitCount(count int, suitsCount map[suit.Suit]int, exclude ...suit.Suit) suit.Suit {
	var highest suit.Suit
	for k, v := range suitsCount {
		if (!slices.Contains(exclude, k)) && v >= count && (uint8(k) > uint8(highest)) {
			highest = k
		}
	}
	return highest
}

func hasRankCount(count int, ranksCount map[rank.Rank]int, exclude ...rank.Rank) rank.Rank {
	var highest rank.Rank
	for k, v := range ranksCount {
		if (!slices.Contains(exclude, k)) && v >= count && (uint8(k) > uint8(highest)) {
			highest = k
		}
	}
	return highest
}

func longestRun(considered rank.Ranks) rank.Ranks {
	var currentChain rank.Ranks
	var longestChain rank.Ranks
	lo.ForEach(considered, func(r rank.Rank, _ int) {
		switch {
		case len(currentChain) == 0:
			currentChain = append(currentChain, r)
		case uint8(currentChain.Last()) == uint8(r):
			// noop
		case uint8(r) == uint8(currentChain.Last())+1:
			currentChain = append(currentChain, r)
		default:
			if len(currentChain) > len(longestChain) {
				longestChain = currentChain
			}
			currentChain = nil
		}
	})
	if len(currentChain) > len(longestChain) {
		longestChain = currentChain
	}
	return longestChain
}

func straightResult(cards card.Cards, straight rank.Ranks) (card.Cards, string) {
	if len(straight) > 5 {
		straight = straight[len(straight)-5:]
	}
	crds := lo.Map(straight, func(r rank.Rank, _ int) *card.Card {
		return cards.WithRank(r, 1)[0]
	})
	str := strings.Join(lo.Map(straight, func(item rank.Rank, _ int) string {
		return item.Key()
	}), ", ")
	return crds, str
}
