package easy

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	stack := []byte{}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < length; i++ {
		if m[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// @lc code=end
