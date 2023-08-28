package extract

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

func returnSame(s string) string {
	return s
}

func Extract[T any](v string, f func(s string) T, logger util.Logger) []T {
	x, err := parseExtract(v)
	if err != nil {
		logger.Errorf("error extracting [%s]: %+v", v, err)
	}
	return lo.Map(x, func(s string, _ int) T {
		return f(s)
	})
}

func ExtractDouble[T any](v string, f func(s string) T, logger util.Logger) [][]T {
	return lo.Map(Extract(v, returnSame, logger), func(x string, _ int) []T {
		return Extract(x, f, logger)
	})
}

func ExtractMap(v string, logger util.Logger) util.ValueMap {
	return lo.SliceToMap(Extract(v, returnSame, logger), func(s string) (string, any) {
		return KVFor(s)
	})
}

func ExtractDoubleMap(v string, logger util.Logger) []util.ValueMap {
	return lo.Map(Extract(v, returnSame, logger), func(x string, _ int) util.ValueMap {
		return ExtractMap(x, logger)
	})
}

func parseExtract(v string) ([]string, error) {
	var ret []string

	lIdx, rIdx := strings.Index(v, "("), strings.LastIndex(v, ")")
	if lIdx == -1 || rIdx == -1 {
		return nil, errors.New("missing paren")
	}

	v = v[lIdx+1 : rIdx]

	pCount, inStr := 0, false
	var curr string
	for _, ch := range v {
		switch ch {
		case '(':
			pCount++
			curr += string(ch)
		case ')':
			pCount--
			curr += string(ch)
		case '"':
			if inStr {
				inStr = false
			} else {
				inStr = true
			}
			curr += string(ch)
		case ',':
			if !inStr && pCount == 0 {
				ret = append(ret, strings.TrimSpace(curr))
				curr = ""
			} else {
				curr += string(ch)
			}
		case '\n':
			if !inStr && pCount == 0 && curr != "" {
				ret = append(ret, strings.TrimSpace(curr))
				curr = ""
			} else {
				curr += string(ch)
			}
		default:
			curr += string(ch)
		}
	}
	if strings.TrimSpace(curr) != "" {
		ret = append(ret, strings.TrimSpace(curr))
	}

	return ret, nil
}
