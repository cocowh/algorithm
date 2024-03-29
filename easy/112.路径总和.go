package easy

func HasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSum(root, targetSum)
}

/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// @lc code=end
