package medium

/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(s string) []int {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	res, start, end := []int{}, 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}

// @lc code=end
