package medium

func RotateRight(head *ListNode, k int) *ListNode {
	return rotateRight(head, k)
}

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	length, tailNode := 0, head
	for tailNode.Next != nil {
		length++
		tailNode = tailNode.Next
	}
	length++                                        // 弥补最后一个节点，未计算
	tailNode.Next = head                            // 形成环
	for i := 0; i < (length-k%length)%length; i++ { // 右移1位，则对应要遍历length-1位找到断开点
		tailNode = tailNode.Next
	}
	head = tailNode.Next
	tailNode.Next = nil
	return head
}

// @lc code=end

// 1、获取链表长度及链尾节点，首位相连形成环
// 2、向右移动k个位置，则延链表遍历(length-k%length)%length（%length，优化case k=lenght）个节点，为断开点
