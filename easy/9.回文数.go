package easy

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if 0 < x && x < 9 {
		return true
	}
	reverse, temp := 0, x
	for temp != 0 {
		rem := temp % 10
		temp /= 10
		reverse = reverse*10 + rem
	}
	return reverse^x == 0
}

// @lc code=end
