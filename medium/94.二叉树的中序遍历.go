package medium

func InorderTraversal(root *TreeNode) []int {
	return inorderTraversal(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversalByStack(root *TreeNode) []int {
	res, stack, point := []int{}, []*TreeNode{}, root
	for point != nil || len(stack) > 0 {
		for point != nil {
			stack = append(stack, point)
			point = point.Left
		}
		point = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, point.Val)
		point = point.Right
	}
	return res
}

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// @lc code=end

// 中序遍历：左->根->右。

// 先输出左
// 回溯输出根
// 输出右
