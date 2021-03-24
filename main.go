package main

import (
	"algorithm/medium"
	"fmt"
)

var (
	ArrInt  = []int{1, 3, 4, 2, 2}
	StrChar = "qewrtyuiop"
)

func main() {
	// medium.FindDuplicate(ArrInt)
	fmt.Println(medium.LengthOfLongestSubstring(StrChar))
}
