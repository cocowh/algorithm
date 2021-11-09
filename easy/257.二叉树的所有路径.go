package easy

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	var helper func(node *TreeNode, prex string)
	helper = func(node *TreeNode, prex string) {
		if node.Left == nil && node.Right == nil {
			res = append(res, prex)
			return
		}
		if node.Left != nil {
			helper(node.Left, prex+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			helper(node.Right, prex+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	helper(root, strconv.Itoa(root.Val))
	return res
}

// @lc code=end
