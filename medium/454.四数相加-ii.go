package medium

/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	countAB := map[int]int{}
	for _, v := range nums1 {
		for _, w := range nums2 {
			countAB[v+w]++
		}
	}
	res := 0
	for _, v := range nums3 {
		for _, w := range nums4 {
			res += countAB[-v-w]
		}
	}
	return res
}

// @lc code=end
