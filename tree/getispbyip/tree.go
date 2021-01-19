package ipisp

import (
	"strings"
)

type tree struct {
	Bit   uint8
	Value ipIspItem
	LNote *tree
	RNote *tree
}

var root = new(tree)

type treeIpIsp struct{}

func (t treeIpIsp) AddData(lines []string) error {
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		l := strings.Split(lines[i], string([]byte{9})) // txt文件以水平定位符作为分隔符了，ascii码十进制表示为9
		ipMask := l[0]
		country := l[1]
		province := l[2]
		isp := strings.Replace(l[3], "\n", "", -1)
		ispCode := getIspCode(isp)
		path := getPath(province)

		// 按bit位构建二叉树
		v, _, mask, _ := unpack(ipMask)
		root.Insert(v, ipIspItem{
			CIDR:     ipMask,
			Isp:      ispCode,
			Province: province,
			Country:  country,
			Path:     path,
		}, mask)
	}
	return nil
}

func (t treeIpIsp) GetIspInfo(ip string) (interface{}, error) {
	v, err := IPV42Uint32(ip)
	if err != nil {
		return nil, err
	}
	return root.Find(v), nil
}

func (t treeIpIsp) DelOld() {
	root.RNote = nil
	root.LNote = nil
	return
}

func (p *tree) Insert(ip uint32, value ipIspItem, mask int) {
	cur := p
	for i := 1; i <= mask; i++ {
		bit := uint8((ip & (1 << (32 - i))) >> (32 - i))
		if cur.LNote != nil && cur.LNote.Bit == bit {
			cur = cur.LNote
			continue
		}
		if cur.RNote != nil && cur.RNote.Bit == bit {
			cur = cur.RNote
			continue
		}

		cnode := &tree{
			Bit:   bit,
			Value: ipIspItem{},
			LNote: nil,
			RNote: nil,
		}
		if i == mask {
			cnode.Value = value
		}

		if bit == 0 {
			cur.LNote = cnode
			cur = cur.LNote
		} else {
			cur.RNote = cnode
			cur = cur.RNote
		}
	}
}

func (p *tree) Find(v uint32) interface{} {
	cur := p
	for i := 1; i <= 32; i++ {
		if cur.LNote == nil && cur.RNote == nil {
			return cur.Value
		}
		bit := uint8((v & (1 << (32 - i))) >> (32 - i))
		if cur.LNote != nil && cur.LNote.Bit == bit {
			cur = cur.LNote
			continue
		}
		if cur.RNote != nil && cur.RNote.Bit == bit {
			cur = cur.RNote
			continue
		}

		// if
		//    node
		//   +    +
		//  0
		// but bit==1, This data is not entered，can't find this ip, break
		break
	}
	return "error"
}
