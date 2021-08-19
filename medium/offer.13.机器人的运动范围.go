package medium

import "fmt"

func MovingCount(m int, n int, k int) int {
	return movingCount(m, n, k)
}

func movingCount(m int, n int, k int) int {
	count := 0
	for i := 0; i < m; i++ {
		if calSum(i) > k {
			return count
		}
		for j := 0; j < n; j++ {
			if calSum(j)+calSum(i) > k {
				continue
			}
			fmt.Println(i, j, calSum(j)+calSum(i))
			count++
		}
	}
	return count
}

func calSum(a int) int {
	sum := 0
	for a >= 10 {
		sum += a % 10
		a /= 10
	}
	sum += a
	return sum
}
