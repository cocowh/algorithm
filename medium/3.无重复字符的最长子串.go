package medium

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	length, maxLen := len(s), 0
	pMap := make(map[byte]int, length)
	for i, j := 0, 0; j < length; j++ {
		if _, ok := pMap[s[j]]; ok {
			i = max(i, pMap[s[j]])
		}
		maxLen = max(maxLen, (j - i + 1))
		pMap[s[j]] = j + 1
	}
	return maxLen
}

// @lc code=end
