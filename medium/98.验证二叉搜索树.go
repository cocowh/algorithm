package medium

func IsValidBST(root *TreeNode) bool {
	return isValidBST(root)
}

func IsValidBSTBy(root *TreeNode) bool {
	arr := []int{}
	inOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func helper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	return min < node.Val && node.Val < max && helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return helper(root, -1<<63, 1<<63-1)
}

// @lc code=end

// 在遍历树的同时保留结点的上界与下界，在比较时不仅比较子结点的值，也要与上下界比较。
