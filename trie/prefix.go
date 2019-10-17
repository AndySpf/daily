package main

import (
	"fmt"
	"strings"
)

var root node

//var maxCount int

type node struct {
	childs []*node
	//childHeads string
	value string
	count int
	Tp    string
	Path  string
}

func main() {
	root.Tp = "root"
	root.Path = "|"

	strs := []string{"herld", "hello", "hellr"}
	genTree(strs)

	root.compressTrie()

	fmt.Println(root.search("hello"))

	root.insert("world")
	// 新插入的是未压缩的
	fmt.Println(root.search("world"))
}

func getMaxEqualCount(s1, s2 string) int {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			return i
		}
	}

	if len(s1) > len(s2) {
		return len(s2)
	} else {
		return len(s1)
	}
}

func genTree(strs []string) {
	for _, str := range strs {
		linkChild(&root, str)
	}
}

func linkChild(p *node, str string) {
	maxEqualCount := getMaxEqualCount(p.value, str)
	// 最大相同元素个数, 0代表没有相同元素
	if maxEqualCount > 0 && len(str) == maxEqualCount {
		// 类似父节点为hal 这个字符串本身就是hal，则字符串长度等于索引加1
		return
	}

	// 排除直接找到节点的情况，则需要查看子节点中是否有以str[maxEqualIndex+1]开头的
	for _, item := range p.childs {
		if item.value[0] == str[maxEqualCount] {
			linkChild(item, str[maxEqualCount:])
			return
		}
	}

	// 如果没有子节点以str[maxEqualIndex]开头，则新建子节点
	newNode := &node{
		value: string(str[maxEqualCount]),
		Path:  p.Path + "->" + string(str[maxEqualCount]),
	}
	p.childs = append(p.childs, newNode)
	if len(str) == 1 {
		return
	}
	linkChild(newNode, str[maxEqualCount:])
}

func (p *node) compressTrie() {
	if p.Tp == "root" {
		for _, child := range p.childs {
			child.compressTrie()
		}
		return
	}

	// 将子节点只有一个的合并压缩
	switch len(p.childs) {
	case 0:
		return
	case 1:
		p.value += p.childs[0].value
		// 更新被压缩的那一项的Path
		splitList := strings.Split(p.childs[0].Path, "->")
		p.Path = p.Path + splitList[len(splitList)-1]
		// 更新被压缩的那一项的子节点Path
		for _, i := range p.childs[0].childs {
			i.Path = p.Path + "->" + i.value
		}
		p.childs = p.childs[0].childs
		p.compressTrie()
	default:
		for _, child := range p.childs {
			child.compressTrie()
		}
	}
}

func (p *node) insert(str string) {
	linkChild(p, str)
}

// 查询函数可以使用getMaxEqualCount优化
func (p *node) search(str string) string {
	if p.Tp == "root" {
		for _, child := range p.childs {
			if child.value[0] == str[0] {
				return child.search(str)
			}
		}
		return "尚未加载该元素"
	}

	// 寻找最大相同量
	maxPrefixCount := getMaxEqualCount(p.value, str)

	// 如果maxPrefixCount等于str的长度，则证明改字符串时终止于此节点
	if maxPrefixCount == len(str) {
		return p.Path
	}

	// 如果如果maxPrefixCount小于str的长度，则证明该字符串的一部分属于该节点，剩下的一部分去子节点中匹配
	if maxPrefixCount < len(str) {
		for _, child := range p.childs {
			if string(child.value[0]) == string(str[maxPrefixCount]) {
				return child.search(str[maxPrefixCount:])
			}
		}
		// 子节点中遍历不到
		return "尚未加载该元素"
	}

	// 不存在maxPrefixCount大于str长度的情况
	return "尚未加载该元素"
}
