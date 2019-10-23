package R_B

import (
	"fmt"
	"testing"
)

func TestRBNode(t *testing.T){
	data := []int{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}
	//data := []int{12, 1, 9, 2, 0, 11, 7, 19, 4}
	fmt.Println("data is ", data)
	root := NewRBTree()
	for _, v := range data{
		root.Insert(v)
		if root.ParentNode != nil{
			// 每次插入后，有可能会更换根节点，可通过判断当前根节点的父节点是否为空进行处理
			root = root.ParentNode
		}
		root.Draw()
	}
	fmt.Println(root.LeftNode)
}

func TestTT(t *testing.T) {
	r := &TT{Name:"zhangsan"}
	fmt.Println("11", &r)
	r.Te()
	fmt.Println("11", &r)
	fmt.Println(r)
	// obj的方法在调用时，方法内obj的指针和调用的指针不相等
}

type TT struct {
	Name string
}

func (t *TT)Te(){
	t1 := &TT{Name:"lisi"}
	fmt.Println("22", &t)
	t = t1
	fmt.Println("22", &t)
}

