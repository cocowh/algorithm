package medium

func BuildTree(preorder []int, inorder []int) *TreeNode {
	return buildTree(preorder, inorder)
}

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return BuildPreIn2TreeDFS(preorder, 0, len(preorder)-1, 0, inPos)
	// return buildTreeByRecursion(preorder, inorder)
	// return BuildTreeByIteration(preorder, inorder)
}

func BuildTreeByIteration(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{root}
	inorderIndex := 0
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

func BuildTreeByRecursion(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root, i := &TreeNode{preorder[0], nil, nil}, 0
	for ; i < len(inorder); i++ { // 每次都遍历找根节点在中序遍历中的位置，区分左右子树，可优化
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = BuildTreeByRecursion(preorder[1:len(inorder[:i])+1], inorder[:i+1])
	root.Right = BuildTreeByRecursion(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 优化
// preStart 前序遍历中当前根节点的位置，pre
// preEnd   前序遍历中
// inStart  根节点在中序遍历中的位置，确定左右子树
// inPos    先行处理为map，便于后续查找根节点在中序遍历中的位置

func BuildPreIn2TreeDFS(pre []int, preStart int, preEnd int, inStart int, inPos map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: pre[preStart]}
	rootIdx := inPos[pre[preStart]]                                                    // 找根节点位置
	leftLen := rootIdx - inStart                                                       // 左子树长度
	root.Left = BuildPreIn2TreeDFS(pre, preStart+1, preStart+leftLen, inStart, inPos)  //preStart+1 左子树根节点，preStart+leftLen 左子树右边界
	root.Right = BuildPreIn2TreeDFS(pre, preStart+leftLen+1, preEnd, rootIdx+1, inPos) //prestart+leftLen+1 右子树根节点，rootIdx+1
	return root
}

// @lc code=end
