package medium

func CuttingRopeII(n int) int {
	return cuttingRopeII(n)
}

func cuttingRopeII(n int) int {
	if n < 4 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return calPow(3, a) % 1000000007
	} else if b == 1 {
		return calPow(3, a-1) * 4 % 1000000007
	} else {
		return calPow(3, a) * 2 % 1000000007
	}
}

func calPow(base int, num int) int {
	res := 1
	for ; num > 0; num >>= 1 {
		if num&1 == 1 {
			res *= base
			res %= 1000000007
		}
		base *= base
		base %= 1000000007
	}
	return res
}
