package easy

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// 快慢指针

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

// @lc code=end
