package design_pattern

import (
	"fmt"
	"time"
)

type Product interface {
	FmtMe()
}

type Product1 struct {
	Time int64
	Name string
}

func (p *Product1) Create() Product {
	return &Product1{Time: time.Now().Unix(), Name: "product1"}
}

func (p *Product1) FmtMe() {
	fmt.Println(fmt.Sprintf("%s %d", p.Name, p.Time))
}

type Product2 struct {
	Time int64
	Name string
}

func (p *Product2) Create() Product {
	return &Product2{Time: time.Now().Unix(), Name: "product2"}
}

func (p *Product2) FmtMe() {
	fmt.Println(fmt.Sprintf("%s %d", p.Name, p.Time))
}

type AbstractFactory interface {
	Create() Product
}

func GetProduct() {
	new(Product1).Create().FmtMe()
	new(Product2).Create().FmtMe()
}
