package R_B

import (
	"fmt"
	"testing"
)

func TestRBNode(t *testing.T){
	data := []int{12, 1, 9}
	root := NewRBTree()
	for _, v := range data{
		root.Insert(v)
	}
	fmt.Println(root)
	root.Draw()
}

