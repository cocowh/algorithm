package medium

func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 || len(matrix[0]) == 0 {
		return false
	}
	columns := len(matrix[0])
	row, col := 0, columns-1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
