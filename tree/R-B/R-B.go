package R_B

import "daily/tree"

const (
	black = 1
	red   = 2
)

type RBTreeNode struct {
	Value     int
	LeftNode  *RBTreeNode
	RightNode *RBTreeNode
	Color     int
	Tp        string
	Height    int
}

func NewRBTree()*RBTreeNode{
	return &RBTreeNode{
		Value:     0,
		LeftNode:  nil,
		RightNode: nil,
		Color:     black,
		Tp:        "root",
		Height:    1,
	}
}

func (R *RBTreeNode) Insert(value interface{}) (bool, error) {
	panic("implement me")
}

func (R *RBTreeNode) Delete(value interface{}) error {
	panic("implement me")
}

func (R RBTreeNode) Find(value interface{}) (tree.Tree, error) {
	panic("implement me")
}

func (R *RBTreeNode) GetValue() interface{} {
	panic("implement me")
}
