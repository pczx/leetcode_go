package solution50

func firstUniqChar(s string) byte {
	count := make([]int, 128)
	for _, c := range s {
		count[c-'a']++
	}

	for _, c := range s {
		if count[c-'a'] == 1 {
			return byte(c)
		}
	}
	return ' '
}
