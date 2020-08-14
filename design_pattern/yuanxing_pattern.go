package design_pattern

import "fmt"

type CopyObj interface {
	clone() CopyObj
}

type Struct1 struct {
	Name string
}

func (p *Struct1) clone() CopyObj {
	cp := *p
	return &cp
}

func getCopyObj() {
	obj := Struct1{
		Name: "struct1",
	}
	copyObj := obj.clone()
	fmt.Println(copyObj.(*Struct1).Name)
}
