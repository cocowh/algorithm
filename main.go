package main

import (
	"algorithm/easy"
	"fmt"
)

var (
	ArrInt     = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	StrChar    = "qewrtyuiop"
	StrTempOne = "wordgoodgoodgoodbestword"
	StrArr     = []string{"word", "good", "best", "word"}
)

func main() {
	// medium.FindDuplicate(ArrInt)
	// fmt.Println(medium.LengthOfLongestSubstring(StrChar))
	// easy.RemoveDuplicates(ArrInt)
	// PrintArrInt(ArrInt)
	// fmt.Println(hard.FindSubstring(StrTempOne, StrArr))
	// fmt.Println(medium.GrayCode(4))
	fmt.Println(easy.Fib(95))
}

func PrintArrInt(arr []int) {
	for i, v := range arr {
		fmt.Println(i, " : ", v)
	}
}
