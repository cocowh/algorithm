package medium

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// 二分

// @lc code=start
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 旋转点
	pointIndex := 0
	if nums[0] > nums[length-1] {
		for left, right := 0, length-1; left <= right; {
			mid := (left + right) / 2
			if nums[mid] > nums[mid+1] {
				pointIndex = mid + 1
				break
			} else if nums[mid] < nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	// 按旋转点划分进行二分查找
	if nums[pointIndex] == target {
		return pointIndex
	} else if pointIndex == 0 { // 未旋转
		return binarySearch(nums, target, 0, length-1)
	} else if nums[0] > target { // 在右边集合二分查找
		return binarySearch(nums, target, pointIndex+1, length-1)
	} else { // 在左边集合二分查找
		return binarySearch(nums, target, 0, pointIndex-1)
	}
}

// 二分查找
func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end
