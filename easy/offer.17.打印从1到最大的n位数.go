package easy

import (
	"math"
)

func PrintNumbers(n int) []int {
	return printNumbers(n)
}

func printNumbers(n int) []int {
	cap := int(math.Pow10(n)) - 1
	res := make([]int, cap)
	for i := 0; i < cap; i++ {
		res[i] = i + 1
	}
	return res
}
