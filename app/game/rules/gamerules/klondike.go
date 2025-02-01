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
		{Name: "Wikipedia", URL: "en.wikipedia.org/wiki/Klondike_(solitaire)"},
		{Name: "Solitaire Central", URL: "www.solitairecentral.com/rules/Klondike.html"},
		{Name: "Pretty Good Solitaire", URL: "www.goodsol.com/pgshelp/klondike.htm"},
		{Name: "dogMelon", URL: "www.dogmelon.com.au/solhelp/Klondike%20Solitaire.shtml"},
		{Name: "Xolitaire", URL: "www.escapedivision.com/xolitaire/en/games/klondike.html"},
		{Name: "Lady Cadogan's Illustrated Games of Solitaire or Patience", URL: "www.gutenberg.org/files/21642/21642-h/21642-h.htm#canfield"},
		{Name: "An 1897 description", URL: "howtoplaysolitaire.blogspot.com/2010/06/canfield-or-klondike-single-deck.html"},
		{Name: "Robert Abbott's Strategy Guide", URL: "www.logicmazes.com/sol/"},
		{Name: "Boris Sandberg's Strategy Tips", URL: "www.solitairecentral.com/articles/KlondikeSolitaireWinningStrategyTips.html"},
		{Name: "A Strategy for Winning Klondike Solitaire", URL: "www.jupiterscientific.org/sciinfo/AStrategryForWinningKlondikeSolitaire.html"},
		{Name: "Usman Latif's analysis of the number of Klondike cards in which no cards can be played.", URL: "www.techuser.net/klondikeprob.html"},
		{Name: "Bill's Solitaire Tester", URL: "www.roziturnbull.com/bill/Solitaire/solitaire.htm"},
	},

	Stock:       &rules.Stock{},
	Waste:       &rules.Waste{},
	Foundations: []*rules.Foundation{{NumPiles: 4, AutoMoveCards: true}},
	Tableaus:    []*rules.Tableau{{Extra: util.ValueMap{"emptyFilledWith": "FillEmptyWith.HighRank"}}},

	Error: "",
}
