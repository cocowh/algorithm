package easy

func Fib(n int) int {
	return fib(n)
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	first, second, current := 0, 1, 0
	for i := 2; i <= n; i++ {
		current = (first + second) % 1000000007
		first = second
		second = current
	}
	return current
}
