package constraint

import (
	"github.com/kyleu/solitaire/app/card"
	"github.com/kyleu/solitaire/app/pile"
	"github.com/kyleu/solitaire/app/util"
	"github.com/samber/lo"
)

type Fn func(src *pile.Pile, tgt pile.Pile, crds card.Cards, state any) bool

type Constraint struct {
	ID   string        `json:"id"`
	Fn   Fn            `json:"-"`
	Opts util.ValueMap `json:"opts,omitempty"`
}

func newConstraint(id string, fn Fn, opts util.ValueMap) *Constraint {
	return &Constraint{ID: id, Fn: fn, Opts: opts}
}

func Always(id string) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, _ pile.Pile, _ card.Cards, _ any) bool {
		return true
	}, nil)
}

func Never(id string) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, _ pile.Pile, _ card.Cards, _ any) bool {
		return false
	}, nil)
}

func Empty(id string) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, tgt pile.Pile, _ card.Cards, _ any) bool {
		return tgt.Empty()
	}, nil)
}

func AllOf(id string, constraints ...*Constraint) *Constraint {
	return newConstraint(id, func(src *pile.Pile, tgt pile.Pile, cards card.Cards, state any) bool {
		return lo.EveryBy(constraints, func(c *Constraint) bool {
			return c.Fn(src, tgt, cards, state)
		})
	}, nil)
}

// TODO
func FiniteTimes(id string, count int) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, _ pile.Pile, _ card.Cards, _ any) bool {
		_ = count // TODO
		return true
	}, nil)
}

func FaceUp(id string) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, _ pile.Pile, cards card.Cards, _ any) bool {
		return !lo.ContainsBy(cards, func(c *card.Card) bool {
			return !c.FaceUp
		})
	}, nil)
}

func FaceDown(id string) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, _ pile.Pile, cards card.Cards, _ any) bool {
		return !lo.ContainsBy(cards, func(c *card.Card) bool {
			return c.FaceUp
		})
	}, nil)
}

func SpecificRanks(id string, ranks ...card.Rank) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, _ pile.Pile, cards card.Cards, _ any) bool {
		return lo.EveryBy(cards, func(c *card.Card) bool {
			return lo.ContainsBy(ranks, func(r card.Rank) bool {
				return r.Key == c.Rank.Key
			})
		})
	}, nil)
}

func AllEmpty(id string, piles ...*pile.Pile) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, _ pile.Pile, cards card.Cards, _ any) bool {
		return lo.EveryBy(piles, func(p *pile.Pile) bool {
			return p.Empty()
		})
	}, nil)
}

func AllNonEmpty(id string, piles ...*pile.Pile) *Constraint {
	return newConstraint(id, func(_ *pile.Pile, _ pile.Pile, cards card.Cards, _ any) bool {
		return lo.EveryBy(piles, func(p *pile.Pile) bool {
			return !p.Empty()
		})
	}, nil)
}
