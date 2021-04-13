package solution60

import "math"

func dicesProbability(n int) []float64 {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 6*n+1)
	}
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= 6*n; j++ {
			for k := 1; k <= 6; k++ {
				if j-k < 1 {
					break
				}
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	total := math.Pow(6, float64(n))
	ans := make([]float64, 5*n+1)
	for i := n; i <= 6 * n; i++ {
		ans[i - n] = float64(dp[n][i]) / total
	}
	return ans
}
