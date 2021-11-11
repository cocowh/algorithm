package medium

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	// 找到中间节点
	midNode, fastNode := head, head
	for fastNode.Next != nil && fastNode.Next.Next != nil {
		midNode = midNode.Next
		fastNode = fastNode.Next.Next
	}
	// 反转后半部链表
	l1, l2 := head, midNode.Next
	midNode.Next = nil
	var prev, cur *ListNode = nil, l2
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	l2 = prev
	// 合并链表
	var l1Temp, l2Temp *ListNode
	for l1 != nil && l2 != nil {
		l1Temp = l1.Next
		l2Temp = l2.Next

		l1.Next = l2
		l1 = l1Temp

		l2.Next = l1
		l2 = l2Temp
	}
}

// @lc code=end
