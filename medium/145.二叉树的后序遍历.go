package medium

func PostorderTraversal(root *TreeNode) []int {
	return postorderTraversal(root)
}

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func PostorderTraversalByDFS(root *TreeNode) []int {
	res := []int{}
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

func postorderTraversal(root *TreeNode) []int {
	res, stack, point := []int{}, []*TreeNode{}, root
	var prev *TreeNode
	for point != nil || len(stack) > 0 {
		for point != nil {
			stack = append(stack, point)
			point = point.Left
		}
		point = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if point.Right == nil || point.Right == prev {
			res = append(res, point.Val)
			prev = point
			point = nil
		} else {
			stack = append(stack, point)
			point = point.Right
		}
	}
	return res
}

// @lc code=end
