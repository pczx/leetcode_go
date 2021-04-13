package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		str := in.Text()
		fmt.Println(sort_string(str))
	}
}

func sort_string(str string) string {
	bytes := []byte{}
	for j := 0; j < 26; j++ {
		for i := 0; i < len(str); i++ {
			if int(str[i]-'a') == j || int(str[i]-'A') == j {
				bytes = append(bytes, str[i])
			}
		}
	}
	b := []byte(str)
	for i, k := 0, 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			b[i] = bytes[k]
			k++
		}
	}
	return string(b)
}
