package algorithm

var x, y, last *TreeNode

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	x = nil
	y = nil
	last = nil
	inOrderRecoverTree(root) // 中序遍历过程中找到异常的两个点
	x.Val, y.Val = y.Val, x.Val
}

//     3          1         123   321
//        2     3
//           1   2
func inOrderRecoverTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		last = inOrderRecoverTree(node.Left)
	}

	if last != nil && node.Val < last.Val {
		if x == nil {
			x = last
		}
		y = node
	}
	last = node

	if node.Right != nil {
		last = inOrderRecoverTree(node.Right)
	}

	return last
}
