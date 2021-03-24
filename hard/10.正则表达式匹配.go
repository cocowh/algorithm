package hard

func IsMatch(s string, p string) bool {
	return isMatch(s, p)
}

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

//  * 每次匹配s[0]=p[0] || p[0]='.'
//  * 有*则可匹配当前字符0次，跳过该字符；匹配一次或多次则只需s后移一位s[1:]再次匹配
//  * 无*作普通字符匹配，s和p后移，s[1:]和p[1:]继续匹配
// @lc code=start
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

// @lc code=end

// 能力有限，后续再来理解其他解法。

// 优化解法：

//      [动态规划](https://leetcode-cn.com/problems/regular-expression-matching/solution/zheng-ze-biao-da-shi-pi-pei-by-leetcode/)

// 其他解法：

//      [有限状态机](https://leetcode-cn.com/problems/regular-expression-matching/solution/yi-bu-dao-wei-zhi-jie-an-zheng-ze-biao-da-shi-de-s/)
