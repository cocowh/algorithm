package medium

import (
	"fmt"
	"math"
)

func VrifyPostorder(postorder []int) bool {
	return verifyPostorder(postorder)
}

func verifyPostorder(postorder []int) bool {
	stack, root := []int{}, math.MaxInt64
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) != 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
		}
		stack = append(stack, postorder[i])
	}
	return true
}
