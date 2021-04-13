package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve(s string, t string) string {
	// write code here
	var ans []rune
	var carry int
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 || carry != 0 {
		var x, y int
		if i < 0 {
			x = 0
		} else {
			x = int(s[i] - '0')
			i--
		}
		if j < 0 {
			y = 0
		} else {
			y = int(t[j] - '0')
			j--
		}
		sum := x + y + carry
		ans = append(ans, rune(sum%10+'0'))
		carry = sum / 10
	}
	for i, j = 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
