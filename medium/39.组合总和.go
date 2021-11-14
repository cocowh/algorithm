package medium

import "sort"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	if len(candidates) == 0 {
		return ret
	}
	sort.Ints(candidates)
	var findCombinationSum func(stack []int, target, start int)
	findCombinationSum = func(stack []int, target, start int) {
		if target == 0 {
			temp := make([]int, len(stack))
			copy(temp, stack)
			ret = append(ret, temp)
			return
		}
		for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
			stack = append(stack, candidates[i])
			findCombinationSum(stack, target-candidates[i], i)
			stack = stack[:len(stack)-1]
		}
	}
	findCombinationSum([]int{}, target, 0)
	return ret
}

// @lc code=end
