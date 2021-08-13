package medium

func Exist(board [][]byte, word string) bool {
	return exist(board, word)
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	for row := range board {
		for col := range board[row] {
			if dfsHelper(board, word, rows, cols, row, col, 0) {
				return true
			}
		}
	}
	return false
}

func dfsHelper(board [][]byte, word string, rows, cols, row, col, point int) bool {
	if row >= rows || row < 0 || col >= cols || col < 0 || point >= len(word) || board[row][col] != word[point] {
		return false
	}
	if point == len(word)-1 {
		return true
	}
	board[row][col] = ' '
	res := dfsHelper(board, word, rows, cols, row+1, col, point+1) ||
		dfsHelper(board, word, rows, cols, row, col+1, point+1) ||
		dfsHelper(board, word, rows, cols, row-1, col, point+1) ||
		dfsHelper(board, word, rows, cols, row, col-1, point+1)
	board[row][col] = word[point]
	return res
}
