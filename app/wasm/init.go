//go:build js

package main

import (
	"github.com/kyleu/solitaire/app/card"
	"github.com/kyleu/solitaire/app/pile"
	"github.com/kyleu/solitaire/app/rules"
	"github.com/kyleu/solitaire/app/util"
)

func initWASM(l util.Logger) {
	Audit("x{0}{1}{0}x", false, "y", "z")
	randoms := pile.Pile{ID: "test", Options: &pile.Options{}, Cards: card.RandomCards(10, card.SuitsCommon, 1)}
	Audit(util.ToJSON(randoms), true)
	outcome := rules.CheckPoker(randoms.Cards)
	Audit(util.ToJSON(outcome), true)
}
