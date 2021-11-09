package hard

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 */

// 遍历链表，k个元素分割一次，进行反转，递归链接

func reverseKGroup(head *ListNode, k int) *ListNode {
	// k个元素链表的表头
	node := head
	// 取k个节点
	for i := 0; i < k; i++ {
		// 不够k个不翻转，直接返回head
		if node == nil {
			return head
		}
		node = node.Next
	}
	// 对k个节点进行反转，并返回头节点
	fmt.Println(head.Val)
	fmt.Println(node.Val)
	newHead := reverseHelper(head, node)
	// 对剩余链表进行k个一组翻转并返回头节点进行连接
	head.Next = reverseKGroup(node, k)

	return newHead
}

/**
 * first 当前组连标的开始节点
 * last 下一组链表的开始节点
 */
func reverseHelper(first *ListNode, last *ListNode) *ListNode {
	// 当前节点first翻转后应指向下一组链表的开始last
	prev := last
	// 从头遍历并翻转
	for first != last {
		// 暂存当前节点的下了一个节点
		temp := first.Next
		// 修改next指向上一个节点
		first.Next = prev
		// 记录上一个节点为当前节点
		prev = first
		// 当前节点移动到下一个节点
		first = temp
	}
	return prev
}

// @lc code=end
