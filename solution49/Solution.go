package solution49

func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	uglyNumbers := make([]int, n)
	uglyNumbers[0] = 1
	nextUglyNumber := 1
	multiply2, multiply3, multiply5 := 0, 0, 0
	for nextUglyNumber < n {
		uglyNumbers[nextUglyNumber] = min(min(uglyNumbers[multiply2]*2, uglyNumbers[multiply3]*3), uglyNumbers[multiply5]*5)
		for uglyNumbers[multiply2]*2 <= uglyNumbers[nextUglyNumber] {
			multiply2++
		}
		for uglyNumbers[multiply3]*3 <= uglyNumbers[nextUglyNumber] {
			multiply3++
		}
		for uglyNumbers[multiply5]*5 <= uglyNumbers[nextUglyNumber] {
			multiply5++
		}
		nextUglyNumber++
	}
	return uglyNumbers[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
