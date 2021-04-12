package hard

func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start

func LongestValidParenthesesByDp(s string) int {
	res, dp := 0, make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i > 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func LongestValidParenthesesByStack(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if maxAns < i-stack[len(stack)-1] {
					maxAns = i - stack[len(stack)-1]
				}
			}
		}
	}
	return maxAns
}

func LongestValidParenthesesBetter(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if maxLength < 2*right {
				maxLength = 2 * right
			}
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if maxLength < 2*left {
				maxLength = 2 * left
			}
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func longestValidParentheses(s string) int {
	// return LongestValidParenthesesByDp(s)
	return LongestValidParenthesesBetter(s)
}

// @lc code=end
