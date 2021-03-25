package easy

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	// index 统计有效长度，有效数组下标
	// temp 标记上一个值，用于与当前判重
	temp, index := nums[0], 1
	for i := 0; i < length; i++ {
		if nums[i] == temp { // 重复数字 遍历跳过
			continue
		} else {
			nums[index] = nums[i]
			temp = nums[i]
			index++
		}
	}
	return index
}

// @lc code=end

//原地删除，返回长度    => 原地进行数据替换，根据返回长度读取新数组

//快慢双指针，快指针遍历数组至末尾，读取元素，慢指针标记替换位置进行数据替换
