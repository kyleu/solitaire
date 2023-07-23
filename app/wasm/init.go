//go:build js

package main

import (
	"github.com/kyleu/solitaire/app/card"
	"github.com/kyleu/solitaire/app/pile"
	"github.com/kyleu/solitaire/app/rules"
	"github.com/kyleu/solitaire/app/suit"
	"github.com/kyleu/solitaire/app/util"
)

func initWASM(l util.Logger) {
	randoms := pile.Pile{ID: "test", Options: &pile.Options{}, Cards: card.RandomCards(10, suit.SuitsCommon, 1)}
	Audit("Hand", util.ToJSON(randoms))
	outcome := rules.CheckPoker(randoms.Cards)
	Audit("Outcome", util.ToJSON(outcome))
}
