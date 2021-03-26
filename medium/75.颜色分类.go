package medium

func SortColors(nums []int) {
	sortColors(nums)
}

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	r, w, b := 0, 0, 0
	for _, v := range nums {
		switch v {
		case 0:
			nums[b] = 2
			b++
			nums[w] = 1
			w++
			nums[r] = 0
			r++
		case 1:
			nums[b] = 2
			b++
			nums[w] = 1
			w++
		case 2:
			b++
		}

	}
}

// @lc code=end

// 三路快排 https://segmentfault.com/a/1190000021726667
