package algorithm

//输入：[1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
func findMaxAverage(nums []int, k int) float64 {
	var max int
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	sum := max
	for i := 1; i < len(nums)-k+1; i++ {
		sum := sum - nums[i-1] + nums[i+k-1]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
