package medium

func NumDecodings(s string) int {
	return numDecodings(s)
}

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	length := len(s)
	// 首字符为0或者最后两个字符大于20
	if length == 0 || s[0] == '0' || (length > 1 && s[length-1] == '0' && s[length-2] > '2') {
		return 0
	}
	pre, res := 1,1
	for i := length-2; i > 0; i-- {
		if s[i] == '0' { 
			if s[i-1] == '0' || s[i-1] > '2' { // 中间连续两个0或者大于20，无法编码
				return 0
			}
			res += 
		} else if s[i] == '0' && s[]
		if s[i] == '0' && s[i-1] > '2' {
			return 
		} else if s[i-1] <  {
			
		}
	}
}

// @lc code=end
