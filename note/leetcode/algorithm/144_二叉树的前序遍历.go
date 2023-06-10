package algorithm

var preorderTraversalRes []int

// 递归
func preorderTraversal(root *TreeNode) []int {
	preorderTraversalRes = []int{}
	preOrder(root)
	return preorderTraversalRes
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	preorderTraversalRes = append(preorderTraversalRes, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

// 迭代
func preorderTraversal1(root *TreeNode) []int {
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
	}
	return res
}
