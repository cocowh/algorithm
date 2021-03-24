package medium

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	length := len(s)

	if length < 2 {
		return s
	}

	start, maxLen := 0, 1
	for i := 0; i < length; {
		if length-i <= maxLen/2 {
			break
		}
		b, e := i, i
		for e < length-1 && s[e+1] == s[e] {
			e++
		}
		i = e + 1
		for e < length-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}
		newLen := e + 1 - b
		if newLen > maxLen {
			start = b
			maxLen = newLen
		}
	}
	return s[start : start+maxLen]
}

// @lc code=end

// 参考：[中心扩展算法](https://leetcode-cn.com/problems/longest-palindromic-substring/solution/)
