package medium

func ValidateStackSequences(pushed []int, popped []int) bool {
	return validateStackSequences(pushed, popped)
}

/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack, j, n := []int{}, 0, len(pushed)
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) != 0 && j < n && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1]
			j++
		}
	}
	return j == n
}

// @lc code=end
