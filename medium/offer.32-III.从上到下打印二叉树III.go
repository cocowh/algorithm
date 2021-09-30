package medium

func ZlevelOrder(root *TreeNode) [][]int {
	return zlevelOrder(root)
}

func zlevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(res) {
			res = append(res, []int{})
		}
		if level&1 == 0 {
			res[level] = append(res[level], node.Val)
		} else {
			res[level] = append([]int{node.Val}, res[level]...)
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return res
}
