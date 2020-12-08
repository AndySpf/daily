package algorithm

import (
	"strconv"
)

// 给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456, 579]。
//
//形式上，斐波那契式序列是一个非负整数列表 F，且满足：
//
//0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32 位有符号整数类型）；
//F.length >= 3；
//对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2] 成立。
//另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。
//
//返回从 S 拆分出来的任意一组斐波那契式的序列块，如果不能拆分则返回 []。
func splitIntoFibonacci(S string) []int {
	if len(S) < 2 {
		return nil
	}

	queue := []int{}
	var backFibonacci func(S string, pos int) bool
	backFibonacci = func(S string, pos int) bool {
		if l := len(queue); l >= 3 && queue[l-3]+queue[l-2] != queue[l-1] {
			return false
		}
		if pos == len(S) {
			return true
		}

		for i := pos + 1; i <= len(S); i++ {
			if i-pos > len(S)/2 { // 长度已经大于总长度的二分之一还长还没找到(奇数时可能等于但一定不大于)，返回false
				return false
			}
			number, err := strconv.Atoi(S[pos:i]) // 可以继续优化，每次循环*=10+S[i]-'0'
			if err != nil || number >= (1<<31-1) {
				return false // 其中一项已经大于int32最大值了，返回false
			}

			exit := false
			if S[pos] == '0' { // 如果是0则这一项只能是单个的0跳出循环
				queue = append(queue, 0)
				exit = true
			} else {
				queue = append(queue, number)
			}
			if backFibonacci(S, i) {
				return true
			}
			queue = queue[:len(queue)-1]
			if exit {
				break
			}
		}
		return false
	}
	if backFibonacci(S, 0) {
		return queue
	}
	return nil
}
