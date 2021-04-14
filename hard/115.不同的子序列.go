package hard

func NumDistinct(s string, t string) int {
	return numDistinct(s, t)
}

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	ls, lt := len(s), len(t)
	if ls < lt {
		return 0
	}
	dp := make([][]int, ls+1)
	for i := range dp {
		dp[i] = make([]int, lt+1)
		dp[i][lt] = 1
	}
	for i := ls - 1; i >= 0; i-- {
		for j := lt - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

// @lc code=end
