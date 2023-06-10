package algorithm

func rotateSlice(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	queue := make([]int, 0, k)
	copy(queue, nums[len(nums)-k:])
	copy(nums[k:], nums[:len(nums)-k])
	copy(nums[:k], queue)
	return
}

// 循环替代 当前索引为x 则x位置上的值要替换到(x+k)%n的位置上，x位置上的值不变等转一圈回来后更新
func rotateSlice1(nums []int, k int) {
	k %= len(nums)
	for start, count := 0, gcd(len(nums), k); start < count; start++ {
		pre, cur := nums[start], start
		for {
			to := (cur + k) % len(nums)
			tmp := nums[to]
			nums[to] = pre
			pre = tmp
			cur = to
			if to == start {
				break
			}
		}
	}
}

// 9, 15 欧几里得算法
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 5, 10
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
