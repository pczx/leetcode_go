package solution46

import "strconv"

func translateNum(num int) int {
	if num < 0 {
		return 0
	}
	return translateNumCore(strconv.Itoa(num))
}

func translateNumCore(src string) int {
	counts := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		if i == 0 {
			counts[i] = 1
			continue
		}
		counts[i] = counts[i-1]
		converted := src[i-1 : i+1]
		if converted >= "10" && converted <= "25" {
			if i > 1 {
				counts[i] += counts[i-2]
			} else {
				counts[i]++
			}
		}
	}
	return counts[len(src)-1]
}
