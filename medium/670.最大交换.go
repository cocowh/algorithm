package medium

import "strconv"

/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	bs := []byte(strconv.Itoa(num))

	// indexs 记录了每个数字最后出现的位置
	indexs := make(map[byte]int, len(bs))
	for i := range bs {
		indexs[bs[i]] = i
	}

	// 高位上的数字与低位上的更大的数字交换，才能使得 num 变大
	for i := 0; i < len(bs); i++ {
		for bj := byte('9'); bs[i] < bj; bj-- {
			j := indexs[bj]
			if j > i {
				bs[i], bs[j] = bs[j], bs[i]
				return convert670(bs)
			}
		}
	}
	return convert670(bs)
}

func convert670(bs []byte) int {
	n, _ := strconv.Atoi(string(bs))
	return n
}

// @lc code=end
