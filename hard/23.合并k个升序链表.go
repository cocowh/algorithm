package hard

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 二分、分治
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	mid := length / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoList(left, right)
}

// 合并两个有序链表
func mergeTwoList(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	head := &ListNode{0, nil}
	tail, aPtr, bPtr := head, a, b
	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		tail = tail.Next
	}
	if aPtr == nil {
		tail.Next = bPtr
	} else {
		tail.Next = aPtr
	}
	return head.Next
}

// @lc code=end
