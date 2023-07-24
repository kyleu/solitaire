package poker

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"github.com/kyleu/solitaire/app/game/card"
	"github.com/kyleu/solitaire/app/game/rank"
)

func PokerCheck(cards card.Cards) *PokerResult {
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
