package hard

func Trap(height []int) int {
	return trap(height)
}

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	length := len(height)
	if length <= 2 {
		return 0
	}
	left, right, leftMax, rightMax, area := 0, length-1, 0, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				area += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				area += rightMax - height[right]
			}
			right--
		}
	}
	return area
}

// @lc code=end

// copy https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0042.Trapping-Rain-Water/
