package medium

func NumDecodings(s string) int {
	return numDecodings(s)
}

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	length := len(s)
	// 首字符为0或者最后两个字符大于20
	if length == 0 || s[0] == '0' || (length > 1 && s[length-1] == '0' && s[length-2] > '2') {
		return 0
	}
	pre, res := 1, 1
	for i := 1; i < length; i++ {
		temp := res
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' { // 为0时只有s[i-1]为1或者2时才能编码
				res = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') { // 只有连续两个字符在01-26之间才可编码
			res += pre
		}
		pre = temp
	}
	return res
}

// @lc code=end
