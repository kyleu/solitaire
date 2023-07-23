package rules

import (
	"github.com/kyleu/solitaire/app/rank"
	"github.com/kyleu/solitaire/app/suit"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"github.com/kyleu/solitaire/app/card"
)

type HandType struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var (
	HandTypeHighCard      = &HandType{Name: "High card", Score: 1}
	HandTypePair          = &HandType{Name: "Pair", Score: 2}
	HandTypeTwoPair       = &HandType{Name: "Two pair", Score: 3}
	HandTypeThreeOfaKind  = &HandType{Name: "Three of a kind", Score: 4}
	HandTypeStraight      = &HandType{Name: "Straight", Score: 5}
	HandTypeFlush         = &HandType{Name: "Flush", Score: 6}
	HandTypeFullHouse     = &HandType{Name: "Full house", Score: 7}
	HandTypeFourOfaKind   = &HandType{Name: "Four of a kind", Score: 8}
	HandTypeStraightFlush = &HandType{Name: "Straight flush", Score: 9}
	HandTypeRoyalFlush    = &HandType{Name: "Royal flush", Score: 10}
)

type PokerResult struct {
	Cards   card.Cards `json:"cards"`
	Score   int        `json:"score"`
	Result  *HandType  `json:"-"`
	Message string     `json:"message"`
}

func newResult(cards card.Cards, res *HandType, addScore int, msg string) *PokerResult {
	return &PokerResult{Cards: cards, Score: (res.Score * 10000) + addScore, Result: res, Message: msg}
}

func CheckPoker(cards card.Cards) *PokerResult {
	ranks := cards.Ranks().Sorted()
	ranksCount := lo.CountValues(ranks)

	suits := cards.Suits().Sorted()
	suitsCount := lo.CountValues(suits)

	straight := longestRun(ranks)
	if len(straight) >= 5 {
		for _, s := range suits {
			if straightFlush := longestRun(cards.WithSuit(s, 0).Ranks()); len(straightFlush) >= 5 {
				crds, str := straightResult(cards, straightFlush)
				if uint8(straightFlush.Last()) == uint8(rank.RankAce) {
					return newResult(crds, HandTypeRoyalFlush, 0, "Royal Flush ("+str+" of "+s.Name()+")")
				}
				return newResult(crds, HandTypeStraightFlush, int(crds.MaxRank()), "Straight Flush ("+str+" of "+s.Name()+")")
			}
		}
	}

	if fourOfAKind := hasRankCount(4, ranksCount); fourOfAKind != 0 {
		crds := cards.WithRank(fourOfAKind, 4)
		return newResult(crds, HandTypeFourOfaKind, int(crds.MaxRank()), "Four of a kind ("+fourOfAKind.Plural()+")")
	}

	threeOfAKind := hasRankCount(3, ranksCount)
	if threeOfAKind != 0 {
		if fullHousePair := hasRankCount(2, ranksCount, threeOfAKind); fullHousePair != 0 {
			threeCrds := cards.WithRank(threeOfAKind, 3)
			twoCards := cards.WithRank(fullHousePair, 2)
			crds := append(slices.Clone(threeCrds), twoCards...)
			msg := "Full house (" + threeOfAKind.Plural() + " over " + fullHousePair.Plural() + ")"
			return newResult(crds, HandTypeFullHouse, int(threeCrds.MaxRank()), msg)
		}
	}

	if flush := hasSuitCount(5, suitsCount); flush != 0 {
		crds := cards.WithSuit(flush, 5)
		return newResult(crds, HandTypeFlush, int(crds.MaxRank()), "Flush ("+flush.Name()+")")
	}

	if len(straight) >= 5 {
		crds, str := straightResult(cards, straight)
		return newResult(crds, HandTypeStraight, int(crds.MaxRank()), "Straight ("+str+")")
	}

	if threeOfAKind != 0 {
		crds := cards.WithRank(threeOfAKind, 3)
		return newResult(crds, HandTypeThreeOfaKind, int(crds.MaxRank()), "Three of a kind ("+threeOfAKind.Plural()+")")
	}

	if pair := hasRankCount(2, ranksCount); pair != 0 {
		if twoPair := hasRankCount(2, ranksCount, pair); twoPair != 0 {
			crds := append(cards.WithRank(pair, 2), cards.WithRank(twoPair, 2)...)
			msg := "Two pair (" + pair.Plural() + " and " + twoPair.Plural() + ")"
			return newResult(crds, HandTypeTwoPair, int(crds.MaxRank()), msg)
		}
		crds := cards.WithRank(pair, 2)
		return newResult(crds, HandTypePair, int(crds.MaxRank()), "Pair of "+pair.Plural())
	}

	highRank := ranks[len(ranks)-1]
	crds := cards.WithRank(highRank, 1)
	return newResult(crds, HandTypeHighCard, int(crds.MaxRank()), highRank.Name()+" high")
}

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
