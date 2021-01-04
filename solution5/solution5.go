package main

import "strings"

func replaceSpace(s string) string {
	var sb strings.Builder
	for _, v := range s {
		if v == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(v)
		}
	}
	return sb.String()
}
