package main

import (
	"algorithm/company"
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
	// fmt.Println(medium.Exist([][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }, "ABCCED"))
	// fmt.Println(medium.MovingCount(14, 14, 5))
	// fmt.Println(easy.PrintNumbers(2))
	// fmt.Println(easy.Exchange([]int{1, 2, 3, 4}))
	company.TestComputeIntersect()
}

func PrintArrInt(arr []int) {
	for i, v := range arr {
		fmt.Println(i, " : ", v)
	}
}
