package solution20

import (
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	index := 0

	scanUnsignedInteger := func(s string) bool {
		before := index
		for index != len(s) && s[index] >= '0' && s[index] <= '9' {
			index++
		}
		return index > before
	}

	scanInteger := func(s string) bool {
		if index != len(s) && (s[index] == '-' || s[index] == '+') {
			index++
		}
		return scanUnsignedInteger(s)
	}

	numeric := scanInteger(s)

	if index != len(s) && s[index] == '.' {
		index++
		numeric = scanUnsignedInteger(s) || numeric
	}

	if index != len(s) && (s[index] == 'e' || s[index] == 'E') {
		index++
		numeric = numeric && scanInteger(s)
	}

	return numeric && index == len(s)
}
