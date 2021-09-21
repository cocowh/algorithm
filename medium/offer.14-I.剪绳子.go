package medium

func CuttingRope(n int) int {
	return cuttingRope(n)
}

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		maxV := 0
		for j := 1; j <= i/2; j++ {
			maxV = max(maxV, dp[j]*dp[i-j])
		}
		dp[i] = maxV
	}
	return dp[n]
}
