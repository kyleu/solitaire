//go:build js

package main

import (
	"github.com/kyleu/solitaire/app/util"
)

func initWASM(l util.Logger) {
	//randoms := pile.Pile{ID: "test", Options: &pile.Options{}, Cards: card.RandomCards(10, suit.SuitsCommon, 1)}
	//Audit("Hand", util.ToJSON(randoms))
	//outcome := poker.Check(randoms.Cards)
	//Audit("Outcome", util.ToJSON(outcome))
}
