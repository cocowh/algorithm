package easy

func Exchange(nums []int) []int {
	return exchange(nums)
}

// 双指针首找偶数，尾找奇数，然后交换
func exchange(nums []int) []int {
	length := len(nums)
	for i, j := 0, length-1; i < j; {
		for i < j && (nums[i]&1) == 1 { // 奇数
			i++
		}
		for i < j && (nums[j]&1) == 0 { // 偶数
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
