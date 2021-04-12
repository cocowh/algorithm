package easy

func ClimbStairs(n int) int {
	return climbStairs(n)
}

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	pre, next := 1, 2
	if n == 1 {
		return pre
	}
	if n == 2 {
		return next
	}
	for i := 2; i < n; i++ {
		next, pre = pre+next, next
	}
	return next
}

// @lc code=end
