package medium

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// 先将 curr 的下一个节点记录为 next；
// 把 curr 的下一个节点指向 next 的下一个节点；
// 把 next 的下一个节点指向 pre 的下一个节点；
// 把 pre 的下一个节点指向 next。

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	// 待反转区域的第一个节点 left 的前一个节点
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// 待反转区域的第一个节点 left
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		// 指向curr的下一个节点
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

// @lc code=end
