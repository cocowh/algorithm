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
		if length-i <= maxLen/2 { //剩余的可搜索字串长度小于已寻找的最大长度，无寻找必要
			break
		}
		b, e := i, i
		for e < length-1 && s[e+1] == s[e] { //多重复字符中心时，向右寻找中心边界
			e++
		}
		i = e + 1                                       //下次搜索中心（跳过连续重复）
		for e < length-1 && b > 0 && s[e+1] == s[b-1] { //双边拓展，e=b：单中心，e!=b：多字符中心
			e++
			b--
		}
		newLen := e + 1 - b
		if newLen > maxLen { //此次搜索回文子串是否为最大回文子串
			start = b
			maxLen = newLen
		}
	}
	return s[start : start+maxLen]
}

// @lc code=end

// 参考：[中心扩展算法](https://leetcode-cn.com/problems/longest-palindromic-substring/solution/)
