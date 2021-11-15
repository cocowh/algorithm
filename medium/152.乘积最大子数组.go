package medium

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	maxNum, minNum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxNum, minNum = minNum, maxNum
		}
		if nums[i] > maxNum*nums[i] {
			maxNum = nums[i]
		} else {
			maxNum = maxNum * nums[i]
		}
		if nums[i] < minNum*nums[i] {
			minNum = nums[i]
		} else {
			minNum = minNum * nums[i]
		}
		if res < maxNum {
			res = maxNum
		}
	}
	return res
}

// @lc code=end
