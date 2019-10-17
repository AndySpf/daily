package trie

import (
	"fmt"
	"strings"
	"daily/tree"
)

/*
Trie树即前缀树，以树形结构存储字符串的公共前缀，以提高字符串的查找效率。
使用案例：gin的路由框架，输入框提示等

此外还有后缀树，即将一个字符串的所有后缀都参与构建所得到的树。用来解决最长回文字符串等字符串相关问题
（Ukkonen构造后缀树算法时间复杂度为O(n)，尚未理解透彻）
*/

type TrieNode struct {
	Path   string
	Childs []*TrieNode
	Value  string
	Tp     string
}

var root *TrieNode

func Init() {
	root = &TrieNode{
		Path:   "|",
		Childs: nil,
		Value:  "",
		Tp:     "root",
	}
}

func (t *TrieNode) Insert(value interface{}) error {
	linkChild(root, value.(string))
	root.compressTrie()
	return nil
}
func (t *TrieNode) Delete(value interface{}) error {
	// 字典树的目的主要是为了字符串查询效率，一般不会有删除操作
	return nil
}
func (t *TrieNode) Find(value interface{}) (tree.Tree, error) {
	n := root.search(value.(string))
	if n == nil {
		fmt.Println("未加载该元素：", value.(string))
		return nil, tree.ErrNotExist
	}
	return n, nil
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

func linkChild(p *TrieNode, str string) {
	maxEqualCount := getMaxEqualCount(p.Value, str)
	// 最大相同元素个数, 0代表没有相同元素
	if maxEqualCount > 0 && len(str) == maxEqualCount {
		// 类似父节点为hal 这个字符串本身就是hal，则字符串长度等于索引加1
		return
	}

	// 排除直接找到节点的情况，则需要查看子节点中是否有以str[maxEqualIndex+1]开头的
	for _, item := range p.Childs {
		if item.Value[0] == str[maxEqualCount] {
			linkChild(item, str[maxEqualCount:])
			return
		}
	}

	// 如果没有子节点以str[maxEqualIndex]开头，则新建子节点
	newNode := &TrieNode{
		Value: string(str[maxEqualCount]),
		Path:  p.Path + "->" + string(str[maxEqualCount]),
	}
	p.Childs = append(p.Childs, newNode)
	if len(str) == 1 {
		return
	}
	linkChild(newNode, str[maxEqualCount:])
}

func (p *TrieNode) compressTrie() {
	if p.Tp == "root" {
		for _, child := range p.Childs {
			child.compressTrie()
		}
		return
	}

	// 将子节点只有一个的合并压缩
	switch len(p.Childs) {
	case 0:
		return
	case 1:
		p.Value += p.Childs[0].Value
		// 更新被压缩的那一项的Path
		splitList := strings.Split(p.Childs[0].Path, "->")
		p.Path = p.Path + splitList[len(splitList)-1]
		// 更新被压缩的那一项的子节点Path
		for _, i := range p.Childs[0].Childs {
			i.Path = p.Path + "->" + i.Value
		}
		p.Childs = p.Childs[0].Childs
		p.compressTrie()
	default:
		for _, child := range p.Childs {
			child.compressTrie()
		}
	}
}

func (p *TrieNode) search(str string) *TrieNode {
	if p.Tp == "root" {
		for _, child := range p.Childs {
			if child.Value[0] == str[0] {
				return child.search(str)
			}
		}
		return nil
	}

	// 寻找最大相同量
	maxPrefixCount := getMaxEqualCount(p.Value, str)

	// 如果maxPrefixCount等于str的长度，则证明改字符串时终止于此节点
	if maxPrefixCount == len(str) {
		return p
	}

	// 如果如果maxPrefixCount小于str的长度，则证明该字符串的一部分属于该节点，剩下的一部分去子节点中匹配
	if maxPrefixCount < len(str) {
		for _, child := range p.Childs {
			if string(child.Value[0]) == string(str[maxPrefixCount]) {
				return child.search(str[maxPrefixCount:])
			}
		}
		// 子节点中遍历不到
		return nil
	}

	// 不存在maxPrefixCount大于str长度的情况
	return nil
}

func (p *TrieNode) GetValue() interface{} {
	return p.Path
}
