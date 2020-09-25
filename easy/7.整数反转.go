package easy

import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	ret := 0
	for x != 0 {
		rem := x % 10 //求余
		x /= 10
		if ret > math.MaxInt32/10 || (ret == math.MaxInt32/10 && rem > 7) {
			return 0
		}
		if ret < math.MinInt32/10 || (ret == math.MinInt32/10 && rem < -8) {
			return 0
		}
		ret = ret*10 + rem
	}
	return ret
}

// @lc code=end
