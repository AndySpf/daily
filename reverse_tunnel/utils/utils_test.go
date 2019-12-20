package utils

import (
	"fmt"
	"time"

	//"github.com/zouyx/agollo"
	//"github.com/zouyx/agollo"
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

func TestDaily1(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = dp[i-1]
		if dp[i-2]+nums[i-1] > dp[i-1] {
			dp[i] = dp[i-2] + nums[i-1]
		}
	}
	fmt.Println(dp)
}

func f() {
	defer func() {
		fmt.Println("f")
	}()
}
func f1() {
	defer func() {
		fmt.Println("f1")
	}()
}
