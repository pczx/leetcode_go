package solution12

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}

	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
		return false
	}

	if word[k] == board[i][j] {
		temp := board[i][j]
		board[i][j] = ' '
		ds := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
		flag := false
		for t := 0; t < 4; t++ {
			flag = flag || dfs(board, i+ds[t][0], j+ds[t][1], k+1, word)
		}
		if flag {
			return true
		}
		board[i][j] = temp
	}
	return false
}
