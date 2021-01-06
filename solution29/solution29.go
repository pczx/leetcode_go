package solution29

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, cols := len(matrix), len(matrix[0])
	ans := make([]int, 0, rows*cols)

	start := 0

	for start*2 < rows && start*2 < cols {
		printMatrixInCircle(&matrix, rows, cols, start, &ans)
		start++
	}

	return ans
}

func printMatrixInCircle(matrix *[][]int, rows, cols, start int, ans *[]int) {
	endX, endY := cols-1-start, rows-1-start

	for i := start; i <= endX; i++ {
		*ans = append(*ans, (*matrix)[start][i])
	}

	if start < endY {
		for i := start + 1; i <= endY; i++ {
			*ans = append(*ans, (*matrix)[i][endX])
		}
	}

	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			*ans = append(*ans, (*matrix)[endY][i])
		}
	}

	if start < endX && start < endY-1 {
		for i := endY - 1; i > start; i-- {
			*ans = append(*ans, (*matrix)[i][start])
		}
	}
}
