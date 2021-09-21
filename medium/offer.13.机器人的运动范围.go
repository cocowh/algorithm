package medium

func MovingCount(m int, n int, k int) int {
	return movingCount(m, n, k)
}

func getSum(i int) int {
	res := 0
	for ; i > 0; i = i / 10 {
		res += i % 10
	}
	return res
}

func movingCount(m int, n int, k int) int {
	if k <= 0 {
		return 1
	}
	v, res := make([][]int, m), 1
	for i := range v {
		v[i] = make([]int, n)
	}
	v[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || getSum(i)+getSum(j) > k {
				continue
			}
			if i-1 >= 0 {
				v[i][j] |= v[i-1][j]
			}
			if j-1 >= 0 {
				v[i][j] |= v[i][j-1]
			}
			res += v[i][j]
		}
	}
	return res
}
