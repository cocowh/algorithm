package easy

func RemoveElement(nums []int, val int) int {
	return removeElement(nums, val)
}

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	length, index := len(nums), 0
	for i := 0; i < length; i++ {
		if nums[i] == val {
			continue
		} else {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

// @lc code=end
