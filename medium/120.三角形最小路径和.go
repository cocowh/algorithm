package medium

import "math"

func MinimumTotal(triangle [][]int) int {
	return minimumTotal(triangle)
}

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	return MininumTotal2(triangle)
	// return MinimumTotalByArray(triangle)
}

// 自底向上
func MininumTotal2(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}

// 自顶向下
func MinimumTotalByArray(triangle [][]int) int {
	width := len(triangle[len(triangle)-1])
	dp := make([]int, width)
	minNum, index := math.MaxInt64, 0
	for ; index < len(triangle[0]); index++ {
		dp[index] = triangle[0][index]
	}
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 { // 最左边
				dp[j] += triangle[i][0]
			} else if j == len(triangle[i])-1 { // 最右边
				dp[j] = dp[j-1] + triangle[i][j]
			} else { // 选取最短路径
				dp[j] = min(dp[j-1]+triangle[i][j], dp[j]+triangle[i][j])
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		if dp[i] < minNum {
			minNum = dp[i]
		}
	}
	return minNum
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
