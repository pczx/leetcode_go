package solution10

func numWays(n int) int {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return a
}
