package algorithm

//       1
//    2      3
//  4   5  6   7
// => 1324567
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	var exit, reverse bool

	for !exit {
		reverse = !reverse
		exit = true
		q1 := []*TreeNode{}
		res1 := []int{}
		for len(queue) != 0 {
			item := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			res1 = append(res1, item.Val)
			if item.Left != nil || item.Right != nil {
				exit = false
			}
			if reverse { // 显得有点笨重。可以像正常遍历一样，维护层数，将奇数层数据翻转
				if item.Left != nil {
					q1 = append(q1, item.Left)
				}
				if item.Right != nil {
					q1 = append(q1, item.Right)
				}
			} else {
				if item.Right != nil {
					q1 = append(q1, item.Right)
				}
				if item.Left != nil {
					q1 = append(q1, item.Left)
				}
			}
		}
		res = append(res, res1)
		queue = q1
	}

	return res
}
