package extract

import "strings"

func ExtractSections(content string) []string {
	lines := strings.Split(content, "\n")
	var ret []string
	var curr string
	for idx := range lines {
		line := lines[idx]
		clean := cleanLine(line)
		if len(clean) == 0 || strings.HasPrefix(clean, "package ") || strings.HasPrefix(clean, "import ") || strings.HasPrefix(clean, "object ") {
			continue
		}
		if strings.HasPrefix(line, "  ") {
			if strings.HasPrefix(line, "    ") {
				curr += "\n" + line
				continue
			}
			if clean == ")" || clean == "))" || clean == ")," || clean == "))," {
				curr += "\n" + line
				ret = append(ret, curr)
				curr = ""
				continue
			}
			if strings.Contains(line, "(") {
				if strings.Contains(line, ")") {
					ret = append(ret, line)
				} else {
					curr = line
				}
				continue
			}
			ret = append(ret, line)
		}
	}
	if curr != "" {
		ret = append(ret, curr)
	}
	return ret
}
