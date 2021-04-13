package solution13

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, vis)
}

func dfs(i, j, m, n, k int, vis [][]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n || getDigit(i)+getDigit(j) > k || vis[i][j] {
		return 0
	}
	count := 1
	vis[i][j] = true
	ds := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for t := 0; t < 4; t++ {
		count += dfs(i+ds[t][0], j+ds[t][1], m, n, k, vis)
	}
	return count
}

func getDigit(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}