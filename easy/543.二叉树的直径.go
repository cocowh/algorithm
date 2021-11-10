package easy

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var helper func(node *TreeNode, res *int) int
	helper = func(node *TreeNode, res *int) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left, res)
		right := helper(node.Right, res)
		*res = max543(*res, left+right)
		return max543(left, right) + 1
	}
	helper(root, &res)
	return res
}

func max543(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
