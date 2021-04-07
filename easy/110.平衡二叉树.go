package easy

func IsBalanced(root *TreeNode) bool {
	return isBalanced(root)
}

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	flag, _ := isBalancedHelper(root)
	return flag
}

func isBalancedHelper(node *TreeNode) (bool, int) { // 获取左右子树的最大高度，并判断子树是否平衡二叉树
	if node == nil {
		return true, 0
	}
	flagL, depthL := isBalancedHelper(node.Left)
	flagR, depthR := isBalancedHelper(node.Right)
	if flagL && flagR && depthL-depthR >= -1 && depthL-depthR <= 1 {
		if depthL > depthR {
			return true, depthL + 1
		}
		return true, depthR + 1
	}
	return false, 0
}

// @lc code=end
