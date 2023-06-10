package algorithm

var minTreeDiff int

// 暴力求解 对每一个节点分别求出其与所有子节点的差的绝对值
func getMinimumDifference(root *TreeNode) int {
	minTreeDiff = 0
	inOrder(root)
	return int(minTreeDiff)
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	searchTree(node)
	inOrder(node.Left)
	inOrder(node.Right)
}

func searchTree(node *TreeNode) {
	l := []*TreeNode{}
	if node.Left != nil {
		l = append(l, node.Left)
	}
	if node.Right != nil {
		l = append(l, node.Right)
	}
	for len(l) != 0 {
		setMin(node, l[0])
		if l[0].Left != nil {
			l = append(l, l[0].Left)
		}
		if l[0].Right != nil {
			l = append(l, l[0].Right)
		}
		l = l[1:]
	}
}

func setMin(node1, node2 *TreeNode) {
	diff := abs(node1.Val - node2.Val)
	if minTreeDiff == 0 {
		minTreeDiff = diff
		return
	}

	if diff < minTreeDiff {
		minTreeDiff = diff
	}
	return
}

// 二叉搜索树性质：中序遍历得到升序的有序数组
var pre *TreeNode

func getMinimumDifference1(root *TreeNode) int {
	minTreeDiff = 0
	pre = nil
	var inOrder1 func(node *TreeNode)
	inOrder1 = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder1(node.Left)
		if pre == nil {
			pre = node
		} else {
			setMin(node, pre)
			pre = node
		}

		inOrder1(node.Right)
	}
	inOrder1(root)
	return minTreeDiff
}

//func getMinimumDifference(root *TreeNode) int {
//	ans, pre := math.MaxInt64, -1
//	var dfs func(*TreeNode)
//	dfs = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		dfs(node.Left)
//		if pre != -1 && node.Val-pre < ans {
//			ans = node.Val - pre
//		}
//		pre = node.Val
//		dfs(node.Right)
//	}
//	dfs(root)
//	return ans
//}
