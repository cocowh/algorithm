package easy

func MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func maxDepthByRecursion(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	depth++
	depthLeft := maxDepthByRecursion(node.Left, depth)
	depthRight := maxDepthByRecursion(node.Right, depth)
	if depthLeft > depthRight {
		return depthLeft
	}
	return depthRight
}

func maxDepth(root *TreeNode) int {
	return maxDepthByRecursion(root, 0)
}

// @lc code=end
