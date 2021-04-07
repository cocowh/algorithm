package medium

func SumNumbers(root *TreeNode) int {
	return sumNumbers(root)
}

/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	pathSet, res := []int{}, 0
	sumNumbersHelper(root, &pathSet, 0)
	for _, v := range pathSet {
		res += v
	}
	return res
}
func sumNumbersHelper(node *TreeNode, pathSet *[]int, path int) {
	if node == nil {
		return
	}
	path = path*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*pathSet = append(*pathSet, path)
	}
	sumNumbersHelper(node.Left, pathSet, path)
	sumNumbersHelper(node.Right, pathSet, path)
}

// @lc code=end
