package medium

import (
	"math"
	"strings"
)

func MyAtoi(str string) int {
	return myAtoi(str)
}

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {

	str = strings.Trim(str, " ")
	len := len(str)
	if len == 0 || str == "" {
		return 0
	}
	i := 0
	result := 0
	flag := true
	if str[i] == '+' {
		i++
	} else if str[i] == '-' {
		i++
		flag = false
	} else if str[i] < '0' || str[i] > '9' {
		return 0
	}
	for ; i < len; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			result = result*10 + int(str[i]) - int('0')
			if result > math.MaxInt32 || result < math.MinInt32 {
				if flag {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
		} else {
			break
		}
	}
	if flag {
		return result
	} else {
		return -result
	}
}

// @lc code=end

// * 过滤首尾空字符
// * 确定首字符是否是符号或者数字，标记符号
// * 遍历字符串累加值，判断是否超越边界
// * 根据符号标记返回值
