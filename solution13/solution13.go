package solution13

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, visited)
}

func dfs(i, j, m, n, k int, visited [][]bool) int {
	count := 0
	if i >= 0 && i < m && j >= 0 && j < n && getDigitSum(i)+getDigitSum(j) <= k && !visited[i][j] {
		visited[i][j] = true
		count = 1 + dfs(i, j-1, m, n, k, visited) +
			dfs(i, j+1, m, n, k, visited) +
			dfs(i-1, j, m, n, k, visited) +
			dfs(i+1, j, m, n, k, visited)
	}
	return count
}

func getDigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
