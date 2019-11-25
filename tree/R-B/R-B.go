package R_B

import (
	"daily/tree"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
)

// TODO 旋转涉及到根节点改变的情况能否优化

const (
	black = 30
	red   = 31
)

type RBTreeNode struct {
	Value      int
	LeftNode   *RBTreeNode
	RightNode  *RBTreeNode
	ParentNode *RBTreeNode
	Color      int
	Tp         string
}

func NewRBTree() *RBTreeNode {
	return &RBTreeNode{
		Value:      0,
		LeftNode:   nil,
		RightNode:  nil,
		ParentNode: nil,
		Color:      red,
		Tp:         "root",
	}
}

func (r *RBTreeNode) Insert(value interface{}) (bool, error) {
	var v int
	var ok bool
	if v, ok = value.(int); !ok {
		return false, errors.New("expect a int type param")
	}

	if r.Tp == "root" && r.Color == red { // 根节点赋值,且如果是红色直接置为黑色
		r.Value = v
		r.Color = black
		return false, nil
	}

	newNode, err := insert(v, r)
	if err != nil {
		log.Error("插入节点失败：", err.Error())
		return false, err
	}
	if newNode == nil {
		// 新插入节点为空，则证明没有新建节点，直接返回即可
		return false, nil
	}

	// 平衡操作
	balanceTree(newNode)

	return true, nil
}

func (r *RBTreeNode) Delete(value interface{}) error {
	var v int
	var ok bool
	if v, ok = value.(int); !ok {
		return errors.New("value must int")
	}

	deleteNode(r, v)
	return nil
}

func (r *RBTreeNode) Find(value interface{}) (tree.Tree, error) {
	if v, ok := value.(int); ok {
		if v == r.Value {
			return r, nil
		}
		if v < r.Value {
			if r.LeftNode == nil {
				return nil, tree.ErrNotExist
			}
			return r.LeftNode.Find(value)
		}

		if v > r.Value {
			if r.RightNode == nil {
				return nil, tree.ErrNotExist
			}
			return r.RightNode.Find(value)
		}
	}
	return nil, errors.New("value should int")
}

func (r *RBTreeNode) GetValue() interface{} {
	panic("implement me")
}

func balanceTree(n *RBTreeNode) {
	switch {
	case n.Tp == "root":
		// 父节点为根节点，则将其置为黑色
		n.Color = black
		break
	case n.ParentNode.Color == black:
		// 父节点为黑色，证明是正确的红黑树,不用处理
		break
	case n.ParentNode.Color == red && n.ParentNode.ParentNode.LeftNode == n.ParentNode && n.ParentNode.ParentNode.RightNode != nil:
		// 父节点为红色，且父节点是左节点，且叔节点不为空，然后根据叔节点颜色走不同方案
		if n.ParentNode.ParentNode.RightNode.Color == red {
			// 叔节点为红
			n.ParentNode.Color = black
			n.ParentNode.ParentNode.RightNode.Color = black
			n.ParentNode.ParentNode.Color = red
			balanceTree(n.ParentNode.ParentNode)
		} else {
			// 叔节点为黑
			revolve(n)
		}
		break
	case n.ParentNode.Color == red && n.ParentNode.ParentNode.RightNode == n.ParentNode && n.ParentNode.ParentNode.LeftNode != nil:
		// 父节点为红色，且父节点是右节点，且叔节点不为空，然后根据叔节点颜色走不同方案
		if n.ParentNode.ParentNode.LeftNode.Color == red {
			n.ParentNode.Color = black
			n.ParentNode.ParentNode.LeftNode.Color = black
			n.ParentNode.ParentNode.Color = red
			balanceTree(n.ParentNode.ParentNode)
		} else {
			revolve(n)
		}
		break
	case n.ParentNode.Color == red && (n.ParentNode.ParentNode.LeftNode != nil || n.ParentNode.ParentNode.RightNode != nil):
		// 父节点为红色，且叔节点为空（空节点认为是黑色）
		revolve(n)
		break
	}
	return
}

