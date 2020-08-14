package main

import (
	"fmt"
	"testing"
)

func TestUint32ToBinary(t *testing.T) {
	uint32ToBinary()
}

func TestGetUInt32ByIPNet(t *testing.T) {
	v, mask, err := getUInt32ByIPNet("192.168.29.0/24")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(v)
	fmt.Println(mask)
}

func Benchmark_Tree(b *testing.B) {
	b.StopTimer()
	LoadData()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v, _ := getUInt32ByIP("134.175.68.57")
		getIspByIP(v)
		//root.Find(v)
	}
}
