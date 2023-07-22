package rules

import (
	"github.com/kyleu/solitaire/app/card"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"strings"
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
				if straightFlush.Last() == card.RankAce {
					return newResult(crds, HandTypeRoyalFlush, 0, "Royal Flush ("+str+" of "+s.Name+")")
				}
				return newResult(crds, HandTypeStraightFlush, crds.MaxRank().Index, "Straight Flush ("+str+" of "+s.Name+")")
			}
		}
	}

	if fourOfAKind := hasRankCount(4, ranksCount); fourOfAKind != nil {
		crds := cards.WithRank(fourOfAKind, 4)
		return newResult(crds, HandTypeFourOfaKind, crds.MaxRank().Index, "Four of a kind ("+fourOfAKind.Plural+")")
	}

	threeOfAKind := hasRankCount(3, ranksCount)
	if threeOfAKind != nil {
		if fullHousePair := hasRankCount(2, ranksCount, threeOfAKind); fullHousePair != nil {
			threeCrds := cards.WithRank(threeOfAKind, 3)
			twoCards := cards.WithRank(fullHousePair, 2)
			crds := append(threeCrds, twoCards...)
			msg := "Full house (" + threeOfAKind.Plural + " over " + fullHousePair.Plural + ")"
			return newResult(crds, HandTypeFullHouse, threeCrds.MaxRank().Index, msg)
		}
	}

	if flush := hasSuitCount(5, suitsCount); flush != nil {
		crds := cards.WithSuit(flush, 5)
		return newResult(crds, HandTypeFlush, crds.MaxRank().Index, "Flush ("+flush.Name+")")
	}

	if len(straight) >= 5 {
		crds, str := straightResult(cards, straight)
		return newResult(crds, HandTypeStraight, crds.MaxRank().Index, "Straight ("+str+")")
	}

	if threeOfAKind != nil {
		crds := cards.WithRank(threeOfAKind, 3)
		return newResult(crds, HandTypeThreeOfaKind, crds.MaxRank().Index, "Three of a kind ("+threeOfAKind.Plural+")")
	}

	if pair := hasRankCount(2, ranksCount); pair != nil {
		if twoPair := hasRankCount(2, ranksCount, pair); twoPair != nil {
			crds := append(cards.WithRank(pair, 2), cards.WithRank(twoPair, 2)...)
			msg := "Two pair (" + pair.Plural + " and " + twoPair.Plural + ")"
			return newResult(crds, HandTypeTwoPair, crds.MaxRank().Index, msg)
		}
		crds := cards.WithRank(pair, 2)
		return newResult(crds, HandTypePair, crds.MaxRank().Index, "Pair of "+pair.Plural)
	}

	highRank := ranks[len(ranks)-1]
	crds := cards.WithRank(highRank, 1)
	return newResult(crds, HandTypeHighCard, crds.MaxRank().Index, highRank.Name+" high")
}

func hasSuitCount(count int, suitsCount map[*card.Suit]int, exclude ...*card.Suit) *card.Suit {
	var highest *card.Suit
	for k, v := range suitsCount {
		if (!slices.Contains(exclude, k)) && v >= count && (highest == nil || k.Index > highest.Index) {
			highest = k
		}
	}
	return highest
}

func hasRankCount(count int, ranksCount map[*card.Rank]int, exclude ...*card.Rank) *card.Rank {
	var highest *card.Rank
	for k, v := range ranksCount {
		if (!slices.Contains(exclude, k)) && v >= count && (highest == nil || k.Index > highest.Index) {
			highest = k
		}
	}
	return highest
}

func longestRun(considered card.Ranks) card.Ranks {
	var currentChain card.Ranks
	var longestChain card.Ranks
	lo.ForEach(considered, func(r *card.Rank, _ int) {
		switch {
		case len(currentChain) == 0:
			currentChain = append(currentChain, r)
		case currentChain.Last().Key == r.Key:
			// noop
		case r.Index == currentChain.Last().Index+1:
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

func straightResult(cards card.Cards, straight card.Ranks) (card.Cards, string) {
	if len(straight) > 5 {
		straight = straight[len(straight)-5:]
	}
	crds := lo.Map(straight, func(r *card.Rank, _ int) *card.Card {
		return cards.WithRank(r, 1)[0]
	})
	str := strings.Join(lo.Map(straight, func(item *card.Rank, _ int) string {
		return item.Key
	}), ", ")
	return crds, str
}
