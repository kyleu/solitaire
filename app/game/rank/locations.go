package rank

import (
	"github.com/kyleu/solitaire/app/util"
)

func (r Rank) Locations() util.Positions {
	switch r {
	case Two:
		return util.Positions{
			{X: 0.5, Y: 0.3}, {X: 0.5, Y: 0.7},
		}
	case Three:
		return util.Positions{
			{X: 0.2, Y: 0.7}, {X: 0.5, Y: 0.5}, {X: 0.8, Y: 0.3},
		}
	case Four:
		return util.Positions{
			{X: 0.2, Y: 0.5}, {X: 0.5, Y: 0.3}, {X: 0.5, Y: 0.7}, {X: 0.8, Y: 0.5},
		}
	case Five:
		return util.Positions{
			{X: 0.2, Y: 0.3}, {X: 0.8, Y: 0.3}, {X: 0.5, Y: 0.5}, {X: 0.2, Y: 0.7}, {X: 0.8, Y: 0.7},
		}
	case Six:
		return util.Positions{
			{X: 0.2, Y: 0.4}, {X: 0.2, Y: 0.6}, {X: 0.5, Y: 0.3}, {X: 0.5, Y: 0.7}, {X: 0.8, Y: 0.4},
			{X: 0.8, Y: 0.6},
		}
	case Seven:
		return util.Positions{
			{X: 0.2, Y: 0.4}, {X: 0.2, Y: 0.6}, {X: 0.5, Y: 0.3}, {X: 0.5, Y: 0.5}, {X: 0.5, Y: 0.7},
			{X: 0.8, Y: 0.4}, {X: 0.8, Y: 0.6},
		}
	case Eight:
		return util.Positions{
			{X: 0.2, Y: 0.3}, {X: 0.2, Y: 0.5}, {X: 0.2, Y: 0.7}, {X: 0.5, Y: 0.4}, {X: 0.5, Y: 0.6},
			{X: 0.8, Y: 0.3}, {X: 0.8, Y: 0.5}, {X: 0.8, Y: 0.7},
		}
	case Nine:
		return util.Positions{
			{X: 0.2, Y: 0.4}, {X: 0.2, Y: 0.6}, {X: 0.2, Y: 0.8}, {X: 0.5, Y: 0.3}, {X: 0.5, Y: 0.5},
			{X: 0.5, Y: 0.7}, {X: 0.8, Y: 0.2}, {X: 0.8, Y: 0.4}, {X: 0.8, Y: 0.6},
		}
	case Ten:
		return util.Positions{
			{X: 0.2, Y: 0.3}, {X: 0.2, Y: 0.5}, {X: 0.2, Y: 0.7}, {X: 0.5, Y: 0.2}, {X: 0.5, Y: 0.4},
			{X: 0.5, Y: 0.6}, {X: 0.5, Y: 0.8}, {X: 0.8, Y: 0.3}, {X: 0.8, Y: 0.5}, {X: 0.8, Y: 0.7},
		}
	default:
		return nil
	}
}
