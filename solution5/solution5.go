package main

import "fmt"

func replaceSpace(s string) string {
	spaces := 0
	for _, c := range s {
		if c == ' ' {
			spaces++
		}
	}

	ans := make([]rune, len(s) + 2 * spaces)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans[i] = '%'
			ans[i + 1] = '2'
			ans[i + 2] = '0'
			i = i + 3
		} else {
			ans[i] = rune(s[i])
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}
