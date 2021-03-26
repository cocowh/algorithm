package hard

import (
	"reflect"
)

func FindSubstring(s string, words []string) []int {
	return findSubstring(s, words)
}

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
//
func FindSubstringByMapBeter(s string, words []string) []int {
	ret, wordNum := []int{}, len(words)
	if wordNum == 0 {
		return ret
	}
	wordLen := len(words[0])
	allWards := make(map[string]int)
	for _, word := range words {
		if _, ok := allWards[word]; ok {
			allWards[word]++
		} else {
			allWards[word] = 0
		}
	}
	for j := 0; j < wordLen; j++ {
		num := 0
		hasWords := make(map[string]int)
		for i := j; i < len(s)-wordNum*wordLen+1; i = i + wordLen {
			hasRemoved := false
			for num < wordNum {
				word := s[i+num*wordLen : i+(num+1)*wordLen]
				if _, ok := allWards[word]; ok {
					if _, ok := hasWords[word]; ok {
						hasWords[word]++
					} else {
						hasWords[word] = 0
					}
					if hasWords[word] > allWards[word] {
						hasRemoved = true
						removeNum := 0
						for hasWords[word] > allWards[word] {
							firstWord := s[i+removeNum*wordLen : i+(removeNum+1)*wordLen]
							hasWords[firstWord]--
							removeNum++
						}
						num = num - removeNum + 1
						i = i + (removeNum-1)*wordLen
						break
					}
				} else {
					hasWords = make(map[string]int)
					i = i + num*wordLen
					num = 0
					break
				}
				num++
			}
			if num == wordNum {
				ret = append(ret, i)
			}
			if num > 0 && !hasRemoved {
				firstWord := s[i : i+wordLen]
				hasWords[firstWord]--
				num--
			}
		}
	}
	return ret
}

// 遍历s字串提取并统计单词与words统计进行比较
func FindSubstringByMap(s string, words []string) []int {
	lengthWA := len(words)
	if lengthWA == 0 { // 空
		return []int{}
	}
	lengthS, lengthW := len(s), len(words[0])
	if lengthW*lengthWA > lengthS { //长度不够
		return []int{}
	}
	maxLen := lengthW * lengthWA
	mW := make(map[string]int, lengthWA)
	for _, v := range words { //统计words
		mW[v]++
	}
	ret := []int{}
	for i := 0; i < lengthS-maxLen+1; i++ {
		sT := s[i : i+maxLen]
		mS := make(map[string]int, lengthWA)
		for j := 0; j < maxLen; j += lengthW {
			word := sT[j : j+lengthW]
			mS[word]++
		}
		if reflect.DeepEqual(mW, mS) {
			ret = append(ret, i)
		}
	}
	return ret
}

func findSubstring(s string, words []string) []int {
	return FindSubstringByMapBeter(s, words)
	// lengthWA := len(words)
	// if lengthWA == 0 { // 空
	// 	return []int{}
	// }
	// lengthS, lengthW := len(s), len(words[0])
	// if lengthW*lengthWA > lengthS || lengthS == 0 { //长度不够
	// 	return []int{}
	// }

	// return ret
}

// @lc code=end
