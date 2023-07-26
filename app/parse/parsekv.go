package parse

import (
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/util"
	"strings"
)

func parseKV(k string, v string, r *rules.Rules, logger util.Logger) {
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
		r.AKA = parseFromStringMap(v)
	case "like":
		r.Like = v
	case "related":
		r.Related = parseFromStringSeq(v)
	case "links":
		r.Links = parseLinks(v)

	case "victoryCondition":
		r.VictoryCondition = strings.TrimPrefix(v, "VictoryCondition.") // todo
	case "cardRemovalMethod":
		r.CardRemovalMethod = strings.TrimPrefix(v, "CardRemovalMethod.") // todo
	case "deckOptions":
		r.DeckOptions = extractMap(extract(v, returnSame, logger)) // todo

	case "stock":
		r.Stock = extractMap(extract(v, returnSame, logger)) // todo
	case "waste":
		r.Waste = extractMap(extract(v, returnSame, logger)) // todo
	case "reserves":
		r.Reserves = extractMap(extract(v, returnSame, logger)) // todo
	case "cells":
		r.Cells = extractMap(extract(v, returnSame, logger)) // todo
	case "foundations":
		r.Foundations = extractDoubleMap(v, logger) // todo
	case "tableaus":
		r.Tableaus = extractDoubleMap(v, logger) // todo
	case "pyramids":
		r.Pyramids = extractDoubleMap(v, logger) // todo
	case "special":
		r.Special = extractMap(extract(v, returnSame, logger)) // todo

	default:
		logger.Errorf("Missing [%s]: %s", k, v)
	}
}
