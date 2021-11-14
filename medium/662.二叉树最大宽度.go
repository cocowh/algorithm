package medium

/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
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
func widthOfBinaryTree(root *TreeNode) int {
	lMap := map[int]int{}
	maxWidth := 0
	var helper func(node *TreeNode, level, pos int)
	helper = func(node *TreeNode, level, pos int) {
		if node == nil {
			return
		}
		if _, ok := lMap[level]; !ok {
			lMap[level] = pos
		}
		curWidth := pos - lMap[level] + 1
		if maxWidth < curWidth {
			maxWidth = curWidth
		}
		helper(node.Left, level+1, pos*2)
		helper(node.Right, level+1, pos*2+1)
	}
	helper(root, 0, 0)
	return maxWidth
}

// @lc code=end
