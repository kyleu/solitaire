package parse

import (
	"github.com/kyleu/solitaire/app/game/rules"
	"github.com/kyleu/solitaire/app/util"
	"strings"
)

func ParseRulesContent(content string, logger util.Logger) *rules.Rules {
	sections := splitSections(content)
	ret := &rules.Rules{}
	for _, section := range sections {
		k, v := util.StringSplit(section, '=', true)
		k, v = strings.TrimSpace(k), cleanLine(v)
		v = strings.TrimSuffix(strings.TrimPrefix(v, `"`), `"`)
		err := parseKV(k, v, ret, logger)
		if err != nil {
			ret.Error += err.Error()
		}
	}
	return ret
}
