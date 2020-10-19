package algorithm

//     1
//  2      3
//4   5  6   7

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectLevelOrder(root)
	return root
}

func connectLevelOrder(node *Node) {
	queue := []*Node{node}
	for len(queue) != 0 {
		cp := make([]*Node, len(queue))
		copy(cp, queue)
		queue = nil
		for i := range cp {
			if i > 0 {
				if i == len(cp)-1 {
					cp[i].Next = nil
				}
				cp[i-1].Next = cp[i]
			}

			if cp[i].Left != nil {
				queue = append(queue, cp[i].Left)
			}
			if cp[i].Right != nil {
				queue = append(queue, cp[i].Right)
			}
		}
	}
}

func connect1(root *Node) *Node {
	connectDepthOrder(root, nil)
	return root
}

func connectDepthOrder(cur, next *Node) {
	if cur == nil {
		return
	}
	cur.Next = next
	connectDepthOrder(cur.Left, cur.Right)
	if cur.Next == nil {
		connectDepthOrder(cur.Right, nil)
	} else {
		connectDepthOrder(cur.Right, cur.Next.Left)
	}
}
