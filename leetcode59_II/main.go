package main

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx]
		if r, c := row+dir[0], col+dir[1]; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIdx = (dirIdx + 1) % 4
			dir = dirs[dirIdx]
		}
		row += dir[0]
		col += dir[1]
	}
	return matrix
}
