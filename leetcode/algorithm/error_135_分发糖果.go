package algorithm

// 老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//相邻的孩子中，评分高的孩子必须获得更多的糖果。评分相同则没要求
//那么这样下来，老师至少需要准备多少颗糖果呢？

//输入: [1,5,5,1]
//输出: 4
func candy(ratings []int) int {
	n := len(ratings)
	// ans 结果
	// inc 小的增序序列长度
	// dec 小的降序序列长度，用来给递减序列内孩子都加一个苹果。
	//     如果当前递减序列长度和上一个递增序列长度相等了
	//     （即递增序列最后一个孩子分到的糖果和递减序列第一个孩子一样多，这是不符合规则的）
	//     需要把递增序列最后一个归为递减序列中，给他加一个糖果
	// pre 记录上一个孩子分到的糖果，用来在递增序列中做自增计算或者重置
	ans, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] { // 当前分数大于等于前一个
			dec = 0                         // 递减序列归零
			if ratings[i] == ratings[i-1] { // 如果当前分数等于前一个，则pre置为1作为当前孩子应该分配的糖果
				pre = 1
			} else {
				pre++ // 如果当前分数大于前一个，将前一个自增1，作为当前孩子应该分配的糖果
			}
			ans += pre
			inc = pre
		} else {
			dec++ // 当前孩子分数小于前一个，则递减序列长度加1
			if dec == inc {
				dec++ // 如果递减长度 == 上一个递增长度，把他看做是递减序列的起始点
			}
			ans += dec // 每次加递减序列的长度即可，不用额外给这些孩子加糖果
			pre = 1    // 递减序列中认为上一个孩子永远是1
		}
	}
	return ans
}

// 我们可以将「相邻的孩子中，评分高的孩子必须获得更多的糖果」这句话拆分为两个规则，分别处理。
// 左规则：当ratings[i−1]<ratings[i] 时，i 号学生的糖果数量将比i−1号孩子的糖果数量多。
// 右规则：当ratings[i]>ratings[i+1] 时，i 号学生的糖果数量将比i+1号孩子的糖果数量多。
// 我们遍历该数组两次，处理出每一个学生分别满足左规则或右规则时，最少需要被分得的糖果数量。每个人最终分得的糖果数量即为这两个数量的最大值。
func candy1(ratings []int) int {
	left := make([]int, len(ratings))
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			left[i] = maxNum(left[i], left[i+1]+1)
		}
	}
	return sumInt(left)
}
