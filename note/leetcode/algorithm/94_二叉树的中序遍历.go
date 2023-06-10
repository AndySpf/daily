package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var inorderTraversalRes []int

func inorderTraversal(root *TreeNode) []int {
	inorderTraversalRes = []int{}
	inorder(root)
	return inorderTraversalRes
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	inorderTraversalRes = append(inorderTraversalRes, root.Val)
	inorder(root.Right)
}

type myInorer struct {
	Color int
	Node  *TreeNode
}

// 栈迭代
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return inorderTraversalRes
	}

	inorderStack := []myInorer{
		{
			Color: 0, // 0:未遍历过  1:遍历过
			Node:  root,
		},
	}
	for {
		if len(inorderStack) == 0 {
			break
		}
		cur := inorderStack[len(inorderStack)-1]
		inorderStack = inorderStack[:len(inorderStack)-1]

		if cur.Color == 1 {
			inorderTraversalRes = append(inorderTraversalRes, cur.Node.Val)
			continue
		} else {
			if cur.Node.Right != nil {
				inorderStack = append(inorderStack, myInorer{Color: 0, Node: cur.Node.Right})
			}

			inorderStack = append(inorderStack, myInorer{1, cur.Node})

			if cur.Node.Left != nil {
				inorderStack = append(inorderStack, myInorer{0, cur.Node.Left})
			}
		}
	}
	return inorderTraversalRes
}
