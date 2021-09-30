package medium

func AlevelOrder(root *TreeNode) []int {
	return alevelOrder(root)
}

func alevelOrder(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		if node != nil {
			res = append(res, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return res
}
