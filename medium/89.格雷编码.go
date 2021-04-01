package medium

import "fmt"

func GrayCode(n int) []int {
	return grayCode(n)
}

/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start

func grayCode(n int) []int {
	var l uint = 1 << uint(n)
	out := make([]int, l)
	for i := uint(0); i < l; i++ {
		fmt.Println(i)
		out[i] = int((i >> 1) ^ i)
	}
	return out
}

// @lc code=end
