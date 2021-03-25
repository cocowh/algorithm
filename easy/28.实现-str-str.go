package easy

func StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	lengthH, lengthN := len(haystack), len(needle)
	if lengthN == 0 { //needle为空返回0，与c和java保持一致
		return 0
	}
	if lengthH == 0 || lengthH < lengthN {
		return -1
	}
	for i := 0; i < lengthH; i++ {
		if i+lengthN <= lengthH && haystack[i:i+lengthN] == needle {
			return i
		} else if i+lengthN > lengthH {
			return -1
		}
	}
	return -1
}

// @lc code=end
