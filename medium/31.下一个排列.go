package medium

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// 将一个左边的「较小数」与一个右边的「较大数」交换，以能够让当前排列变大，从而得到下一个排列。
// 让这个「较小数」尽量靠右，而「较大数」尽可能小。交换完成后，「较大数」右边的数需要按照升序重新排列。保证新排列大于原来排列的情况下，使变大的幅度尽可能小

// @lc code=start
func nextPermutation(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	// 从后往前找第一个序列a[left] < a[left+1], [left+1, n)递减
	left := length - 1
	for ; left > 0; left-- {
		if nums[left] <= nums[left-1] {
			continue
		} else {
			break
		}
	}
	// 特殊情况，已是最大数，通过翻转调整为最小数
	if left == 1 && nums[left] <= nums[left-1] || left == 0 {
		for i := 0; i < length/2; i++ {
			temp := nums[i]
			nums[i] = nums[length-1-i]
			nums[length-1-i] = temp
		}
	} else {
		left--
		// 从后往前在[left+1,n)找第一个a[right] > a[left]
		// 交换a[right],a[left]
		for right := length - 1; right > left; right-- {
			if nums[right] > nums[left] {
				temp := nums[right]
				nums[right] = nums[left]
				nums[left] = temp
				break
			}
		}
		// 对右边按升序排序
		i, j := left+1, length-1
		for i < j {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
			j--
		}
	}
}

// @lc code=end
