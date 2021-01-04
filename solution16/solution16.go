package solution16

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1.0 / x
	}
	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n /= 2
	}
	return res
}
