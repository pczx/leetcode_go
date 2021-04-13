package solution58

func reverseWords(s string) string {
	var ans string
	for i, j := 0, len(s)-1; j >= 0; j-- {
		if s[j] != ' ' {
			i = j
			for i >= 0 && s[i] != ' ' {
				i--
			}
			ans += s[i + 1 : j] + " "
			j = i
		}
	}
	if len(s) > 0 {
		ans = ans[:len(s) - 1]
	}
	return ans
}