func revolve(n *RBTreeNode) {
	switch {
	case n == n.ParentNode.LeftNode && n.ParentNode == n.ParentNode.ParentNode.LeftNode:
		// left-left
		n.ParentNode.ParentNode.Color, n.ParentNode.Color = n.ParentNode.Color, n.ParentNode.ParentNode.Color
		rightRevolve(n.ParentNode.ParentNode)
		break
	case n == n.ParentNode.RightNode && n.ParentNode == n.ParentNode.ParentNode.LeftNode:
		// left-right 左旋父节点，并将父节点作为当前节点，然后到left-left
		leftRevolve(n.ParentNode)
		// 左旋后n的父节点就是旋转前n的祖父节点，对其再右旋即可
		n.ParentNode.Color, n.Color = n.Color, n.ParentNode.Color
		rightRevolve(n.ParentNode)
		break
	case n == n.ParentNode.LeftNode && n.ParentNode == n.ParentNode.ParentNode.RightNode:
		// right-left 右旋父节点，并将父节点作为当前节点，然后到left-left
		rightRevolve(n.ParentNode)
		// 右旋后n的父节点就是旋转前n的祖父节点，对其再左旋即可
		n.ParentNode.Color, n.Color = n.Color, n.ParentNode.Color
		leftRevolve(n.ParentNode)
		break
	case n == n.ParentNode.RightNode && n.ParentNode == n.ParentNode.ParentNode.RightNode:
		// right-right
		n.ParentNode.ParentNode.Color, n.ParentNode.Color = n.ParentNode.Color, n.ParentNode.ParentNode.Color
		leftRevolve(n.ParentNode.ParentNode)
		break
	}
	return
}

func insert(v int, r *RBTreeNode) (*RBTreeNode, error) {
	if v > r.Value {
		if r.RightNode == nil {
			r.RightNode = &RBTreeNode{
				Value:      v,
				LeftNode:   nil,
				RightNode:  nil,
				ParentNode: r,
				Color:      red,
				Tp:         "",
			}
			return r.RightNode, nil
		}
		newNode, err := insert(v, r.RightNode)
		return newNode, err
	} else if v < r.Value {
		if r.LeftNode == nil {
			r.LeftNode = &RBTreeNode{
				Value:      v,
				LeftNode:   nil,
				RightNode:  nil,
				ParentNode: r,
				Color:      red,
				Tp:         "",
			}
			return r.LeftNode, nil
		}
		newNode, err := insert(v, r.LeftNode)
		return newNode, err
	} else {
		log.Debug(v, "已存在")
		return nil, nil
	}
}

func deleteNode(node *RBTreeNode, v int) (int, *RBTreeNode) {
	if node == nil {
		return 1, nil
	}
	if v == node.Value {
		switch {
		case node.LeftNode == nil && node.RightNode == nil && node.Color == red:
			// X没有孩子，且X是红色，直接删除X
			realDeleteNode(node)
			return 1, node
		case node.LeftNode == nil && node.RightNode == nil && node.Color == black:
			// X没有孩子，且X是黑色，则以X点进行旋转调色，最后删除X
			balanceDeleteTree(node)
			realDeleteNode(node)
			return 2, node
		case node.LeftNode != nil && node.RightNode != nil:
			// X有两个孩子，从后继中找到最小节点D，交换X和D的数值，再对新X进行删除
			minSuccessor := getMinSuccessor(node.RightNode)
			minSuccessor.Value, node.Value = node.Value, minSuccessor.Value
			deleteNode(minSuccessor, v)
		default:
			// X只有一个孩子C，交换X和C的数值，再对新X进行删除
			if node.LeftNode != nil {
				node.Value, node.LeftNode.Value = node.LeftNode.Value, node.Value
				deleteNode(node.LeftNode, v)
			} else {
				node.Value, node.RightNode.Value = node.RightNode.Value, node.Value
				deleteNode(node.RightNode, v)
			}
		}
	}
	if v > node.Value {
		return deleteNode(node.RightNode, v)
	}
	if v < node.Value {
		return deleteNode(node.LeftNode, v)
	}
	return 1, nil
}

