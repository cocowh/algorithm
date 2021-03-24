package medium

func Convert(s string, numRows int) string {
	return convert(s, numRows)
}

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	gap := 2*numRows - 2
	var ret string
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < length; j += gap {
			ret += string(s[i+j])
			if i != 0 && i != numRows-1 && j+gap-i < length {
				ret += string(s[j+gap-i])
			}
		}
	}
	return ret
}

// @lc code=end

// 想到的遍历字符串确定位置，然后拼接，或者按确定位置在字符串遍历然后拼接。
