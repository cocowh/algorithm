package easy

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		m[v-'a']--
		if m[v-'a'] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end
