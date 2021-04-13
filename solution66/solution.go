package solution66

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return []int{}
	}
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * a[i-1]
	}
	temp := 1
	for j := n - 2; j >= 0; j-- {
		temp *= a[j + 1]
		ans[j] *= temp
	}
	return ans
}
