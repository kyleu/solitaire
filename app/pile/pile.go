package pile

import "github.com/kyleu/solitaire/app/card"

type Pile struct {
	ID      string     `json:"id"`
	Options *Options   `json:"options,omitempty"`
	Cards   card.Cards `json:"cards,omitempty"`
}

func (p *Pile) Empty() bool {
	return len(p.Cards) == 0
}
