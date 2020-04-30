package utils

import (
	"bytes"
	"container/heap"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	"io/ioutil"
	"os"
	"testing"
)

type S struct {
	Name string
}

func TestExecuteCMD(t *testing.T) {
	m := map[string]*S{}

	m["zs"] = &S{
		Name: "zs",
	}
	m["ls"] = &S{
		Name: "ls",
	}

	fmt.Println(m["zs"])
	delete(m, "ww")
	fmt.Println(m["zs"])
	fmt.Println(m["ww"])
	fmt.Println(^uint16(0))
	fmt.Println(uint8(255))
}

func TestFileExist(t *testing.T) {
	fmt.Println(FileExist("/Users/qijing.fqj/go/src/reverse_tunnel/utils/utils1.go"))
}

var m = map[string]chan bool{}

func TestStr2bytes(t *testing.T) {
	fmt.Println(m["a"])
	m["a"] = make(chan bool)
	go tt()
	fmt.Println(m["a"])

	fmt.Println("ccccc")
	m["a"] <- true
	fmt.Println("bbbbb")
}
func tt() {
	<-m["a"]
	fmt.Println("aaaa")
}

func TestAppendFile(t *testing.T) {
	f, err := os.Create("./test1")
	if err != nil {
		t.Fatal(err.Error())
	}
	f.Close()
	fmt.Println([]byte("\n"))
	if err := AppendFile("./test1", []byte("hello")); err != nil {
		fmt.Println(err.Error())
	}
	if err := AppendFile("./test1", []byte("world")); err != nil {
		fmt.Println(err.Error())
	}
	bs, _ := ioutil.ReadFile("./test")
	fmt.Println(string(bs))

	os.Remove("./test1")
}

//type CustomChangeListener struct {
//
//}
//func (c *CustomChangeListener) OnChange(changeEvent *agollo.ChangeEvent) {
//	// TODO Reload()
//	fmt.Println(changeEvent.Namespace)
//	for k, v := range changeEvent.Changes{
//		fmt.Printf(k, ":")
//		fmt.Println(*v)
//	}
//}
//func TestDaily(t *testing.T) {
//	if err := agollo.Start();err != nil{
//		panic(err.Error())
//	}
//
//	agollo.AddChangeListener(&CustomChangeListener{})
//	select{
//
//	}
//}

type INT1 struct {
	Name string `json:"name"`
}

type SS struct {
	Num int `json:"name"`
}

type SSS struct {
	O    *SS
	Name string
}

func TestDaily2(t *testing.T) {
	m := map[string]string{
		"zs": "1",
		"ls": "2",
	}
	s := map[string]*string{}
	for k, v := range m {
		s[k] = &v
	}
	fmt.Println(*s["zs"])
	fmt.Println(*s["ls"])
}

type ts interface {
	name()
}

func remove(slice []int, i int) []int {
	b := slice[:1:2]
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Printf("slice is %p, b is %p\n", slice, b)
	b = append(b, 3)
	fmt.Printf("slice is %p, b is %p\n", slice, b)
	fmt.Printf("slice is %v, b is %v\n", slice, b)
	copy(slice[2:], slice[3:])
	return slice[:len(slice)-1]
}

type StrTest struct {
	Name string `json:"name"`
}

func (p *StrTest) String() string {
	return "aaaaa"
}

type s struct {
	Name time.Time `json:"name"`
}

func (p s) String() string {
	return "aaaa"
}

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

type augmentedReader struct {
	innerReader io.Reader
	augmentFunc func([]byte) []byte
}

// replaces ' ' with '!'
func bangify(buf []byte) []byte {
	return bytes.Replace(buf, []byte(" "), []byte("!"), -1)
}

func (r *augmentedReader) Read(buf []byte) (int, error) {
	tmpBuf := make([]byte, len(buf))
	n, err := r.innerReader.Read(tmpBuf)
	copy(buf[:n], r.augmentFunc(tmpBuf[:n]))
	fmt.Println(n)
	return n, err
}

func BangReader(r io.Reader) io.Reader {
	fmt.Println("Bang")
	return &augmentedReader{innerReader: r, augmentFunc: bangify}
}

func UpcaseReader(r io.Reader) io.Reader {
	fmt.Println("Upcase")
	return &augmentedReader{innerReader: r, augmentFunc: bytes.ToUpper}
}

type Item struct {
	value    string // 优先级队列中的数据，可以是任意类型，这里使用string
	priority int    // 优先级队列中节点的优先级
	index    int    // index是该节点在堆中的位置
}

// 优先级队列需要实现heap的interface
type PriorityQueue []*Item

// 绑定Len方法
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 绑定Less方法，这里用的是小于号，生成的是小根堆
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// 绑定swap方法
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

// 绑定put方法，将index置为-1是为了标识该数据已经出了优先级队列了
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	//item.index = -1
	return item
}

// 绑定push方法
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// 更新修改了优先级和值的item在优先级队列中的位置
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type t1 struct {
	A int `default:"42"`
}

func deadloop() {
	for {
		fmt.Println("deadloop")
		time.Sleep(time.Second * 1)
	}
}

func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func TestDaily1(t *testing.T) {
	A := []int{1, 2, 3, 0, 0, 0}
	B := []int{2, 5, 6}
	merge(A, 3, B, 3)
	fmt.Println(A)
}

func merge(A []int, m int, B []int, n int) {
	aIndex, bIndex := 0, 0
	for {
		if A[aIndex] > B[bIndex] {
			copy(A[aIndex+1:], A[aIndex:])
			A[aIndex] = B[bIndex]
			bIndex++
			aIndex++
		} else {
			aIndex++
		}

		if aIndex == m+bIndex {
			copy(A[aIndex:], B[bIndex:])
			bIndex = n
		}
		if bIndex >= n {
			break
		}
	}
}
