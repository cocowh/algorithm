package medium

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ret := []string{""}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	for _, v := range digits {
		copy := []string{}
		retLen := len(ret)
		for j := 0; j < retLen; j++ {
			arrLen := len(m[byte(v)])
			for i := 0; i < arrLen; i++ {
				copy = append(copy, ret[j]+m[byte(v)][i])
			}
		}
		ret = copy
	}
	return ret
}

// @lc code=end
