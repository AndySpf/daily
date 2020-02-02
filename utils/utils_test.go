package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestExecuteCMD(t *testing.T) {
	s := "./"
	ss, _ := filepath.Abs(s)
	res, _ := ExecuteCMD(fmt.Sprintf("ls -lt %s | awk '{print $9}'", ss))
	res1 := strings.Trim(res, "\n")
	fmt.Println(res1)
}

func TestIndex2Col(t *testing.T) {
	fmt.Println(Index2Col(0))
}

type ss struct {
	Name *string
}

func TestT(t *testing.T) {
	//fmt.Println(T())
	n1 := "zs"
	n2 := "ls"
	a := &ss{Name: &n1}
	b := &ss{Name: &n2}
	fmt.Println(a.Name)
	fmt.Println(b.Name)
	a.Name = b.Name
	fmt.Println(a.Name)
	fmt.Println(b.Name)
	fmt.Println(*a.Name)
}

func TestTs(t *testing.T) {
	format()
}

func TestRemoveItemFromSlice(t *testing.T) {
	m1 := map[string]string{"name": "zs", "age": "1"}
	m2 := make(map[string]string)
	for key, value := range m1 {
		m2[key] = value
	}
	fmt.Println(&m2)
	fmt.Println(&m1)
	m1["name"] = "ls"
	fmt.Println(m1)
	fmt.Println(m2)

	slice := []int{1, 2, 3, 4}
	i := 3
	copy(slice[i:], slice[i+1:])
	slice = slice[: len(slice)-1 : len(slice)-1]
	//copy(slice[i:], slice[i+1:])
	//slice = slice[: len(slice)-1 : len(slice)-1]
	//copy(slice[i:], slice[i+1:])
	//slice = slice[: len(slice)-1 : len(slice)-1]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
