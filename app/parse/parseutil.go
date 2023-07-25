package parse

import "strings"

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
