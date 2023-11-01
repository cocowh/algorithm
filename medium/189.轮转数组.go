package medium

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	reverse := func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// @lc code=end
