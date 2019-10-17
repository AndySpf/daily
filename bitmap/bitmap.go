package main

import "fmt"

// bitMap排序

type MapBit struct {
	bs   []byte
	data []int
}

func main() {
	m := new(MapBit)
	m.data = []int{15, 3, 8, 5, 2, 7, 9}
	m.bs = []byte{}
	m.setBit()

	fmt.Println(m.getBit())
}

func (m *MapBit) setBit() byte {
	var b byte
	for _, i := range m.data {
		index, offset := i/8, 7-i%8
		for index >= len(m.bs) {
			m.bs = append(m.bs, byte(0))
		}
		m.bs[index] |= uint8(1) << uint8(offset)
	}
	return b
}

func (m *MapBit) getBit() []int {
	res := []int{}
	notExist := []int{}

	for index := range m.bs {
		for i := 0; i < 8; i++ {
			v := (m.bs[index] >> uint64(7-i)) & uint8(1)
			if v == 1 {
				res = append(res, i+(index<<3))
			} else {
				notExist = append(notExist, i+(index<<3))
			}
		}
	}
	fmt.Println("前16位数中没有出现的数有", notExist)
	return res
}
