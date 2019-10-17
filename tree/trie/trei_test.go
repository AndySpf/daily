package trie

import (
	"fmt"
	"testing"
)

func TestTrieNode_Insert(t *testing.T) {
	Init()
	root.Insert("hello")
	fmt.Println(*root.Childs[0])
	root.Find("hello")
}
