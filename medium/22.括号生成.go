package medium

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// 对括号左边已'('开头，左括号的数量必定大于右括号数量时才可以加入一个右括号
// 记左右括号的数量，当左括号数量小于n时可以添加左括号
// 当右括号数量小于左括号时可以添加一个右括号

// @lc code=start
func generateParenthesis(n int) []string {
	ret := []string{}
	getParenthesis(&ret, "", 0, 0, n)
	return ret
}

func getParenthesis(ret *[]string, cur string, left, right, max int) {
	if len(cur) == max*2 {
		*ret = append(*ret, cur)
		return
	}
	if left < max {
		getParenthesis(ret, cur+"(", left+1, right, max)
	}
	if right < left {
		getParenthesis(ret, cur+")", left, right+1, max)
	}
}

// @lc code=end
