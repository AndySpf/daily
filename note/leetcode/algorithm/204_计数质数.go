package algorithm

// 统计所有小于非负整数 n 的质数的数量。
// 质数：除了1和他本身之外没有任何因子的数
// 11 * 13 = 143  11是因数则143/13必然也是其因数,因此选择两个中的较小值判断即可
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	judge := func(num int) bool {
		for i := 3; i*i <= num; i += 2 {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	count := 1 // 2
	for i := 3; i < n; i += 2 {
		if judge(i) {
			count++
		}
	}
	return count
}

func countPrimes1(n int) int {
	if n < 3 {
		return 0
	}

	isNotPrime := make([]bool, n)

	count := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
			for j := 2; j*i < n; j++ { // 如果n为质数，则2n,3n,4n... 必然不是质数
				isNotPrime[i*j] = true
			}
		}
	}
	return count
}

// 官方骚气的O(n)解法，线性筛选
func countPrimes2(n int) int {
	primes := []int{}
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}
