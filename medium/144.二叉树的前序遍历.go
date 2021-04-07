package medium

func PreorderTraversal(root *TreeNode) []int {
	return preorderTraversal(root)
}

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	return PreorderTraversalByStack(root)
	// return PreorderTraversalByRecursion(root)
}

func PreorderTraversalByRecursion(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	res := append([]int{}, node.Val)
	res = append(res, PreorderTraversalByRecursion(node.Left)...)
	res = append(res, PreorderTraversalByRecursion(node.Right)...)
	return res
}

func PreorderTraversalByStack(root *TreeNode) []int {
	res, stack, point := []int{}, []*TreeNode{}, root
	for point != nil || len(stack) > 0 {
		if point != nil {
			res = append(res, point.Val)
			stack = append(stack, point)
			point = point.Left
		} else {
			point = stack[len(stack)-1]
			point = point.Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

// @lc code=end
