package medium

func MyPow(x float64, n int) float64 {
	return myPow(x, n)
}

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	b, res := n, 1.0
	if b < 0 {
		x = 1 / x
		b = -b
	}
	for ; b > 0; b >>= 1 {
		if b&1 == 1 {
			res *= x
		}
		x *= x
	}
	return res
}
