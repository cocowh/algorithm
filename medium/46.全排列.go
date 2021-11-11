package medium

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func backtrack(length, first int, ret *[][]int, nums []int) {
	if first == length {
		tmp := make([]int, length)
		copy(tmp, nums)
		*ret = append(*ret, tmp)
	}
	for i := first; i < length; i++ {
		// 动态维护数组
		nums[first], nums[i] = nums[i], nums[first]
		// 继续递归填下一个数
		backtrack(length, first+1, ret, nums)
		// 撤销操作
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func permute(nums []int) [][]int {
	ret := [][]int{}
	backtrack(len(nums), 0, &ret, nums)
	return ret
}

// @lc code=end
