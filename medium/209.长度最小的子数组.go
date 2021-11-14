package medium

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	res := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < length {
		sum += nums[end]
		for sum >= target {
			if res > end-start+1 {
				res = end - start + 1
			}
			sum -= nums[start]
			start++
			fmt.Println(start, end, res)
		}
		end++
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

// @lc code=end
