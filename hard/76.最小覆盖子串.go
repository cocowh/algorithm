package hard

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	// 统计t字符数
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	// 检查当前窗口是否满足
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		// 统计当前窗口内字符
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		// 检查窗口是否可以收缩
		for check() && l <= r {
			// 更小的窗口
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

// @lc code=end
