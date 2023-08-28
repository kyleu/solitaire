package parse

import (
	"strings"

	"github.com/kyleu/solitaire/app/game/rules"
)

func parseLinks(v string) rules.Links {
	ret := rules.Links{}
	var key string
	for {
		start := strings.Index(v, `"`)
		if start == -1 {
			break
		}
		v = v[start+1:]
		end := strings.Index(v, `"`)
		if end == -1 {
			break
		}
		if key == "" {
			key = v[:end]
		} else {
			ret = append(ret, &rules.Link{Name: key, URL: v[:end]})
			key = ""
		}
		v = v[end+1:]
	}
	return ret
}
