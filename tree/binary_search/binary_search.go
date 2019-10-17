package binary_search

import (
	"errors"
	"fmt"
	"uselog/tree"
)

type binarySearchNode struct {
	LeftNode  *binarySearchNode
	RightNode *binarySearchNode
	Value     *int
	tp        string
}

func NewBinarySearchTree() *binarySearchNode {
	return &binarySearchNode{
		LeftNode:  nil,
		RightNode: nil,
		Value:     nil,
		tp:        "root",
	}
}

func (b *binarySearchNode) Insert(value interface{}) error {
	if v, ok := value.(int); ok {
		if b.Value == nil {
			b.Value = &v
			return nil
		}
		if v < *b.Value {
			if b.LeftNode == nil {
				b.LeftNode = new(binarySearchNode)
				b.LeftNode.Value = &v
				return nil
			} else {
				return b.LeftNode.Insert(value)
			}
		}

		if v > *b.Value {
			if b.RightNode == nil {
				b.RightNode = new(binarySearchNode)
				b.RightNode.Value = &v
				return nil
			} else {
				return b.RightNode.Insert(value)
			}
		}

		if v == *b.Value {
			return nil
		}
	}

	return errors.New("value should int")
}

func (b *binarySearchNode) Find(value interface{}) (tree.Tree, error) {
	if v, ok := value.(int); ok {
		if v == *b.Value {
			return b, nil
		}
		if v < *b.Value {
			if b.LeftNode == nil {
				return nil, tree.ErrNotExist
			}
			return b.LeftNode.Find(value)
		}

		if v > *b.Value {
			if b.RightNode == nil {
				return nil, tree.ErrNotExist
			}
			return b.RightNode.Find(value)
		}
	}
	return nil, errors.New("value should int")
}

func (b *binarySearchNode) GetValue() interface{} {
	return *b.Value
}

func (b *binarySearchNode) Delete(value interface{}) error {

	if v, ok := value.(int); ok {
		// targetType:1.没有子节点的叶子节点;2.只有一侧树的节点;3.有两个子节点的节点;4.特殊的3:根节点;5.不存在的点
		// direction: 1.目标节点是其父节点的左子节点。2.目标节点是其父节点的右子节点
		targetType, parent, target, direction := getParentNode(b, v)
		//fmt.Println(fmt.Sprintf("类别：%d, parent：%+v, cur:%+v, direction:%d", targetType, parent, target, direction))
		switch {
		case targetType == 5:
			return tree.ErrNotExist
		case targetType == 1 && direction == 1: // 删除target与父节点的关系
			parent.LeftNode = nil
		case targetType == 1 && direction == 2:
			parent.RightNode = nil
		case targetType == 2 && direction == 1: // 将target与父节点的关系迁移至target子节点与父节点中
			if target.LeftNode != nil {
				parent.LeftNode = target.LeftNode
			} else {
				parent.LeftNode = target.RightNode
			}
		case targetType == 2 && direction == 2:
			if target.LeftNode != nil {
				parent.RightNode = target.LeftNode
			} else {
				parent.RightNode = target.RightNode
			}
		case targetType == 3: // 从右节点开始找到最小后继，将最小后继脱离树，然后将其代替target
			rightLeastNode := getReplaceNodeAndDeleteRelation(target, target.RightNode, 2)
			target.Value = rightLeastNode.Value
		case targetType == 4:
			rightLeastNode := getReplaceNodeAndDeleteRelation(target, target.RightNode, 2)
			b.Value = rightLeastNode.Value
		}
	}

	return nil
}

func getParentNode(b *binarySearchNode, v int) (int, *binarySearchNode, *binarySearchNode, int8) {
	if b.tp == "root" {
		if *b.Value == v {
			return 4, nil, b, 0
		}
	}
	if v < *b.Value {
		if b.LeftNode == nil {
			return 5, nil, nil, 0
		}
		if v == *b.LeftNode.Value {
			// 目标节点的两个子节点均为nil，则是第一种情况
			if b.LeftNode.LeftNode == nil && b.LeftNode.RightNode == nil {
				return 1, b, b.LeftNode, 1
			}
			// 目标节点的两个子节点均不为nil，则是第3种情况
			if b.LeftNode.LeftNode != nil && b.LeftNode.RightNode != nil {
				return 3, b, b.LeftNode, 1
			}
			// 其他则是第2种情况
			return 2, b, b.LeftNode, 1
		}
		return getParentNode(b.LeftNode, v)
	}

	if v > *b.Value {
		if b.RightNode == nil {
			return 5, nil, nil, 2
		}
		if v == *b.RightNode.Value {
			// 目标节点的两个子节点均为nil，则是第一种情况
			if b.RightNode.LeftNode == nil && b.RightNode.RightNode == nil {
				return 1, b, b.RightNode, 2
			}
			// 目标节点的两个子节点均不为nil，则是第3种情况
			if b.RightNode.LeftNode != nil && b.RightNode.RightNode != nil {
				return 3, b, b.RightNode, 2
			}
			// 其他则是第2种情况
			return 2, b, b.RightNode, 2
		}
		return getParentNode(b.RightNode, v)
	}

	return 5, nil, nil, 0
}

func getReplaceNodeAndDeleteRelation(parent, node *binarySearchNode, direction int) *binarySearchNode {
	if node.LeftNode == nil {
		// 删除右最小节点与其父节点见关系
		if direction == 1 {
			parent.LeftNode = nil
		} else {
			parent.RightNode = nil
		}
		return node
	}
	return getReplaceNodeAndDeleteRelation(node, node.LeftNode, 1)
}

type sequential struct {
	arr []int
}

// Print 中序遍历当前树
func (b *binarySearchNode) Print() {
	s := &sequential{}
	s.genArr(b)
	fmt.Println(s.arr)
}

func (s *sequential) genArr(b *binarySearchNode) {
	if b.LeftNode != nil {
		s.genArr(b.LeftNode)
	}
	s.arr = append(s.arr, *b.Value)
	if b.RightNode != nil {
		s.genArr(b.RightNode)
	}
}
