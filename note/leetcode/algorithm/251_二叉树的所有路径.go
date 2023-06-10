package algorithm

import (
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	paths = paths[:0]
	dfsTreePath(root, "")
	return paths
}

func dfsTreePath(root *TreeNode, path string) {
	if root == nil {
		return
	}

	path += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		paths = append(paths, path)
	}

	path += "->"
	dfsTreePath(root.Left, path)
	dfsTreePath(root.Right, path)
}

var (
	pathQueue []string
	nodeQueue []*TreeNode
)

func binaryTreePaths1(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	paths = paths[:0]
	pathQueue = []string{strconv.Itoa(root.Val)}
	nodeQueue = []*TreeNode{root}
	bfsTreePath()
	return paths
}

func bfsTreePath() {
	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]

		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
		}
	}
}