// getMinSuccessor 获取最小后继
func getMinSuccessor(node *RBTreeNode) *RBTreeNode {
	if node.LeftNode == nil {
		return node
	}
	return getMinSuccessor(node.LeftNode)
}

func realDeleteNode(node *RBTreeNode) {
	if node.Tp == "root" {
		*node = RBTreeNode{}
		return
	}

	if node == node.ParentNode.LeftNode {
		node.ParentNode.LeftNode = nil
	}
	if node == node.ParentNode.RightNode {
		node.ParentNode.RightNode = nil
	}
}

func (r *RBTreeNode) isBlack() bool {
	if r == nil {
		return true
	}
	if r.Color == black {
		return true
	}
	return false
}

func balanceDeleteTree(N *RBTreeNode) {
	//N=旋转调色的当前节点，P=N的父亲，W=N的兄弟，Nf=N的远侄子，Nn=N的近侄子
	var P, W, Nf, Nn *RBTreeNode

	// 如果是根节点则没有P
	if N.Tp == "root" {
		N.Color = black
		return
	}

	P = N.ParentNode
	if N == N.ParentNode.LeftNode {
		W = N.ParentNode.RightNode
		Nf = N.ParentNode.RightNode.RightNode
		Nn = N.ParentNode.RightNode.LeftNode
	} else {
		W = N.ParentNode.LeftNode
		Nf = N.ParentNode.LeftNode.LeftNode
		Nn = N.ParentNode.LeftNode.RightNode
	}

	switch {
	case N.Tp == "root" || N.Color == red:
		N.Color = black
	case N.Tp != "root" && N.isBlack() && W.Color == red:
		W.Color = black
		P.Color = red
		if N == P.LeftNode {
			leftRevolve(P)
		} else {
			rightRevolve(P)
		}
		balanceDeleteTree(N)
	case N.Tp != "root" && N.isBlack() && W.isBlack() && Nf.isBlack() && Nn.isBlack():
		W.Color = red
		balanceDeleteTree(P)
	case N.Tp != "root" && N.isBlack() && W.isBlack() && Nf.isBlack() && Nn.Color == red:
		W.Color, Nn.Color = Nn.Color, W.Color
		if N == P.LeftNode {
			rightRevolve(W)
		} else {
			leftRevolve(W)
		}
		balanceDeleteTree(N)
	case N.Tp != "root" && N.isBlack() && W.isBlack() && Nf.Color == red:
		W.Color = P.Color
		P.Color = black
		Nf.Color = black
		if N == P.LeftNode {
			leftRevolve(P)
		} else {
			rightRevolve(P)
		}
	}
}

func rightRevolve(N *RBTreeNode) *RBTreeNode {
	var res = N
	// 将原对象地址表示出来，方便后续建立新的关系
	P := N.ParentNode
	LC := N.LeftNode
	LCRC := N.LeftNode.RightNode
	// 重建三组关系：1.要旋转节点与其父节点关系2.要旋转节点与其左子节点关系3.要旋转节点的左子节点与要旋转节点的左子节点的右子节点的关系
	if N.Tp != "root" {
		N.ParentNode = nil
		if N == P.LeftNode {
			P.LeftNode = nil
			// 新建关系
			P.LeftNode = LC
		} else {
			P.RightNode = nil
			// 新建关系
			P.RightNode = LC
		}
	} else {
		// N是根节点，则这p一段关系不用处理
		N.Tp = ""
		LC.Tp = "root"
		res = LC

	}

	N.LeftNode = nil
	LC.ParentNode = nil

	if LCRC != nil {
		LC.RightNode = nil
		LCRC.ParentNode = nil

		N.LeftNode = LCRC
		LCRC.ParentNode = N
	}

	LC.ParentNode = P
	LC.RightNode = N
	N.ParentNode = LC

	return res
}

