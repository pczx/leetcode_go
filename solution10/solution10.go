package solution10

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return a
}
