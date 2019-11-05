package tree

import "errors"

/*
二叉树：
完全二叉树：除倒数第二行外均有两个子节点，且最后一行拥有两个子树的节点靠左紧凑
满二叉树： 除最后一行的叶子节点外所有结点都有两个子节点
二叉查找树: 左子节点永远比根节点小，右子节点永远比根节点大
AVL树：(强平衡查找树)
红黑树：(弱平衡查找树) https://my.oschina.net/u/3272058/blog/1914452
B树
B+树
B*树
Trie树 字典树，前缀树。每个节点是字符串的最大公共前缀，如gin的路由树。
*/
var ErrNotExist = errors.New("node not exist")

type Tree interface {
	Insert(value interface{}) (bool, error)
	Delete(value interface{}) error
	Find(value interface{}) (Tree, error)
	GetValue() interface{}
}

type DrawTree struct {
	LeftNode  *Tree
	RightNode *Tree
	Height    int
	interval  int
}
