package medium

func FindDuplicate(nums []int) int {
	return findDuplicate(nums)
}

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	// 快慢指针找交点（此时即链表的value）
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	// 再次从前遍历，找到value对应的key，即为重复值
	for slow = 0; slow != fast; {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// @lc code=end
