package easy

func ReversePrint(head *ListNode) []int {
	return reversePrint(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	res := make([]int, length)
	for head != nil {
		length--
		res[length] = head.Val
		head = head.Next
	}
	return res
}
