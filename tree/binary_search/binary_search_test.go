package binary_search

import (
	"fmt"
	"testing"
	"daily/tree"
)

func TestBinarySearchNode(t *testing.T) {
	root := NewBinarySearchTree()
	root.Insert(3)
	root.Insert(1)
	root.Insert(5)
	root.Insert(4)
	root.Insert(2)
	root.Insert(6)
	root.Insert(0)
	root.Insert(8)

	root.Draw()

	fmt.Printf("树中序遍历结果为:")
	root.Print()

	n, err := root.Find(1)
	if err != nil {
		if err != tree.ErrNotExist {
			t.Error("错误的错误")
		}
	}
	fmt.Println("查找1的结果为:", n.GetValue())

	n, err = root.Find(9)
	if err != nil {
		if err != tree.ErrNotExist {
			t.Error("错误的错误")
		}
	}

	root.Delete(1)

	fmt.Printf("删除1后的树中序遍历结果为")
	root.Print()
}

func TestGenArr(t *testing.T) {
	val := []int{0, 1, 2, 3, 4, 5, 6}
	b := &binarySearchNode{
		LeftNode: &binarySearchNode{
			LeftNode: &binarySearchNode{
				LeftNode: &binarySearchNode{
					LeftNode:  nil,
					RightNode: nil,
					Value:     &val[0],
					Tp:        "",
				},
				RightNode: nil,
				Value:     &val[1],
				Tp:        "",
			},
			RightNode: &binarySearchNode{
				LeftNode: nil,
				RightNode: &binarySearchNode{
					LeftNode:  nil,
					RightNode: nil,
					Value:     &val[4],
					Tp:        "",
				},
				Value: &val[3],
				Tp:    "",
			},
			Value: &val[2],
			Tp:    "",
		},
		RightNode: &binarySearchNode{
			LeftNode:  nil,
			RightNode: nil,
			Value:     &val[6],
			Tp:        "",
		},
		Value: &val[5],
		Tp:    "root",
	}
	s := &sequential{}
	s.genArr(b)
	fmt.Println(s.arr)
}

