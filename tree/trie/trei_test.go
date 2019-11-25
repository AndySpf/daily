package trie

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestTrieNode_Insert(t *testing.T) {
	Init()
	root.Insert("hello")
	fmt.Println(*root.Childs[0])
	root.Find("hello")
}

type Widget struct {
	id    int
	attrs []string
}

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

type Server struct {
	ServerName string `key1:"value1" key11:"value11"`
	ServerIP   string `key2:"value2"`
}

const (
	mutexLocked = iota // mutex is locked
	mutexWoken
	mutexStarving    = 4
	mutexWaiterShift = iota
	starvationThresholdNs
)

func TestInit(t *testing.T) {
	a := "xxxxxxxxxx"

	fmt.Println(a)
	x := (*[2]uintptr)(unsafe.Pointer(&a))
	h := [3]uintptr{x[0], x[1], x[1]}
	bs := *(*[]byte)(unsafe.Pointer(&h))
	//strHeader := *(*[2]uintptr)(unsafe.Pointer(&a))
	//
	//sliceHeader := [3]uintptr{
	//	strHeader[0],
	//	strHeader[1],
	//	strHeader[1],
	//}
	//
	//b := *(*[]byte)(unsafe.Pointer(&sliceHeader))
	fmt.Println(bs)

	bss := []byte(a)
	bss[1] = 'a'
	fmt.Println(a)

}

func doSomThing(ss *[]Widget) {
	data := *ss
	data[0].id = 1
}

func doOtherThing(sp []*Widget) {
	fmt.Printf("s2%p\n", sp)
	sp = nil
	fmt.Printf("s2%p\n", sp)
}
