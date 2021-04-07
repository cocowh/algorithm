package medium

func KthSmallest(root *TreeNode, k int) int {
	return kthSmallest(root, k)
}

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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

func KthSmallestHelper(node *TreeNode, k int, count *int, res *int) {
	if node == nil {
		return
	}
	KthSmallestHelper(node.Left, k, count, res)
	*count++
	if *count == k {
		*res = node.Val
	}
	KthSmallestHelper(node.Right, k, count, res)
}

func kthSmallest(root *TreeNode, k int) int {
	// count, res := 0, 0
	// KthSmallestHelper(root, k, &count, &res)
	// return res
	return KthSmallestHelperByStack(root, k)
}

func KthSmallestHelperByStack(root *TreeNode, k int) int {
	stack, point := []*TreeNode{}, root
	for point != nil || len(stack) > 0 {
		for point != nil {
			stack = append(stack, point)
			point = point.Left
		}
		point = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return point.Val
		}
		point = point.Right
	}
	return 0
}

// @lc code=end
