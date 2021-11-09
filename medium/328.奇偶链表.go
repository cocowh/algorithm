package medium

/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 奇链表
	evenHead := head.Next
	// 偶链表
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		// 偶链表next指向奇链表next
		odd.Next = even.Next
		// 移动偶链表至下个节点
		odd = odd.Next
		// 奇链表next指向偶链表next
		even.Next = odd.Next
		//移动奇链表指向下一个节点
		even = even.Next
	}
	// 拼接奇偶链表
	odd.Next = evenHead
	return head
}

// @lc code=end
