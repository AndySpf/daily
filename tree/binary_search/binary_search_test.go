package binary_search

import (
	"fmt"
	"testing"
	"uselog/tree"
)

func TestBinarySearchNode_Delete(t *testing.T) {
	root := NewBinarySearchTree()
	root.Insert(3)
	root.Insert(1)
	root.Insert(5)
	root.Insert(4)
	root.Insert(2)
	root.Insert(6)
	root.Insert(0)

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
					tp:        "",
				},
				RightNode: nil,
				Value:     &val[1],
				tp:        "",
			},
			RightNode: &binarySearchNode{
				LeftNode: nil,
				RightNode: &binarySearchNode{
					LeftNode:  nil,
					RightNode: nil,
					Value:     &val[4],
					tp:        "",
				},
				Value: &val[3],
				tp:    "",
			},
			Value: &val[2],
			tp:    "",
		},
		RightNode: &binarySearchNode{
			LeftNode:  nil,
			RightNode: nil,
			Value:     &val[6],
			tp:        "",
		},
		Value: &val[5],
		tp:    "root",
	}
	s := &sequential{}
	s.genArr(b)
	fmt.Println(s.arr)
}

// 数组传值，切片不扩容传引用，扩容则传值（因为底层指向数组的指针变了）
func TestSliceTransmit(t *testing.T) {
	i := [3]int{1, 2, 3}
	testS(i)
	fmt.Println(i)
}

func testS(i [3]int) {
	//i = append(i, 4, 5, 6, 7)
	i[1] = 10
}
