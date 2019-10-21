package R_B

import (
	"daily/tree"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
)

const (
	black = 30
	red   = 31
)

type RBTreeNode struct {
	Value     int
	LeftNode  *RBTreeNode
	RightNode *RBTreeNode
	ParentNode *RBTreeNode
	Color     int
	Tp        string
	Height    int
}

func NewRBTree()*RBTreeNode{
	return &RBTreeNode{
		Value:     0,
		LeftNode:  nil,
		RightNode: nil,
		ParentNode: nil,
		Color:     red,
		Tp:        "root",
		Height:    1,
	}
}

func (r *RBTreeNode) Insert(value interface{}) (bool, error) {
	var v int
	var ok bool
	if v, ok = value.(int); !ok {
		return false, errors.New("expect a int type param")
	}

	if r.Tp == "root" && r.Color == red{ // 根节点赋值,且如果是红色直接置为黑色
		r.Value = v
		r.Color = black
		return false, nil
	}

	newNode, err := insert(v, r)
	if err != nil{
		log.Error("插入节点失败：", err.Error())
		return false, err
	}
	if newNode == nil {
		// 新插入节点为空，则证明没有新建节点，直接返回即可
		return false, nil
	}

	// 平衡操作
	if newRoot := balanceTree(newNode);newRoot != nil{
		//r.Value = newRoot.Value
		//r.LeftNode = newRoot.LeftNode
		//r.RightNode = newRoot.RightNode
		*r = *newRoot
	}


	return true, nil
}

func balanceTree(n *RBTreeNode) (*RBTreeNode){
	switch{
	case n.Tp == "root":
		n.Color = black
		return nil
	case n.ParentNode.Color == black:
		return nil
	case n.ParentNode.Color == red && (n.ParentNode.LeftNode != nil && n.ParentNode.RightNode != nil):
		// 父节点一定不为空，但叔节点不一定,需要判断出左子还是右子
		if n.ParentNode.LeftNode == n {
			if n.ParentNode.RightNode.Color == red {
				n.ParentNode.Color = black
				n.ParentNode.Color = black
				n.ParentNode.ParentNode.Color = red
				balanceTree(n.ParentNode.ParentNode)
			}
			if n.ParentNode.RightNode.Color == black {
				return revolve(n)
			}
		}
		if n.ParentNode.RightNode == n {
			if n.ParentNode.LeftNode.Color == red {
				n.ParentNode.Color = black
				n.ParentNode.Color = black
				n.ParentNode.ParentNode.Color = red
				balanceTree(n.ParentNode.ParentNode)
			}
			if n.ParentNode.LeftNode.Color == black {
				return revolve(n)
			}
		}
	case n.ParentNode.Color == red && (n.ParentNode.LeftNode != nil || n.ParentNode.RightNode != nil):
		return revolve(n)
	}
	return nil
}

func revolve(n *RBTreeNode) *RBTreeNode{
	var newRoot *RBTreeNode
	switch{
	case n == n.ParentNode.LeftNode && n.ParentNode == n.ParentNode.ParentNode.LeftNode:
		// left-left
		revolveNode := n.ParentNode.ParentNode
		revolveNodeParent := n.ParentNode.ParentNode.ParentNode
		// 交换父节点与祖父节点的颜色
		revolveNode.Color, n.ParentNode.Color = n.ParentNode.Color, revolveNode.Color


		// 断绝父节点与祖父节点关系
		if revolveNodeParent != nil{
			// revolveNodeParent不是根节点才断关系
			revolveNodeParent.LeftNode = n.ParentNode
		}else{
			revolveNode.Tp = ""
			n.ParentNode.Tp = "root"
			newRoot = n.ParentNode
		}
		n.ParentNode.ParentNode = revolveNodeParent



		// 断绝父节点与右子节点的关系
		revolveNode.LeftNode = n.ParentNode.RightNode
		if n.ParentNode.RightNode != nil{
			n.ParentNode.RightNode.ParentNode = revolveNode
		}

		// 建立与父节点间的新关系
		n.ParentNode.RightNode = revolveNode
		revolveNode.ParentNode = n.ParentNode

		//n.ParentNode.Draw()
		break
	case n == n.ParentNode.RightNode && n.ParentNode == n.ParentNode.ParentNode.LeftNode:
		// right-left
		// 1.左旋父节点，并将父节点作为当前节点，然后到left-left
		revolveNode := n.ParentNode    // 旋转节点
		revolveNodeParent := n.ParentNode.ParentNode  // 旋转节点的父节点

		// 断绝父节点与祖父节点关系
		revolveNodeParent.LeftNode = n
		n.ParentNode = revolveNodeParent

		// 断绝与左子节点的关系
		revolveNode.RightNode = n.LeftNode
		if n.LeftNode != nil{
			n.LeftNode.ParentNode = revolveNode
		}

		// 建立与父节点间的新关系
		n.LeftNode = revolveNode
		revolveNode.ParentNode = n
		newRoot = revolve(revolveNode)
		break
	case n == n.ParentNode.LeftNode && n.ParentNode == n.ParentNode.ParentNode.RightNode:
		// left-right
	case n == n.ParentNode.RightNode && n.ParentNode == n.ParentNode.ParentNode.RightNode:
		// right-right
	}
	return newRoot
}

