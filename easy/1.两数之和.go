package easy

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))

	for k, v := range nums {
		if i, ok := indexMap[target-v]; ok {
			return []int{k, i}
		}
		indexMap[v] = k
	}
	return nil
}

// @lc code=end
