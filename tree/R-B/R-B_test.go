package R_B

import (
	"fmt"
	"testing"
)

func TestRBTreeNode_Insert(t *testing.T) {
	data := []int{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}
	//data := []int{12, 1, 9, 2, 0, 11, 7, 19, 4}
	fmt.Println("data is ", data)
	root := NewRBTree()
	for _, v := range data {
		root.Insert(v)
		//fmt.Println("=================================")
		if root.ParentNode != nil {
			// 每次插入后，有可能会更换根节点，可通过判断当前根节点的父节点是否为空进行处理
			root = root.ParentNode
		}
	}

	root.Draw()

}

func TestRBTreeNode_Delete(t *testing.T) {
	data := []int{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}
	//data := []int{12, 1, 9, 2, 0, 11, 7, 19, 4}
	fmt.Println("data is ", data)
	root := NewRBTree()
	for _, v := range data {
		root.Insert(v)
		//fmt.Println("=================================")
		if root == nil {
			// 每次插入后，有可能会更换根节点，可通过判断当前根节点的父节点是否为空进行处理
			root = root.ParentNode
		}
	}

	data1 := []int{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}
	//data1 := []int{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}
	for _, v := range data1 {
		root.Delete(v)
		fmt.Println(root)
		if root.ParentNode != nil {
			// 每次插入后，有可能会更换根节点，可通过判断当前根节点的父节点是否为空进行处理
			root = root.ParentNode
		}
	}

	root.Draw()
}

func TestRBTreeNode_GetHeight(t *testing.T) {
	root := &RBTreeNode{
		Value: 1,
		LeftNode: &RBTreeNode{
			Value: 2,
			LeftNode: &RBTreeNode{
				Value:    4,
				LeftNode: nil,
				RightNode: &RBTreeNode{
					Value:      5,
					LeftNode:   nil,
					RightNode:  nil,
					ParentNode: nil,
					Color:      0,
					Tp:         "",
				},
				ParentNode: nil,
				Color:      0,
				Tp:         "",
			},
			RightNode:  nil,
			ParentNode: nil,
			Color:      0,
			Tp:         "",
		},
		RightNode: &RBTreeNode{
			Value:    3,
			LeftNode: nil,
			RightNode: &RBTreeNode{
				Value:      4,
				LeftNode:   nil,
				RightNode:  nil,
				ParentNode: nil,
				Color:      0,
				Tp:         "",
			},
			ParentNode: nil,
			Color:      0,
			Tp:         "",
		},
		ParentNode: nil,
		Color:      0,
		Tp:         "",
	}

	fmt.Println(root.GetHeight())
	fmt.Println(root.LeftNode.GetHeight())
	fmt.Println(root.RightNode.GetHeight())
}

func TestRightResolve(t *testing.T) {
	root := &RBTreeNode{
		Value:      7,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: nil,
		Color:      0,
		Tp:         "",
	}
	left1 := &RBTreeNode{
		Value:      5,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: root,
		Color:      0,
		Tp:         "",
	}
	right1 := &RBTreeNode{
		Value:      9,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: root,
		Color:      0,
		Tp:         "",
	}
	root.LeftNode = left1
	root.RightNode = right1
	left2 := &RBTreeNode{
		Value:      3,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: left1,
		Color:      0,
		Tp:         "",
	}
	right2 := &RBTreeNode{
		Value:      6,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: left1,
		Color:      0,
		Tp:         "",
	}
	left1.LeftNode = left2
	left1.RightNode = right2
	left3 := &RBTreeNode{
		Value:      2,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: left2,
		Color:      0,
		Tp:         "",
	}
	right3 := &RBTreeNode{
		Value:      4,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: left2,
		Color:      0,
		Tp:         "",
	}
	left2.LeftNode = left3
	left2.RightNode = right3

	rightRevolve(root.LeftNode)
}

func TestLeftResolve(t *testing.T) {
	root := &RBTreeNode{
		Value:      3,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: nil,
		Color:      0,
		Tp:         "root",
	}
	left1 := &RBTreeNode{
		Value:      2,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: root,
		Color:      0,
		Tp:         "",
	}
	right1 := &RBTreeNode{
		Value:      8,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: root,
		Color:      0,
		Tp:         "",
	}
	root.LeftNode = left1
	root.RightNode = right1
	left2 := &RBTreeNode{
		Value:      7,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: right1,
		Color:      0,
		Tp:         "",
	}
	right2 := &RBTreeNode{
		Value:      10,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: right1,
		Color:      0,
		Tp:         "",
	}
	right1.LeftNode = left2
	right1.RightNode = right2
	left3 := &RBTreeNode{
		Value:      9,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: right2,
		Color:      0,
		Tp:         "",
	}
	right3 := &RBTreeNode{
		Value:      11,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: right2,
		Color:      0,
		Tp:         "",
	}
	right2.LeftNode = left3
	right2.RightNode = right3

	res := leftRevolve(root)

	fmt.Println(res)
}
