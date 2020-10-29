package algorithm

// 根到叶子节点路径所代表的数字的和
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	return dfsSumNumbers(root, 0, 0, sum)
}

func dfsSumNumbers(node *TreeNode, num int, height int, sum int) int {
	// 23 => 423
	num = num*10 + node.Val

	if node.Left == nil && node.Right == nil {
		sum += num
	}

	if node.Left != nil {
		sum = dfsSumNumbers(node.Left, num, height+1, sum)
	}
	if node.Right != nil {
		sum = dfsSumNumbers(node.Right, num, height+1, sum)
	}
	return sum
}

func sumNumbers1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return bfsSumNumbers([]*TreeNode{root}, []int{root.Val})

}

func bfsSumNumbers(nodeQueue []*TreeNode, pathQueue []int) int {
	sum := 0
	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path*10+node.Left.Val)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path*10+node.Right.Val)
		}
		if node.Left == nil && node.Right == nil {
			sum += pathQueue[i]
		}
	}
	return sum
}
