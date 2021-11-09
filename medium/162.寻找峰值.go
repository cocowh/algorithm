package medium

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		mid := (head + tail) / 2
		if nums[mid] > nums[mid+1] {
			tail = mid
		} else {
			head = mid + 1
		}
	}
	return head
}

// @lc code=end