func leftRevolve(N *RBTreeNode) *RBTreeNode {
	var res = N
	// 将原对象地址表示出来，方便后续建立新的关系
	P := N.ParentNode
	RC := N.RightNode
	RCLC := N.RightNode.LeftNode
	// 重建三组关系：1.要旋转节点与其父节点关系2.要旋转节点与其左子节点关系3.要旋转节点的左子节点与要旋转节点的左子节点的右子节点的关系
	if N.Tp != "root" {
		N.ParentNode = nil
		if N == P.LeftNode {
			P.LeftNode = nil
			// 新建关系
			P.LeftNode = RC
		} else {
			P.RightNode = nil
			// 新建关系
			P.RightNode = RC
		}
	} else {
		// N是根节点，则这一段关系不用处理
		N.Tp = ""
		RC.Tp = "root"
		res = RC
	}

	N.RightNode = nil
	RC.ParentNode = nil

	if RCLC != nil {
		RC.LeftNode = nil
		RCLC.ParentNode = nil

		N.RightNode = RCLC
		RCLC.ParentNode = N
	}

	RC.ParentNode = P
	RC.LeftNode = N
	N.ParentNode = RC

	return res
}

type drawInfo struct {
	parentPos   int
	currentNode *RBTreeNode
	direction   int // 1左2右
}

func (r *RBTreeNode) Draw() {
	// 高度越高，两个孩子中间距离越大才行
	drawList := []drawInfo{}
	if r.Tp == "root" {
		pos := (2 << uint64(r.GetHeight()-1)) - 1
		fmt.Printf("\x1b[%dC", pos)
		fmt.Printf("\x1b[%dm\x1b[47m%d\x1b[0m\n", black, r.Value)

		if r.LeftNode != nil {
			drawList = append(drawList, drawInfo{parentPos: pos + 1, currentNode: r.LeftNode, direction: 1})
		}
		if r.RightNode != nil {
			drawList = append(drawList, drawInfo{parentPos: pos + 1, currentNode: r.RightNode, direction: 2})
		}
	}
	for {
		tmpDrawList := []drawInfo{}
		fmt.Print("\x1b[s") // 当前光标在第n行的第一列，保存光标位置

		// 求出该层应该位移的度
		maxHeight := 0

		for _, n := range drawList {
			if n.currentNode.GetHeight() > maxHeight {
				maxHeight = n.currentNode.GetHeight()
			}
		}

		//offset := (maxHeight * maxHeight + 3 * maxHeight) / 2
		offset := (2 << uint64(maxHeight-1)) - 1

		for _, n := range drawList {
			pos := n.parentPos
			if n.direction == 1 {
				pos = pos - (offset + 1)
				// 光标移动到父节点位置，然后向左移动两列
				fmt.Printf("\x1b[%dC", n.parentPos)
				fmt.Printf("\x1b[%dD", offset+2)
				fmt.Printf("\x1b[%dm\x1b[47m%d\x1b[0m", n.currentNode.Color, n.currentNode.Value)
				fmt.Print("\x1b[u") // 恢复光标位置 恢复光标和Attrs <ESC> 8
			} else {
				pos = pos + (offset + 1)
				// 光标移动到父节点位置，然后向右移动两列
				fmt.Printf("\x1b[%dC", n.parentPos)
				fmt.Printf("\x1b[%dC", offset)
				fmt.Printf("\x1b[%dm\x1b[47m%d\x1b[0m", n.currentNode.Color, n.currentNode.Value)
				fmt.Print("\x1b[u") // 恢复光标位置 恢复光标和Attrs <ESC> 8
			}

			if n.currentNode.LeftNode != nil {
				tmpDrawList = append(tmpDrawList, drawInfo{parentPos: pos, currentNode: n.currentNode.LeftNode, direction: 1})
			}
			if n.currentNode.RightNode != nil {
				tmpDrawList = append(tmpDrawList, drawInfo{parentPos: pos, currentNode: n.currentNode.RightNode, direction: 2})
			}
		}
		fmt.Print("\n")
		drawList = tmpDrawList
		tmpDrawList = nil
		if len(drawList) == 0 {
			break
		}
	}

}

func (r *RBTreeNode) GetHeight() int {
	leftHeight := 1
	rightHeight := 1
	if r.LeftNode != nil {
		leftHeight += r.LeftNode.GetHeight()
	}
	if r.RightNode != nil {
		rightHeight += r.RightNode.GetHeight()
	}
	if leftHeight > rightHeight {
		return leftHeight
	}
	return rightHeight
}
