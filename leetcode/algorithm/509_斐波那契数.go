package algorithm

func fib(n int) int {
	if n < 2 {
		return n
	}
	f1 := 0
	f2 := 1
	res := 0
	for i := 2; i <= n; i++ {
		res = f1 + f2
		f1 = f2
		f2 = res
	}
	return res
}
