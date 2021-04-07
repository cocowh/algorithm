package medium

func SortedListToBST(head *ListNode) *TreeNode {
	return sortedListToBST(head)
}

/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head != nil && head.Next == nil {
		return &TreeNode{Val: head.Val, Left: nil, Right: nil}
	}
	preNode, midNode := findMid(head)
	if midNode == nil {
		return nil
	}
	if preNode != nil { // 中断前部分链表
		preNode.Next = nil
	}
	if midNode == head { // 遍历左叶子节点
		head = nil
	}
	return &TreeNode{Val: midNode.Val, Left: sortedListToBST(head), Right: sortedListToBST(midNode.Next)}
}

func findMid(head *ListNode) (midPre, mid *ListNode) { // 快慢指针找中间节点，返回中间节点的前节点是为了中断链表
	if head == nil || head.Next == nil {
		return nil, head
	}
	mid, temp := head, head
	for temp.Next != nil && temp.Next.Next != nil {
		midPre = mid
		mid = mid.Next
		temp = temp.Next.Next
	}
	return midPre, mid
}

// @lc code=end
