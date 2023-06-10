package design_pattern

import (
	"fmt"
	"time"
)

type Obj interface {
	Fmt()
}

type Obj1 struct {
	Time   int64
	Length int
}

func (p *Obj1) Fmt() {
	fmt.Println(fmt.Sprintf("obj1 %d", p.Time))
}

type Obj2 struct {
	Time   int64
	Length int
}

func (p *Obj2) Fmt() {
	fmt.Println(fmt.Sprintf("obj2 %d", p.Time))
}

type Factory struct{}

func (p *Factory) CreateObj(index int) Obj {
	switch index {
	case 1:
		return &Obj1{time.Now().Unix(), 16}
	case 2:
		return &Obj2{time.Now().Unix(), 32}
	}
	return nil
}

func GetObj() {
	obj := new(Factory).CreateObj(1)
	obj.Fmt()
}
