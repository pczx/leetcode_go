package solution47

import (
	"math"
)

func maxValue(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch {
			case i == 0 && j == 0:
				continue
			case i == 0:
				grid[i][j] += grid[i][j-1]
			case j == 0:
				grid[i][j] += grid[i-1][j]
			default:
				grid[i][j] += int(math.Max(float64(grid[i-1][j]), float64(grid[i][j-1])))
			}
		}
	}

	return grid[rows-1][cols-1]
}
