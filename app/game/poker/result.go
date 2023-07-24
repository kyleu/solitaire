package poker

import "github.com/kyleu/solitaire/app/game/card"

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

type Result struct {
	Cards   card.Cards `json:"cards"`
	Score   int        `json:"score"`
	Result  *HandType  `json:"-"`
	Message string     `json:"message"`
}

func newResult(cards card.Cards, res *HandType, addScore int, msg string) *Result {
	return &Result{Cards: cards, Score: (res.Score * 1000) + addScore, Result: res, Message: msg}
}
