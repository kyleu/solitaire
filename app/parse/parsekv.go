package parse

import (
	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/util"
)

func parseKV(k string, v string, r *rules.Rules, logger util.Logger) error {
	switch k {
	case "id":
		r.Key = v
	case "completed":
		r.Completed = v == "true"
	case "title":
		if v != "" {
			r.Name = v
		}
	case "layout":
		r.Layout = v

	case "aka":
		r.AKA = map[string]string{"X": v} // todo
	case "like":
		r.Like = v // todo
	case "related":
		r.Related = parseFromStringSeq(v) // todo
	case "links":
		r.Links = rules.Links{{Name: v}} // todo

	case "victoryCondition":
		r.VictoryCondition = v // todo
	case "cardRemovalMethod":
		r.CardRemovalMethod = v // todo
	case "deckOptions":
		r.DeckOptions = v // todo

	case "stock":
		r.Stock = v // todo
	case "waste":
		r.Waste = v // todo
	case "reserves":
		r.Reserves = v // todo
	case "cells":
		r.Cells = v // todo
	case "foundations":
		r.Foundations = v // todo
	case "tableaus":
		r.Tableaus = v // todo
	case "pyramids":
		r.Pyramids = v // todo
	case "special":
		r.Special = v // todo

	default:
		return errors.Errorf("[%s]: %s", k, v)
	}
	return nil
}
