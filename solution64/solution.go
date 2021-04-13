package solution64

func sumNums(n int) int {
	if n == 0 {
		return 0
	}
	return sumNums(n - 1) + n
}
