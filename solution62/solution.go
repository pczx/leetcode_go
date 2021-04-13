package solution62

func lastRemaining(n int, m int) int {
	return f(n, m)
}

func f(n, m int) int {
	if n == 1 {
		return 0
	}
	x := f(n-1, m)
	return (m + x) % n
}
