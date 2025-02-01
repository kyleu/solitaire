package sandbox

import (
	"context"
	"github.com/kyleu/solitaire/app/game/deck"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/lib/filesystem"
	"github.com/kyleu/solitaire/app/parse"
	"github.com/kyleu/solitaire/app/util"
)

var parseRules = &Sandbox{Key: "parseRules", Title: "Parse Rules", Icon: "heart", Run: onParseRules}

func onParseRules(ctx context.Context, as *app.State, logger util.Logger) (any, error) {
	return onParseRulesJSON(ctx, as, logger)
	//return onParseRulesScala(ctx, as, logger)
}

type JSONLink struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type JSONFoundation struct {
	Name                      string `json:"name,omitempty"`
	SetNumber                 int    `json:"setNumber,omitempty"`
	NumPiles                  int    `json:"numPiles,omitempty"`
	CardsShown                int    `json:"cardsShown,omitempty"`
	LowRank                   int    `json:"lowRank,omitempty"`
	InitialCardRestriction    int    `json:"initialCardRestriction,omitempty"`
	InitialCards              int    `json:"initialCards,omitempty"`
	SuitMatchRule             string `json:"suitMatchRule,omitempty"`
	RankMatchRule             string `json:"rankMatchRule,omitempty"`
	Wrap                      bool   `json:"wrap,omitempty"`
	MoveCompleteSequencesOnly bool   `json:"moveCompleteSequencesOnly,omitempty"`
	MaxCards                  int    `json:"maxCards,omitempty"`
	CanMoveFrom               int    `json:"canMoveFrom,omitempty"`
	MayMoveToFrom             any    `json:"mayMoveToFrom,omitempty"`
	Visible                   bool   `json:"visible,omitempty"`
	AutoMoveCards             bool   `json:"autoMoveCards,omitempty"`
	AutoMoveFrom              any    `json:"autoMoveFrom,omitempty"`
}

type JSONTableau struct {
	Name                         string   `json:"name,omitempty"`
	SetNumber                    int      `json:"setNumber,omitempty"`
	NumPiles                     int      `json:"numPiles,omitempty"`
	CardsShown                   int      `json:"cardsShown,omitempty"`
	InitialCards                 any      `json:"initialCards,omitempty"`
	CustomInitialCards           []string `json:"customInitialCards,omitempty"`
	UniqueRanks                  bool     `json:"uniqueRanks,omitempty"`
	CardsFaceDown                string   `json:"cardsFaceDown,omitempty"`
	SuitMatchRuleForBuilding     string   `json:"suitMatchRuleForBuilding,omitempty"`
	RankMatchRuleForBuilding     string   `json:"rankMatchRuleForBuilding,omitempty"`
	Wrap                         bool     `json:"wrap,omitempty"`
	SuitMatchRuleForMovingStacks string   `json:"suitMatchRuleForMovingStacks,omitempty"`
	RankMatchRuleForMovingStacks string   `json:"rankMatchRuleForMovingStacks,omitempty"`

	AutoFillEmptyFrom string `json:"autoFillEmptyFrom,omitempty"`
	EmptyFilledWith   int    `json:"emptyFilledWith,omitempty"`

	MayMoveToNonEmptyFrom []any `json:"mayMoveToNonEmptyFrom,omitempty"`
	MayMoveToEmptyFrom    []any `json:"mayMoveToEmptyFrom,omitempty"`

	MaxCards                  int    `json:"maxCards,omitempty"`
	ActionDuringDeal          string `json:"actionDuringDeal,omitempty"`
	ActionAfterDeal           string `json:"actionAfterDeal,omitempty"`
	PilesWithLowCardsAtBottom int    `json:"pilesWithLowCardsAtBottom,omitempty"`
}

