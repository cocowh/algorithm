package medium

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 设交换的两个节点为head和next，head连接后面交换完成的子链表，next连接head，完成交换
// 当无节点或只剩一个节点（head == nil || head.Next == nil）,无法进行交换，终止

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

// @lc code=end
