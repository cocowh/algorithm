/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	i, j, maxA := 0, len(height)-1, 0
	for i != j {
		if height[i] < height[j] {
			maxA = max(maxA, height[i]*(j-i))
			i++
		} else {
			maxA = max(maxA, height[j]*(j-i))
			j--
		}
	}
	return maxA
}

// @lc code=end

