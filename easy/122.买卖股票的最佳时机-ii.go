package easy

func MaxProfitII(prices []int) int {
	return maxProfitII(prices)
}

/**
 * dp[i] = dp[]
 */

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start

func maxProfitII(prices []int) int {
	return MaxProfitIIByDP(prices)
}
func MaxProfitIIByGreedy(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

func MaxProfitIIByDP(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		if dp1+prices[i] > dp0 {
			dp0 = dp1 + prices[i]
		}
		if dp0-prices[i] > dp1 {
			dp1 = dp0 - prices[i]
		}
	}
	return dp0
}

// @lc code=end
