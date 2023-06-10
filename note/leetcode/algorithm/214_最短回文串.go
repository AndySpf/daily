package algorithm

import (
	"fmt"
	"strings"
	"unsafe"
)

//.*aacecaaa    aacecaaa   0:2 3:5; 0:3 4:7   #a#a#c#e#c#a#a#a#
// aaacecaaa
func shortestPalindrome(s string) string {
	if s == "" {
		return s
	}
	// s当中以第一个字符开头的最大回文子串的尾索引i，然后将s[i+1:]翻转添加到首部
	newS := fmt.Sprintf("#%s#", strings.Join(strings.Split(s, ""), "#"))
	maxIndex := 0
	var index int
	for index = 1; index <= (len(newS)-1)/2; index++ {
		if newS[0:index] == bytes2string(reverseSlice(string2bytes(newS[index+1:2*index+1]))) {
			if index > maxIndex {
				maxIndex = index
			}
		}
	}
	res := []byte{}
	for i := maxIndex + maxIndex + 1; i < len(newS); i += 2 {
		res = append(res, newS[i])
	}
	s = bytes2string(reverseSlice(res)) + s
	return s
}

func reverseSlice(a []byte) []byte {
	tmp := make([]byte, len(a))
	copy(tmp, a) // 终究要复制的，就没必要无拷贝进行bytes和string转换了
	for i := len(tmp)/2 - 1; i >= 0; i-- {
		opp := len(tmp) - 1 - i
		tmp[i], tmp[opp] = tmp[opp], tmp[i]
	}

	return tmp
}

func bytes2string(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func string2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
