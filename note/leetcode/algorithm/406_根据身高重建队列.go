package algorithm

import (
	"sort"
)

// 输入:
//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
//输出:
//[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
// 按照身高降序（k升序）排列后，遍历数组，将每一个遍历到的people[i]放入距离队列头people[i][1]的位置
// 因为是降序，所以后确定位置的人是不会影响已经确定位置的人的相对顺序的.所以按照插空思想只用实现插入函数
// 按位置插入即可
func reconstructQueue(people [][]int) [][]int {
	if len(people) <= 1 {
		return people
	}
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		// 如果k相同，则要按照k升序。保证后面的插入时不能影响前面的
		// 如果k为2的先排序，则k为1插入位置是会影响k为2的相对位置的。反之则不会
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	var insert func(index int, sli [][]int, pos int)
	insert = func(index int, sli [][]int, pos int) {
		tmp := sli[index]
		copy(sli[pos+1:index+1], sli[pos:index])
		sli[pos] = tmp
	}

	for i, item := range people {
		insert(i, people, item[1])
	}
	return people
}
