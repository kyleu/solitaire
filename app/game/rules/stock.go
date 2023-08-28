package rules

import (
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

type Stock struct {
	Name                  string `json:"name,omitempty"`
	CardsShown            int    `json:"cardsShown,omitempty"`
	DealTo                string `json:"dealTo,omitempty"`
	MaximumDeals          int    `json:"maximumDeals,omitempty"`
	CardsDealt            string `json:"cardsDealt,omitempty"`
	StopAfterPartialDeal  bool   `json:"stopAfterPartialDeal,omitempty"`
	CreatePocketWhenEmpty bool   `json:"createPocketWhenEmpty,omitempty"`
	GalleryMode           bool   `json:"galleryMode,omitempty"`
}

func StockFromMap(m util.ValueMap, logger util.Logger) *Stock {
	if len(m) == 0 {
		return nil
	}
	ret := &Stock{}
	lo.ForEach(lo.Keys(m), func(k string, _ int) {
		switch k {
		case "name":
			ret.Name = m.GetStringOpt(k)
		case "cardsShown":
			ret.CardsShown = m.GetIntOpt(k)
		case "dealTo":
			ret.DealTo = strings.TrimPrefix(m.GetStringOpt(k), "StockDealTo.")
		case "maximumDeals":
			ret.MaximumDeals = m.GetIntOpt(k)
		case "cardsDealt":
			ret.CardsDealt = strings.TrimPrefix(m.GetStringOpt(k), "StockCardsDealt.")
		case "stopAfterPartialDeal":
			ret.StopAfterPartialDeal = m.GetBoolOpt(k)
		case "createPocketWhenEmpty":
			ret.CreatePocketWhenEmpty = m.GetBoolOpt(k)
		case "galleryMode":
			ret.GalleryMode = m.GetBoolOpt(k)
		default:
			logger.Errorf("unhandled stock options key [%s]", k)
		}
	})
	return ret
}
