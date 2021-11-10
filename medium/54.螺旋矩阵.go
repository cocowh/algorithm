package medium

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	minRow, maxRow, minCol, maxCol, res := 0, len(matrix)-1, 0, len(matrix[0])-1, []int{}
	for minRow <= maxRow && minCol <= maxCol {
		//上层
		for c := minCol; c <= maxCol; c++ {
			res = append(res, matrix[minRow][c])
		}
		//右层
		for r := minRow + 1; r <= maxRow; r++ {
			res = append(res, matrix[r][maxCol])
		}
		if minRow < maxRow && minCol < maxCol {
			//下层
			for c := maxCol - 1; c > minCol; c-- {
				res = append(res, matrix[maxRow][c])
			}
			//左层
			for r := maxRow; r > minRow; r-- {
				res = append(res, matrix[r][minCol])
			}
		}
		//收缩
		minCol++
		minRow++
		maxCol--
		maxRow--
	}
	return res
}

// @lc code=end
