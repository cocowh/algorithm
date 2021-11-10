package easy

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}
	res := make([]string, len(a)+1)
	// 从后往前累加
	i, j, k, c := len(a)-1, len(b)-1, len(a), 0
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		// 计算累加值
		res[k] = strconv.Itoa((ai + bj + c) % 2)
		// 计算进位
		c = (ai + bj + c) / 2
		i--
		j--
		k--
	}
	// 字符串a更长，i > 0则未处理完，继续计算进位
	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		res[k] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) / 2
		i--
		k--
	}
	// 判断进位是否处理位，确定首位
	if c > 0 {
		res[k] = strconv.Itoa(c)
	}

	return strings.Join(res, "")
}

// @lc code=end
