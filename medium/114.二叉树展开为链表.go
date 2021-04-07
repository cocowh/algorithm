package medium

func Flatten(root *TreeNode) {
	flatten(root)
	flatten1(root)
}

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
// 对先序遍历结果进行处理
func flatten1(root *TreeNode) {
	// list := PreorderTraversalByDFSHelper(root)
	list := PreorderTraversalByStackHelper(root)
	for i := 1; i < len(list); i++ {
		list[i-1].Left, list[i-1].Right = nil, list[i]
	}
}

// 递归前序遍历
func PreorderTraversalByDFSHelper(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, PreorderTraversalByDFSHelper(root.Left)...)
		list = append(list, PreorderTraversalByDFSHelper(root.Right)...)
	}
	return list
}

// 利用栈进行递归迭代
func PreorderTraversalByStackHelper(root *TreeNode) []*TreeNode {
	list, stack, point := []*TreeNode{}, []*TreeNode{}, root
	for point != nil || len(stack) > 0 {
		if point != nil {
			list = append(list, point)
			stack = append(stack, point)
			point = point.Left
		} else {
			point = stack[len(stack)-1]
			point = point.Right
			stack = stack[:len(stack)-1]
		}
	}
	return list
}

// 利用栈前序遍历时展开 保存前序遍历的上一个节点，对其左右自节点进行操作
func FlattenByStack(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var prev *TreeNode
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Left, prev.Right = nil, curr
		}
		left, right := curr.Left, curr.Right
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
		prev = curr
	}
}

// 找根节点的左子树最后节点作为右子树的前驱节点
func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

// @lc code=end
