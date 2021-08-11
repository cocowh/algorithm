package easy

func NumWays(n int) int {
	return numWays(n)
}

func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first, second, current := 1, 2, 2
	for i := 3; i <= n; i++ {
		current = (first + second) % 1000000007
		first = second
		second = current
	}
	return current
}
