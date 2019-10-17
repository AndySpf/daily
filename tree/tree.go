package tree

import "errors"

/*
二叉树：
完全二叉树：除倒数第二行外均有两个子节点，且最后一行拥有两个子树的节点靠左紧凑
满二叉树： 除最后一行的叶子节点外所有结点都有两个子节点
二叉查找树
AVL树（平衡查找树）
红黑树：
B树
B+树
B*树
Trie树
*/
var ErrNotExist = errors.New("node not exist")

type Tree interface {
	Insert(value interface{}) error
	Delete(value interface{}) error
	Find(value interface{}) (Tree, error)
	GetValue() interface{}
}
