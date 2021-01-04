package solution17

import (
	"strconv"
)

func printNumbers(n int) []int {
	res := []int{}
	if n <= 0 {
		return res
	}

	str := make([]byte, n)
	for i := 0; i < 10; i++ {
		str[0] = byte(i + int('0'))
		dfs(str, n, 0, &res)
	}
	return res
}

func dfs(str []byte, n, index int, res *[]int) {
	if index == n-1 {
		v, _ := strconv.Atoi(string(str))
		if v != 0 {
			*res = append(*res, v)
		}
		return
	}

	for i := 0; i < 10; i++ {
		str[index+1] = byte(i + int('0'))
		dfs(str, n, index+1, res)
	}
}
