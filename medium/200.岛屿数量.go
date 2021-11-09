package medium

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func dfsNumIslands(grid [][]byte, r, c int) {
	nr, nc := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfsNumIslands(grid, r-1, c)
	dfsNumIslands(grid, r+1, c)
	dfsNumIslands(grid, r, c-1)
	dfsNumIslands(grid, r, c+1)
}
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	nr, nc, res := len(grid), len(grid[0]), 0
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				res++
				dfsNumIslands(grid, i, j)
			}
		}
	}
	return res
}

// @lc code=end
