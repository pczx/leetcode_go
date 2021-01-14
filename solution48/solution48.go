package solution48

func lengthOfLongestSubstring(s string) int {
	res, tmp, pos := 0, 0, make(map[rune]int)
	for j, c := range s {
		i := -1
		if v, ok := pos[c]; !ok {
			i = v
		}
		pos[c] = j
		if tmp < j-i {
			tmp++
		} else {
			tmp = j - i
		}
		res = max(res, tmp)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
