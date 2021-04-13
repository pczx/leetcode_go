package main

import (
	"fmt"
	"strconv"
)

func printNumbers(n int) []int {
	ans := []int{}
	if n <= 0 {
		return ans
	}
	str := make([]byte, n)
	var dfs func(int)
	dfs = func(index int) {
		if index == n {
			v, _ := strconv.Atoi(string(str[:n]))
			if v != 0 {
				ans = append(ans, v)
			}
			return
		}
		for i := 0; i < 10; i++ {
			str[index] = byte(i + '0')
			dfs(index + 1)
		}
	}
	dfs(0)
	return ans
}

func main () {
	fmt.Println(printNumbers(4))
}
