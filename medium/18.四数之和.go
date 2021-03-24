package medium

import "sort"

func FourSum(nums []int, target int) [][]int {
	return fourSum(nums, target)
}

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(nums) //排序
	length := len(nums)
	for i := 0; i < length-3; i++ { // 定第一位
		if i > 0 && nums[i] == nums[i-1] { //非首次，第一位重复，跳过
			continue
		}
		for j := i + 1; j < length-2; j++ { // 定第二位
			if j > i+1 && nums[j] == nums[j-1] { // 非首次，第二位重复，跳过
				continue
			}
			l, r := j+1, length-1 // 快慢指针确定后两位
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r] // 确定偏移
				if sum < target || (l != j+1 && nums[l] == nums[l-1]) {
					l++
				} else if sum > target || (r != length-1 && nums[r] == nums[r+1]) {
					r--
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				}
			}
		}
	}
	return ret
}

// @lc code=end
