package medium

func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	return isSubStructure(A, B)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (helperOffer26(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func helperOffer26(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return helperOffer26(A.Left, B.Left) && helperOffer26(A.Right, B.Right)
}
