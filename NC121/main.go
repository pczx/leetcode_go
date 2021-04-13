package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return string字符串一维数组
 */

import "sort"

func Permutation(str string) []string {
	// write code here
	if len(str) == 0 {
		return []string{""}
	}
	ans := []string{}
	strs := []byte(str)
	dfs(0, strs, &ans)
	sort.Strings(ans)
	return ans
}

func dfs(start int, str []byte, ans *[]string) {
	if start == len(str) {
		*ans = append(*ans, string(str))
		return
	}
	set := make(map[byte]struct{})
	for i := start; i < len(str); i++ {
		if _, ok := set[str[i]]; ok {
			continue
		}
		set[str[i]] = struct{}{}
		str[start], str[i] = str[i], str[start]
		dfs(start+1, str, ans)
		str[start], str[i] = str[i], str[start]
	}
}
