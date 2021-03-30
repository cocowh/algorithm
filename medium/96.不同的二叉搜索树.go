package medium

func NumTrees(n int) int {
	return numTrees(n)
}

// 公式
func NumTreesByFormula(n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret = ret * 2 * (2*i + 1) / (i + 2)
	}
	return ret
}

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

// @lc code=end

// 官方题解 https://leetcode-cn.com/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/
