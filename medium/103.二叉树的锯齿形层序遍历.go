package medium

func ZigzagLevelOrder(root *TreeNode) [][]int {
	return zigzagLevelOrder(root)
}

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func ZigzagLevelOrderByBFS(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q, flag := []*TreeNode{root}, false
	for len(q) > 0 {
		size := len(q)
		tmp := []*TreeNode{}
		lay := make([]int, size)
		j := size - 1
		for i := 0; i < size; i++ {
			root = q[0]
			q = q[1:]
			if !flag {
				lay[i] = root.Val
			} else {
				lay[j] = root.Val
				j--
			}
			if root.Left != nil {
				tmp = append(tmp, root.Left)
			}
			if root.Right != nil {
				tmp = append(tmp, root.Right)
			}

		}
		res = append(res, lay)
		flag = !flag
		q = tmp
	}
	return res
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	// return ZigzagLevelOrderByBFS(root)
	res := [][]int{}
	zigzagSearch(root, 0, &res)
	return res
}

func zigzagSearch(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	for len(*res) < depth+1 {
		*res = append(*res, []int{})
	}
	if depth%2 == 0 {
		(*res)[depth] = append((*res)[depth], root.Val)
	} else {
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	}
	zigzagSearch(root.Left, depth+1, res)
	zigzagSearch(root.Right, depth+1, res)
}

// @lc code=end