func insert(v int, r *RBTreeNode) (*RBTreeNode, error){
	if v > r.Value {
		if r.RightNode == nil {
			if r.LeftNode == nil {
				r.Height += 1 // 如要新建节点，则需先判断兄弟节点是否存在，不存在的话父节点高度加1
			}
			r.RightNode = &RBTreeNode{
				Value:       v,
				LeftNode:    nil,
				RightNode:   nil,
				ParentNode:  r,
				Color:       red,
				Tp:          "",
				Height:      1,
			}
			return r.RightNode, nil
		}
		newNode, err := insert(v, r.RightNode)
		if newNode != nil {
			r.Height = r.RightNode.Height + 1 // 递归向下，如果有新建节点，则父节点高度等于子节点+1
		}
		return newNode, err
	}else if v < r.Value{
		if r.LeftNode == nil {
			if r.RightNode == nil {
				r.Height += 1 // 如要新建节点，则需先判断兄弟节点是否存在，不存在的话父节点高度加1
			}
			r.LeftNode = &RBTreeNode{
				Value:       v,
				LeftNode:    nil,
				RightNode:   nil,
				ParentNode:  r,
				Color:       red,
				Tp:          "",
				Height:      1,
			}
			return r.LeftNode, nil
		}
		newNode, err := insert(v, r.LeftNode)
		if newNode == nil {
			r.Height = r.LeftNode.Height + 1 // 递归向下，如果有新建节点，则父节点高度等于子节点+1
		}
		return newNode, err
	} else {
		log.Debug(v, "已存在")
		return nil, nil
	}
}

func (r *RBTreeNode) Delete(value interface{}) error {
	panic("implement me")
}

func (r RBTreeNode) Find(value interface{}) (tree.Tree, error) {
	panic("implement me")
}

func (r *RBTreeNode) GetValue() interface{} {
	panic("implement me")
}


type drawInfo struct{
	parentPos int
	currentNode *RBTreeNode
	direction int // 1左2右
}

func (b *RBTreeNode) Draw() {
	// 间距害的调，高度越高，两个孩子中间距离越大才行
	var interval = 4
	drawList := []drawInfo{}
	if b.Tp == "root"{
		pos := b.Height * interval
		fmt.Printf("\x1b[%dC", pos)
		fmt.Printf("\x1b[%dm\x1b[47m%d\x1b[0m\n", black, b.Value)

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
				fmt.Printf("\x1b[%dm\x1b[47m%d\x1b[0m", n.currentNode.Color, n.currentNode.Value)
				fmt.Print("\x1b[u") // 恢复光标位置 恢复光标和Attrs <ESC> 8
			}else{
				pos += interval
				// 光标移动到父节点位置，然后向右移动两列
				fmt.Printf("\x1b[%dC", n.parentPos)
				fmt.Printf("\x1b[%dC", n.currentNode.Height+ 1)
				fmt.Printf("\x1b[%dm\x1b[47m%d\x1b[0m", n.currentNode.Color, n.currentNode.Value)
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
