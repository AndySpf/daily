package design_pattern

import "fmt"

type builder struct {
	Builder1 b1
	Builder2 b2
}

type b1 struct {
	Item1 string
}

type b2 struct {
	Item3 int
}

func (p *builder) WithB1() *builder {
	p.Builder1 = b1{}
	return p
}
func (p *builder) WithB2() *builder {
	p.Builder2 = b2{}
	return p
}
func (p *builder) WithItem1(s string) *builder {
	p.Builder1.Item1 = s
	return p
}
func (p *builder) WithItem3(i int) *builder {
	p.Builder2.Item3 = i
	return p
}

func getBuilder() {
	b := &builder{}
	b.WithB1().WithItem1("hello").WithB2().WithItem3(123)
	fmt.Println(b)
}
