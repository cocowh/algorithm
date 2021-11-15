package hard

import "fmt"

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	length := len(ratings)
	// inc 递增序列长度
	// dec 递减序列长度
	// pre 上一个分配数
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < length; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
		fmt.Println(inc, dec, pre)
	}
	return ans
}

// @lc code=end
