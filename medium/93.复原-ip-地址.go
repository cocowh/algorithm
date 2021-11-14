package medium

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	res := []string{}
	stack := make([]int, 4)
	var findAddress func(seqId, start int)
	findAddress = func(seqId, start int) {
		if seqId == 4 && start == len(s) {
			ipAddr := ""
			for i := 0; i < 4; i++ {
				ipAddr += strconv.Itoa(stack[i])
				if i != 3 {
					ipAddr += "."
				}
			}
			res = append(res, ipAddr)
			return
		}
		if seqId >= 4 || start >= len(s) {
			return
		}
		if s[start] == '0' {
			stack[seqId] = 0
			findAddress(seqId+1, start+1)
		}
		addr := 0
		for i := start; i < len(s); i++ {
			addr = addr*10 + int(s[i]-'0')
			if addr > 0 && addr <= 0xFF {
				stack[seqId] = addr
				findAddress(seqId+1, i+1)
			} else {
				break
			}
		}
	}
	findAddress(0, 0)
	return res
}

// @lc code=end
