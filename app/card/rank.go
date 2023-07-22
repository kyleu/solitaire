package card

import (
	"github.com/kyleu/solitaire/app/util"
	"github.com/pkg/errors"
)

type Rank struct {
	Index  int
	Key    string
	Name   string
	Plural string
}

func newRank(idx int, key string, name string) *Rank {
	return &Rank{Index: idx, Key: key, Name: name, Plural: util.StringToPlural(name)}
}

func (s *Rank) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(s.Key, false), nil
}

func (s *Rank) UnmarshalJSON(data []byte) error {
	var key string
	err := util.FromJSON(data, &key)
	if err != nil {
		return errors.Wrapf(err, "Invalid rank key [%s]", string(data))
	}
	curr, ok := ranksByKey[key]
	if !ok {
		return errors.Wrapf(err, "Incorrect rank key [%s]", string(data))
	}
	s = curr
	return nil
}

var (
	RankTwo     = newRank(2, "2", "Two")
	RankThree   = newRank(3, "3", "Three")
	RankFour    = newRank(4, "4", "Four")
	RankFive    = newRank(5, "5", "Five")
	RankSix     = newRank(6, "6", "Six")
	RankSeven   = newRank(7, "7", "Seven")
	RankEight   = newRank(8, "8", "Eight")
	RankNine    = newRank(9, "9", "Nine")
	RankTen     = newRank(10, "X", "Ten")
	RankJack    = newRank(11, "J", "Jack")
	RankQueen   = newRank(12, "Q", "Queen")
	RankKing    = newRank(13, "K", "King")
	RankAce     = newRank(14, "A", "Ace")
	RankUnknown = newRank(0, "?", "Unknown")
)
