package easy

import (
	"math"
)

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start

func MaxProfit1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min, maxPrice := prices[0], math.MinInt64
	for _, v := range prices {
		if v-min > maxPrice {
			maxPrice = v - min
		}
		if v < min {
			min = v
		}
	}
	return maxPrice
}

func MaxProfit2(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	res, stack := 0, []int{prices[0]}
	for _, v := range prices {
		if v > stack[len(stack)-1] {
			stack = append(stack, v)
		} else {
			index := len(stack) - 1
			for ; index >= 0; index-- {
				if stack[index] < v {
					break
				}
			}
			stack = stack[:index+1]
			stack = append(stack, v)
		}
		res = max(res, stack[len(stack)-1]-stack[0])
	}
	return res
}

func maxProfit(prices []int) int {
	return MaxProfit1(prices)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
