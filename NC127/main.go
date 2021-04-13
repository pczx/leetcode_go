package main

/**
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
func LCS(str1 string, str2 string) string {
	// write code here
	m, n := len(str1), len(str2)
	maxlen := 0
	index := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxlen {
					maxlen = dp[i][j]
					index = i
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	if maxlen == 0 {
		return "-1"
	}
	return str1[index-maxlen+1 : index+1]
}
