package parse

import (
	"github.com/kyleu/solitaire/app/util"
	"strings"
)

func kvFor(section string) (string, string) {
	k, v := util.StringSplit(section, '=', true)
	k, v = strings.TrimSpace(k), cleanLine(v)
	v = strings.TrimSuffix(strings.TrimPrefix(v, `"`), `"`)
	return k, v
}

func cleanLine(l string) string {
	l = strings.TrimSuffix(strings.TrimSpace(l), ",")
	if strings.HasPrefix(l, "Some(") {
		l = strings.TrimPrefix(strings.TrimSuffix(l, ")"), "Some(")
	}
	return l
}

func parseFromStringSeq(v string) []string {
	var ret []string
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
		ret = append(ret, v[:end])
		v = v[end+1:]
	}
	return ret
}

func parseFromStringMap(v string) map[string]string {
	ret := map[string]string{}
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
			ret[key] = v[:end]
			key = ""
		}
		v = v[end+1:]
	}
	return ret
}
