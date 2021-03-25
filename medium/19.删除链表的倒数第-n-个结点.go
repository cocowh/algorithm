package medium

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head, n)
}

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for fast != nil { //快慢指针，间距n后同步移动
		fast = fast.Next
		if n < 0 {
			slow = slow.Next
		}
		n--
	}
	if n == 0 { // length = n, 删除第一个节点
		return head.Next
	}
	slow.Next = slow.Next.Next // 删除节点
	return head
}

// @lc code=end
