package medium

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return headNew
}

// @lc code=end
