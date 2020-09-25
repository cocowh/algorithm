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
// 思考：根据题意，整数值范围在[-214748368,2147483647]，
// 大于math.MaxInt32或者小于math.MinInt32为溢出，返回0。
// 对原值依次除10取余，商为下次迭代初始值，结果值依次乘10加上所得余数，直至原值除尽为0。
// 在期间判断值是否超出范围。
