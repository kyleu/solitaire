package extract

import (
	"github.com/kyleu/solitaire/app/util"
	"strconv"
	"strings"
)

func KVFor(section string) (string, string) {
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

func ParseFromStringSeq(v string) []string {
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

func ParseFromStringMap(v string) map[string]string {
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

func parseIntFrom(s string, prefix string, logger util.Logger) int {
	if prefix != "" {
		s = strings.TrimPrefix(strings.TrimSuffix(s, ")"), prefix+"(")
	}
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		logger.Errorf("invalid int from [%s]", s)
	}
	return int(i)
}
