package algorithm

//输入：[0,1,1]
//输出：[true,false,false]
//解释：
//输入数字为 0, 01, 011；也就是十进制中的 0, 1, 3 。只有第一个数可以被 5 整除，因此 answer[0] 为真。
func prefixesDivBy5(A []int) []bool {
	if len(A) == 0 {
		return nil
	}
	res := make([]bool, len(A))
	var cur int
	for i := 0; i < len(A); i++ {
		cur = (cur<<1 + A[i]) % 5
		if cur == 0 {
			res[i] = true
		}
	}
	return res
}
