package easy

func DeleteNode(head *ListNode, val int) *ListNode {
	return deleteNode(head, val)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	for ; head.Val == val; head = head.Next {
	}
	res := head
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return res
}
