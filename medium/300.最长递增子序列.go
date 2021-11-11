package medium

import "sort"

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
// 解法一 O(n log n) DP
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, v := range nums {
		i := sort.SearchInts(dp, v)
		if i == len(dp) {
			dp = append(dp, v)
		} else {
			dp[i] = v
		}
	}
	return len(dp)
}

// 解法二 O(n^2) DP
func lengthOfLIS2(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}
	return res
}

// @lc code=end
