package medium

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	prevMax, currMax := 0, 0
	for _, v := range nums {
		if prevMax+v > currMax {
			currMax, prevMax = prevMax+v, currMax
		} else {
			prevMax = currMax
		}
	}
	return currMax
}

func robByDp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		curMax := dp[i-2] + nums[i]
		if curMax > dp[i-1] {
			dp[i] = curMax
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)-1]
}

// @lc code=end
