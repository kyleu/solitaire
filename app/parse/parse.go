package parse

import (
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/util"
)

func ParseRulesContent(content string, logger util.Logger) *rules.Rules {
	sections := splitSections(content)
	ret := &rules.Rules{}
	for _, section := range sections {
		k, v := kvFor(section)
		parseKV(k, v, ret, logger)
	}
	return ret
}
