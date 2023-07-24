package pile

import (
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/game/card"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Set struct {
	Behavior   any       `json:"behavior"`
	Piles      Piles     `json:"piles"`
	Hidden     bool      `json:"hidden,omitempty"`
	Position   *Position `json:"position,omitempty"`
	Rows       float64   `json:"rows,omitempty"`
	Dimensions *Position `json:"dimensions,omitempty"`
	Divisor    float64   `json:"divisor,omitempty"`
}

func (s *Set) Visible() bool {
	return !s.Hidden
}

type Sets []*Set

func (s Sets) Piles() Piles {
	return lo.FlatMap(s, func(x *Set, _ int) []*Pile {
		return s.Piles()
	})
}

func (s Sets) PilesByID() map[string]*Pile {
	return lo.SliceToMap(lo.FlatMap(s, func(x *Set, _ int) []*Pile {
		return s.Piles()
	}), func(p *Pile) (string, *Pile) {
		return p.ID, p
	})
}

func (s Sets) CardsByID() map[int]*card.Card {
	piles := lo.FlatMap(s, func(x *Set, _ int) []*Pile {
		return s.Piles()
	})
	return lo.SliceToMap(lo.FlatMap(piles, func(x *Pile, _ int) []*card.Card {
		return x.Cards
	}), func(c *card.Card) (int, *card.Card) {
		return c.ID, c
	})
}
