package hard

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	const intMax = int(^uint(0) >> 1)    //定义最小整数常量
	const intMin = ^intMax               //定义最大整数常量
	len1, len2 := len(nums1), len(nums2) //计算数组长度
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1) //对长度小的作二分提高效率
	}
	L1, L2, R1, R2, c1, c2, lo, hi := 0, 0, 0, 0, 0, 0, 0, 2*len1+1 //割两个数组参数初始化，虚拟加长使长度变奇数为2*len1 + 1
	for lo <= hi {
		c1 = (hi + lo) / 2    //数组1二分（数组1初始割点）
		c2 = len1 + len2 - c1 //数组2分割（数组2初始割点）
		if c1 == 0 {          //数组1割在零位（两个数组为空数组）
			L1 = intMin //左1为空取L1为整数最小值
		} else {
			L1 = nums1[(c1-1)/2] //左1有值取其左边最大值（前面加*2 + 1，此处-1/2）
		}
		if c1 == 2*len1 { //右1为空取R1为整数最大值
			R1 = intMax
		} else {
			R1 = nums1[c1/2] //不为空取R1为右边最小值
		}
		if c2 == 0 {
			L2 = intMin //同上
		} else {
			L2 = nums2[(c2-1)/2]
		}
		if c2 == 2*len2 {
			R2 = intMax //同上
		} else {
			R2 = nums2[c2/2] //同上
		}

		if L1 > R2 {
			hi = c1 - 1 //若不满足L1 < R2说明左1太大（多），割点左移（向左二分）
		} else if L2 > R1 {
			lo = c1 + 1 //若不满足L2 < R1说明左2太大（多），割点左移（向左二分）
		} else {
			break
		}
	}
	return ((math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2.0) //当满足L1 < R2 && L2 < R1时说明此时切割的四份已经有序，取有序后中间两个数的平均数即为中位数
}

// @lc code=end

// 容易想到的是对两个数组遍历进行排序然后取中位数，但是数间复杂度为O(m+n)，不符合要求O(log(m+n))。

// 此题解法目前参考他人思路[【分步详解】两个有序数组中的中位数和Top K问题](https://blog.csdn.net/hk2291976/article/details/51107778)，是对官方题解的一种更清晰易懂的解释说明。

// 参考[官方](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/)
