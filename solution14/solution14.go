package solution14

func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	products := make([]int, n+1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3
	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}
		}
		products[i] = max
	}
	return products[n]
}
