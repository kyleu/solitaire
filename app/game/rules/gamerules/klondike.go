package gamerules

import (
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/util"
)

var Klondike = &rules.Rules{
	Key:       "klondike",
	Completed: true,
	Name:      "Klondike",
	Layout:    "swf|t",

	AKA:  map[string]string{"testx": "Test X", "testy": "Test Y"},
	Like: "testz",
	Related: []string{
		"whitehorse", "kingsley", "trigon", "goldmine", "thoughtful", "klondikegallery", "chineseklondike", "athena",
		"saratoga", "endlessharp", "smokey", "spike", "gilbert", "jumboklondike", "chinaman", "klondike1card", "sevendevils",
	},
	Links: rules.Links{
		{"Wikipedia", "en.wikipedia.org/wiki/Klondike_(solitaire)"},
		{"Solitaire Central", "www.solitairecentral.com/rules/Klondike.html"},
		{"Pretty Good Solitaire", "www.goodsol.com/pgshelp/klondike.htm"},
		{"dogMelon", "www.dogmelon.com.au/solhelp/Klondike%20Solitaire.shtml"},
		{"Xolitaire", "www.escapedivision.com/xolitaire/en/games/klondike.html"},
		{"Lady Cadogan's Illustrated Games of Solitaire or Patience", "www.gutenberg.org/files/21642/21642-h/21642-h.htm#canfield"},
		{"An 1897 description", "howtoplaysolitaire.blogspot.com/2010/06/canfield-or-klondike-single-deck.html"},
		{"Robert Abbott's Strategy Guide", "www.logicmazes.com/sol/"},
		{"Boris Sandberg's Strategy Tips", "www.solitairecentral.com/articles/KlondikeSolitaireWinningStrategyTips.html"},
		{"A Strategy for Winning Klondike Solitaire", "www.jupiterscientific.org/sciinfo/AStrategryForWinningKlondikeSolitaire.html"},
		{"Usman Latif's analysis of the number of Klondike cards in which no cards can be played.", "www.techuser.net/klondikeprob.html"},
		{"Bill's Solitaire Tester", "www.roziturnbull.com/bill/Solitaire/solitaire.htm"},
	},

	Stock:       &rules.Stock{},
	Waste:       &rules.Waste{},
	Foundations: []*rules.Foundation{{NumPiles: 4, AutoMoveCards: true}},
	Tableaus:    []*rules.Tableau{{Extra: util.ValueMap{"emptyFilledWith": "FillEmptyWith.HighRank"}}},

	Error: "",
}
