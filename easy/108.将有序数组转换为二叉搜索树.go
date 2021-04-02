package easy

func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST(nums)
}

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func helperSortArrayToBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helperSortArrayToBST(nums, left, mid-1)
	root.Right = helperSortArrayToBST(nums, mid+1, right)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helperSortArrayToBST(nums, 0, len(nums)-1)
}

// @lc code=end
