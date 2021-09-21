package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	mid := head
	for i := 0; i < k; i++ {
		if mid != nil {
			mid = mid.Next
		} else {
			return nil
		}
	}
	for mid != nil {
		head = head.Next
		mid = mid.Next
	}
	return head
}