type JSONCell struct {
	Name          string   `json:"name,omitempty"`
	PluralName    string   `json:"pluralName,omitempty"`
	NumPiles      int      `json:"numPiles,omitempty"`
	MayMoveToFrom []string `json:"mayMoveToFrom,omitempty"`
	InitialCards  int      `json:"initialCards,omitempty"`
	NumEphemeral  int      `json:"numEphemeral,omitempty"`
}

type JSONRules struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed,omitempty"`
	Title     string `json:"title"`
	Layout    string `json:"layout"`

	AKA     map[string]string `json:"aka,omitempty"`
	Like    string            `json:"like,omitempty"`
	Related []string          `json:"related,omitempty"`
	Links   []*JSONLink       `json:"links,omitempty"`

	MaxPlayers        int           `json:"maxPlayers,omitempty"`
	VictoryCondition  string        `json:"victoryCondition,omitempty"`
	CardRemovalMethod int           `json:"cardRemovalMethod,omitempty"`
	DeckOptions       *deck.Options `json:"deckOptions,omitempty"`

	Stock       *rules.Stock      `json:"stock,omitempty"`
	Waste       *rules.Waste      `json:"waste,omitempty"`
	Reserves    *rules.Reserve    `json:"reserves,omitempty"`
	Cells       *JSONCell         `json:"cells,omitempty"`
	Foundations []*JSONFoundation `json:"foundations,omitempty"`
	Tableaus    []*JSONTableau    `json:"tableaus,omitempty"`
	Pyramids    []*JSONPyramid    `json:"pyramids,omitempty"`
	Special     *rules.Special    `json:"special,omitempty"`

	Error string `json:"error,omitempty"`
}

type JSONPyramid struct {
	Name          string `json:"name,omitempty"`
	SetNumber     int    `json:"setNumber,omitempty"`
	PyramidType   int    `json:"pyramidType,omitempty"`
	Height        int    `json:"height,omitempty"`
	CardsFaceDown int    `json:"cardsFaceDown,omitempty"`

	SuitMatchRuleForBuilding     string `json:"suitMatchRuleForBuilding,omitempty"`
	RankMatchRuleForBuilding     string `json:"rankMatchRuleForBuilding,omitempty"`
	Wrap                         bool   `json:"wrap,omitempty"`
	SuitMatchRuleForMovingStacks string `json:"suitMatchRuleForMovingStacks,omitempty"`
	RankMatchRuleForMovingStacks string `json:"rankMatchRuleForMovingStacks,omitempty"`

	MayMoveToNonEmptyFrom []string `json:"mayMoveToNonEmptyFrom,omitempty"`
	MayMoveToEmptyFrom    []string `json:"mayMoveToEmptyFrom,omitempty"`
	EmptyFilledWith       int      `json:"emptyFilledWith,omitempty"`
}

func onParseRulesJSON(_ context.Context, _ *app.State, logger util.Logger) (any, error) {
	fs, err := filesystem.NewFileSystem("./tmp", false, "")
	if err != nil {
		return nil, err
	}
	b, err := fs.ReadFile("rules.json")
	if err != nil {
		return nil, err
	}
	var ret []*JSONRules
	err = util.FromJSONStrict(b, &ret)
	return ret, err
}

func onParseRulesScala(_ context.Context, _ *app.State, logger util.Logger) (any, error) {
	fs, err := filesystem.NewFileSystem("../solitaire.gg/shared/src/main/scala/models/rules/impl", false, "")
	if err != nil {
		return nil, err
	}
	files := lo.Map(fs.ListFiles(".", nil, logger), func(e *filesystem.FileInfo, _ int) string {
		return e.Name
	})
	ret := lo.Map(files, func(filename string, _ int) *rules.Rules {
		b, err := fs.ReadFile(filename)
		if err != nil {
			name := strings.TrimSuffix(filename, ".scala")
			return &rules.Rules{Name: name, Error: err.Error()}
		}
		content := string(b)
		return parse.RulesContent(content, logger)
	})
	return ret, nil
}
