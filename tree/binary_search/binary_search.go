package binary_search

import (
	"daily/tree"
	"errors"
	"fmt"
)

type binarySearchNode struct {
	LeftNode  *binarySearchNode
	RightNode *binarySearchNode
	Value     *int
	Tp        string
	Height    int
}

func NewBinarySearchTree() *binarySearchNode {
	return &binarySearchNode{
		LeftNode:  nil,
		RightNode: nil,
		Value:     nil,
		Tp:        "root",
		Height:    1,
	}
}

// Insert return: whether the new node is created successful and error
func (b *binarySearchNode) Insert(value interface{}) (bool, error) {
	if v, ok := value.(int); ok {
		if b.Value == nil { // 根节点初始化时是没值的，该情况针对根节点赋值,没有新建节点
			b.Value = &v
			return false, nil
		}
		if v < *b.Value {
			if b.LeftNode == nil {
				if b.RightNode == nil{
					b.Height += 1 // 建新左子树前判断右子树是否存在，如果不存在则是第一次建孩子，父节点高度先加1
				}
				b.LeftNode = new(binarySearchNode)
				b.LeftNode.Value = &v
				b.LeftNode.Height = 1
				return true, nil
			} else {
				// 如果递归，则判断递归是否有新建节点，如果建了，则父节点的高度要更新
				ok, err := b.LeftNode.Insert(value)
				if ok {
					b.Height = b.LeftNode.Height + 1
				}
				return ok, err
			}
		}

		if v > *b.Value {
			if b.RightNode == nil {
				if b.LeftNode == nil{
					b.Height += 1 // 建新右子树前判断左子树是否存在，如果不存在则是第一次建孩子，父节点高度先加1
				}
				b.RightNode = new(binarySearchNode)
				b.RightNode.Value = &v
				b.RightNode.Height = 1
				return true, nil
			} else {
				// 如果递归，则判断递归是否有新建节点，如果建了，则父节点的高度要更新
				ok, err := b.RightNode.Insert(value)
				if ok {
					b.Height = b.RightNode.Height + 1
				}
				return ok, err
			}
		}

		if v == *b.Value {
			return false, nil
		}
	}

	return false, errors.New("value should int")
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
		targetType, target := getTargetAndTargetType(b, v)
		//fmt.Println(fmt.Sprintf("类别：%d, parent：%+v, cur:%+v, direction:%d", targetType, parent, target, direction))
		switch {
		case targetType == 5:
			return tree.ErrNotExist
		case targetType == 1:
			target = nil
		case targetType == 2:
			if target.RightNode != nil{
				target.Value = target.RightNode.Value
				target.LeftNode = target.RightNode.LeftNode
				target.RightNode = target.RightNode.RightNode
			}
			if target.LeftNode != nil{
				target.Value = target.LeftNode.Value
				target.LeftNode = target.LeftNode.LeftNode
				target.RightNode = target.LeftNode.RightNode
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

func getTargetAndTargetType(b *binarySearchNode, v int) (int,  *binarySearchNode) {
	if b.Tp == "root" {
		if *b.Value == v {
			return 4, b
		}
	}
	if v < *b.Value {
		if b.LeftNode == nil {
			return 5, nil
		}
		if v == *b.LeftNode.Value {
			// 目标节点的两个子节点均为nil，则是第一种情况
			if b.LeftNode.LeftNode == nil && b.LeftNode.RightNode == nil {
				return 1, b.LeftNode
			}
			// 目标节点的两个子节点均不为nil，则是第3种情况
			if b.LeftNode.LeftNode != nil && b.LeftNode.RightNode != nil {
				return 3, b.LeftNode
			}
			// 其他则是第2种情况
			return 2, b.LeftNode
		}
		return getTargetAndTargetType(b.LeftNode, v)
	}

	if v > *b.Value {
		if b.RightNode == nil {
			return 5, nil
		}
		if v == *b.RightNode.Value {
			// 目标节点的两个子节点均为nil，则是第一种情况
			if b.RightNode.LeftNode == nil && b.RightNode.RightNode == nil {
				return 1, b.RightNode
			}
			// 目标节点的两个子节点均不为nil，则是第3种情况
			if b.RightNode.LeftNode != nil && b.RightNode.RightNode != nil {
				return 3, b.RightNode
			}
			// 其他则是第2种情况
			return 2, b.RightNode
		}
		return getTargetAndTargetType(b.RightNode, v)
	}

	return 5, nil
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


type drawInfo struct{
	parentPos int
	currentNode *binarySearchNode
	direction int // 1左2右
}

func (b *binarySearchNode) Draw() {
	// 间距害的调，高度越高，两个孩子中间距离越大才行
	var interval = 4
	drawList := []drawInfo{}
	if b.Tp == "root"{
		pos := b.Height * interval
		fmt.Printf("\x1b[%dC", pos)
		fmt.Println(*b.Value)

		if b.LeftNode != nil{
			drawList = append(drawList, drawInfo{parentPos: pos, currentNode: b.LeftNode, direction:1})
		}
		if b.RightNode != nil{
			drawList = append(drawList, drawInfo{parentPos: pos, currentNode: b.RightNode, direction:2})
		}
	}
	for {
		tmpDrawList := []drawInfo{}
		fmt.Print("\x1b[s")   // 当前光标在第n行的第一列，保存光标位置
		for _, n := range drawList{
			pos := n.parentPos
			if n.direction == 1{
				pos -= interval
				// 光标移动到父节点位置，然后向左移动两列
				fmt.Printf("\x1b[%dC", n.parentPos)
				fmt.Printf("\x1b[%dD", n.currentNode.Height+ 1)
				fmt.Print(*n.currentNode.Value)
				fmt.Print("\x1b[u") // 恢复光标位置 恢复光标和Attrs <ESC> 8
			}else{
				pos += interval
				// 光标移动到父节点位置，然后向右移动两列
				fmt.Printf("\x1b[%dC", n.parentPos)
				fmt.Printf("\x1b[%dC", n.currentNode.Height+ 1)
				fmt.Print(*n.currentNode.Value)
				fmt.Print("\x1b[u") // 恢复光标位置 恢复光标和Attrs <ESC> 8
			}

			if n.currentNode.LeftNode != nil{
				tmpDrawList = append(tmpDrawList, drawInfo{parentPos: pos, currentNode:n.currentNode.LeftNode, direction:1})
			}
			if n.currentNode.RightNode != nil{
				tmpDrawList = append(tmpDrawList, drawInfo{parentPos: pos, currentNode:n.currentNode.RightNode, direction:2})
			}
		}
		fmt.Print("\n")
		drawList = tmpDrawList
		tmpDrawList = nil
		if len(drawList) == 0{
			break
		}
	}

}

