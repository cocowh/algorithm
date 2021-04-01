package medium

func BuildTree106(inorder []int, postorder []int) *TreeNode {
	return buildTree106(inorder, postorder)
}

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

func BuildTree106ByRecursion(inorder, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(int, int) *TreeNode
	build = func(inorderLeft, inorderRight int) *TreeNode {
		// 无剩余节点
		if inorderLeft > inorderRight {
			return nil
		}
		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := idxMap[val]
		root.Right = build(inorderRootIndex+1, inorderRight)
		root.Left = build(inorderLeft, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}

func buildTree106(inorder []int, postorder []int) *TreeNode {
	return BuildTree106ByRecursion(inorder, postorder)
}

// @lc code=end
