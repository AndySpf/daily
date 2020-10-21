package algorithm

import (
	"fmt"
)

// IsMatch p:abf*b
// 		   s:abfff
func IsMatch(s, p string) bool {
	if p[0] == '*' {
		return false
	}

	status := make([][]bool, len(s)+1) // 二维数组，status[i,j] s的前i项可以被p的前j项匹配
	for i := range status {
		status[i] = make([]bool, len(p)+1)
	}
	// 初始化p或者s为空的情况：
	// p为空的情况一定为false，即二维数组中，每一个小数组的第一项为false
	// s，p都为空的情况认为是能够匹配的，为true
	// 将s为空，p不为空的情况待确定（存在p如c*的情况，这种情况就算s为空也是能够匹配的）
	/*
		        true   x   x   x   x   x待初始化
				false  Y   Y   Y   Y   Y自底向上待计算
				false  Y   Y   Y   Y
				false  Y   Y   Y   Y
	*/
	status[0][0] = true
	for j := 0; j < len(p); j++ { //排除类似c*等一开始就匹配不到的内容,这样就可以认为c为0次
		if p[j] == '*' && status[0][j-1] {
			status[0][j+1] = true
		}
	}

	// 开始计算
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '.' || s[i] == p[j] {
				status[i+1][j+1] = status[i][j]
			}
			if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' { // 匹配不到前一个
					status[i+1][j+1] = status[i+1][j-1]
				} else {
					status[i+1][j+1] = status[i+1][j] || status[i][j+1] || status[i+1][j-1]
				}
			}
		}
	}
	return status[len(s)][len(p)]
}

// IsMatch1 按照下一跳路径，递归
func IsMatch1(s string, p string) bool {
	begin := new(regexNode)
	begin.C = '>'
	generatePattern(begin, p, 0)
	fmt.Println(begin)
	return check(begin, s, 0)
}

type regexNode struct {
	C        byte
	Children map[byte][]*regexNode
	End      bool
}

func (n *regexNode) append(c byte, child *regexNode) {
	m := n.Children
	if m == nil {
		m = make(map[byte][]*regexNode)
		n.Children = m
	}
	list := m[c]
	if list == nil {
		list = make([]*regexNode, 0)
	}
	for _, v := range list {
		if v == child {
			m[c] = list
			return
		}
	}
	list = append(list, child)
	m[c] = list
}

func generatePattern(now *regexNode, str string, idx int) {
	if len(str) <= idx {
		now.End = true
		return
	}
	vnow := now
	switch str[idx] {
	case '*':
		now.append(now.C, now)
	default:
		node := new(regexNode)
		node.C = str[idx]
		now.append(str[idx], node)
		vnow = node
	}
	generatePattern(vnow, str, idx+1)
	return
}

func check(now *regexNode, str string, idx int) bool {
	if len(str) <= idx {
		return now.End
	}
	list := now.Children['.']
	for _, v := range now.Children[str[idx]] {
		list = append(list, v)
	}
	for _, v := range list {
		r := check(v, str, idx+1)
		if r {
			return true
		}
	}
	return false
}
