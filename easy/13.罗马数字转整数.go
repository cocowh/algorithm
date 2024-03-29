package easy

func RomanToInt(s string) int {
	return romanToInt(s)
}

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	intNum := 0
	intEnum := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := intEnum[s[i]]
		sign := 1
		if temp < last {
			sign = -1
		}
		intNum += sign * temp
		last = temp
	}
	return intNum
}

// @lc code=end
