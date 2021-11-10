package medium

import "fmt"

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	fmt.Println(dp)
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min322(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min322(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
