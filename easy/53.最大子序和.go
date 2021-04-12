package easy

func MaxSubArray(nums []int) int {
	return maxSubArray(nums)
}

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	res, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}

// @lc code=end

// 记连续和sum，最大连续和maxSum，每次求得连续和与maxSum比较取大值赋给maxSum。
// 若sum > 0，表示连续和目前还不会对后续和产生负增益，还可以继续保留为后续做"贡献"
// 若sum < 0，表示当前连续和会减少后续和的值，应当舍弃，将当前连续和设为当前值。
