package main

import (
	"bytes"
	"strconv"
)

func compressString(S string) string {
	var T bytes.Buffer
	count := 0
	for i := 0; i < len(S); {
		j := i
		for j < len(S) && S[j] == S[i] {
			count++
			j++
		}
		T.WriteByte(S[i])
		T.WriteString(strconv.Itoa(count))
		i = j
		count = 0
	}
	if T.Len() >= len(S) {
		return S
	}
	return T.String()
}
