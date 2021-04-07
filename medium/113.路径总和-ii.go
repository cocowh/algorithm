package medium

func PathSum(root *TreeNode, targetSum int) [][]int {
	return pathSum(root, targetSum)
}

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	pathSumHelper(root, targetSum, &res, []int{})
	return res
}

func pathSumHelper(node *TreeNode, targetSum int, res *[][]int, set []int) {
	if node == nil {
		return
	} else if targetSum-node.Val == 0 && node.Left == nil && node.Right == nil {
		temp := make([]int, len(set)+1)
		copy(temp, append(set, node.Val))
		*res = append(*res, temp)
	} else {
		pathSumHelper(node.Left, targetSum-node.Val, res, append(set, node.Val))
		pathSumHelper(node.Right, targetSum-node.Val, res, append(set, node.Val))
	}
}

// @lc code=end
