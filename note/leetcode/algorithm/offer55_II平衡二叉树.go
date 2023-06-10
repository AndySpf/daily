package algorithm

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历
func isBalanced(root *TreeNode) bool {
	return postorderraversal(root) != -1
}

func postorderraversal(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := postorderraversal(root.Left)
	if left == -1 {
		return -1
	}
	right := postorderraversal(root.Right)
	if right == -1 {
		return -1
	}
	diff := abs(left - right)
	if diff < 2 {
		return maxDepth(left, right) + 1
	} else {
		return -1
	}
}

func abs(x int) int {
	if x < 0 {
		return ^x + 1
	}
	return x
}

func maxDepth(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
