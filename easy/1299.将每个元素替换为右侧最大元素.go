package easy

/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {
	last := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	if len(arr) <= 1 {
		return arr
	}
	for i := len(arr) - 2; i >= 0; i-- {
		temp := arr[i]
		if arr[i+1] > last {
			arr[i] = arr[i+1]
		} else {
			arr[i] = last
		}
		last = temp
	}
	return arr
}

// @lc code=end
