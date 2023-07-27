package parse

import (
	"github.com/kyleu/solitaire/app/game/deck"
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/parse/extract"
	"github.com/kyleu/solitaire/app/util"
	"strings"
)

func RulesContent(content string, logger util.Logger) *rules.Rules {
	sections := extract.ExtractSections(content)
	ret := &rules.Rules{}
	for _, section := range sections {
		k, v := extract.KVFor(section)
		kv(k, v, ret, logger)
	}
	return ret
}

func kv(k string, v string, r *rules.Rules, logger util.Logger) {
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
		r.AKA = extract.ParseFromStringMap(v)
	case "like":
		r.Like = v
	case "related":
		r.Related = extract.ParseFromStringSeq(v)
	case "links":
		r.Links = parseLinks(v)

	case "victoryCondition":
		r.VictoryCondition = strings.TrimPrefix(v, "VictoryCondition.")
	case "cardRemovalMethod":
		r.CardRemovalMethod = strings.TrimPrefix(v, "CardRemovalMethod.")
	case "deckOptions":
		r.DeckOptions = deck.OptionsFromMap(extract.ExtractMap(v, logger), logger)

	case "stock":
		r.Stock = rules.StockFromMap(extract.ExtractMap(v, logger), logger)
	case "waste":
		r.Waste = rules.WasteFromMap(extract.ExtractMap(v, logger), logger)
	case "reserves":
		r.Reserves = rules.ReserveFromMap(extract.ExtractMap(v, logger), logger)
	case "cells":
		r.Cells = rules.CellFromMap(extract.ExtractMap(v, logger), logger)
	case "foundations":
		r.Foundations = rules.FoundationsFromMaps(extract.ExtractDoubleMap(v, logger), logger)
	case "tableaus":
		r.Tableaus = rules.TableausFromMaps(extract.ExtractDoubleMap(v, logger), logger)
	case "pyramids":
		r.Pyramids = rules.PyramidsFromMaps(extract.ExtractDoubleMap(v, logger), logger)
	case "special":
		r.Special = rules.SpecialFromMap(extract.ExtractMap(v, logger), logger)

	default:
		logger.Errorf("Missing [%s]: %s", k, v)
	}
}
