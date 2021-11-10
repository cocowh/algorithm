package medium

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[1] {
			if intervals[i][1] >= temp[1] {
				temp[1] = intervals[i][1]
			}
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)
	return res
}

// @lc code=end
