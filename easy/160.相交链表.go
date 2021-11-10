package easy

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// 根据条件在返回结果后，两个链表仍须保持原有的结构。，不能直接使用头节点遍历。
// 首先排序双层遍历，其时间复杂度O(m*n)
// 其次想到hash存节点，已存则代表有交点，时间复杂度O(max(n,m))，空间复杂度O(n+m-1)
// 最后参考双指针表尾互指，发现运行时间不理想，没考虑不相交多次互指情景，抄标准答案，加上互指状态位标记

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB, hasLinkA, hasLinkB := headA, headB, false, false
	for pA != nil && pB != nil {
		if pA == pB {
			return pA
		}
		pA, pB = pA.Next, pB.Next
		if pA == nil && !hasLinkB {
			pA = headB
			hasLinkB = true
		}
		if pB == nil && !hasLinkA {
			pB = headA
			hasLinkA = true
		}
	}
	return nil
}

// @lc code=end
