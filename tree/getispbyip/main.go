// 测试根据ip查找对应ip的信息两种算法性能
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var allIsp = make([]ispItem, 0, 1<<19)
var root = new(t)

type ispItem struct {
	StartIp uint32
	EndIp   uint32
	ISP     string
}

type t struct {
	Value uint8
	ISP   string
	LNote *t
	RNote *t
}

type isp struct {
	v   uint32
	isp string
}

// 110000001010100000011100 00001100
// 110000001010100000011101 00001100
func main() {
	LoadData()

	http.HandleFunc("/getisp/tree", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		fmt.Println(ip)
		v, _ := getUInt32ByIP(ip)
		start := time.Now().UnixNano()
		isp := root.Find(v)
		fmt.Println(time.Now().UnixNano() - start)
		w.Write([]byte(isp))
	})

	// 45M左右   采用二叉树内存约为二分法的二倍，查找时间几乎相同
	http.HandleFunc("/getisp/binary_search", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")
		fmt.Println(ip)
		v, _ := getUInt32ByIP(ip)
		start := time.Now().UnixNano()
		isp := getIspByIP(v)
		fmt.Println(time.Now().UnixNano() - start)
		w.Write([]byte(isp))
	})
	http.HandleFunc("/getisp/binary_search_update", func(w http.ResponseWriter, r *http.Request) {
		allIsp = make([]ispItem, 0, 1<<19)
		LoadData()
	})

	http.ListenAndServe("127.0.0.1:19999", nil)

}

func LoadData() {
	file := "/Users/spf/go/src/daily/tree/getispbyip/ipisp.txt"
	f, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		panic(err.Error())
		return
	}
	defer f.Close()

	start := time.Now().Unix()
	buf := bufio.NewReader(f)
	lines := make([]string, 1000)
	i := 0
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok! use ", time.Now().Unix()-start)
				addData(lines[:i+1])
				break
			} else {
				fmt.Println("read isp file error: ", err.Error())
				return
			}
		}
		lines[i] = line
		i++
		if i%1000 == 0 {
			addData(lines)
			i = 0
		}
	}
}

func addData(lines []string) {
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		l := strings.Split(lines[i], string([]byte{9})) // txt文件以水平定位符作为分隔符了，ascii码十进制表示为9
		ipMask := l[0]
		//country := l[1]
		province := l[2]
		ispStr := strings.Replace(l[3], "\n", "", -1)

		// 按bit位构建二叉树
		//v, mask, _ := getUInt32ByIPNet(ipMask)
		//root.Insert(isp{
		//	v:   v,
		//	isp: province + "|" + ispStr,
		//}, mask)

		// 构建二分查找切片
		start, end, _ := unpack(ipMask)
		allIsp = append(allIsp, ispItem{
			StartIp: start,
			EndIp:   end,
			ISP:     province + "|" + ispStr,
		})
	}
}

func (p *t) Insert(isp isp, mask int) {
	cur := p
	for i := 1; i <= mask; i++ {
		bit := uint8((isp.v & (1 << (32 - i))) >> (32 - i))
		if cur.LNote != nil && cur.LNote.Value == bit {
			cur = cur.LNote
			continue
		}
		if cur.RNote != nil && cur.RNote.Value == bit {
			cur = cur.RNote
			continue
		}

		cnode := &t{
			Value: bit,
			ISP:   "",
			LNote: nil,
			RNote: nil,
		}
		if i == mask {
			cnode.ISP = isp.isp
		}

		if bit == 0 {
			cur.LNote = cnode
			cur = cur.LNote
		} else {
			cur.RNote = cnode
			cur = cur.RNote
		}
	}
}

func (p *t) Find(v uint32) string {
	cur := p
	for i := 1; i <= 32; i++ {
		if cur.LNote == nil && cur.RNote == nil {
			return cur.ISP
		}
		bit := uint8((v & (1 << (32 - i))) >> (32 - i))
		if cur.LNote != nil && cur.LNote.Value == bit {
			cur = cur.LNote
			continue
		}
		if cur.RNote != nil && cur.RNote.Value == bit {
			cur = cur.RNote
			continue
		}

		// 没有走到树的下一层。证明左节点或者右节点有值，但不是该ip对应位的值，说明没有录入该ip段
		break
	}
	return "error"
}

func getUInt32ByIP(ip string) (uint32, error) {
	sList := strings.Split(ip, ".")
	iList := make([]int, len(sList))
	for i := range sList {
		var err error
		iList[i], err = strconv.Atoi(sList[i])
		if err != nil {
			return 0, err
		}
		if iList[i]>>8 > 0 || iList[i] < 0 { // bigger than 255 or smaller than 0
			return 0, errors.New("invalid ip " + ip)
		}
	}

	v := uint32(iList[0]<<24 + iList[1]<<16 + iList[2]<<8 + iList[3])
	return v, nil
}

func getUInt32ByIPNet(ipNet string) (uint32, int, error) {
	l := strings.Split(ipNet, "/")
	if len(l) < 2 {

	}
	mask, err := strconv.Atoi(l[1])
	if err != nil {

	}
	if mask > 32 {

	}
	sList := strings.Split(l[0], ".")
	iList := make([]int, len(sList))
	for i := range sList {
		var err error
		iList[i], err = strconv.Atoi(sList[i])
		if err != nil {
			return 0, 0, err
		}
		if iList[i]>>8 > 0 || iList[i] < 0 { // bigger than 255 or smaller than 0
			return 0, 0, errors.New("invalid ip " + ipNet)
		}
	}

	v := uint32(iList[0]<<24 + iList[1]<<16 + iList[2]<<8 + iList[3])
	return v, mask, nil
}

func uint32ToBinary() {
	v := uint8(100) // 0110 0100
	for i := 1; i <= 8; i++ {
		bit := (v & (1 << (8 - i))) >> (8 - i) // 1000 0000
		fmt.Println(bit)
	}
}

func IpIspHandler1() {
	ip := "192.168.29.89"
	v, err := getUInt32ByIP(ip)
	if err != nil {
		return
	}
	fmt.Println(v)

	getIspByIP(v)
}

func getIspByIP(ip uint32) string {
	left := 0
	right := len(allIsp) - 1
	for {
		mid := (right-left)/2 + left
		if isMatch(ip, left) {
			return allIsp[left].ISP
		}
		if isMatch(ip, right) {
			return allIsp[right].ISP
		}
		if isMatch(ip, mid) {
			return allIsp[mid].ISP
		}
		if ip > allIsp[mid].EndIp {
			left = mid + 1
			continue
		}
		if ip < allIsp[mid].StartIp {
			right = mid - 1
			continue
		}
	}
}

func isMatch(ip uint32, index int) bool {
	if ip >= allIsp[index].StartIp && ip <= allIsp[index].EndIp {
		return true
	}
	return false
}

func unpack(ipMask string) (uint32, uint32, error) {
	l := strings.Split(ipMask, "/")
	if len(l) < 2 {
		ip, err := getUInt32ByIP(ipMask)
		if err != nil {
			return 0, 0, err
		}
		return ip, ip, nil
	}
	ip := l[0]

	mask, err := strconv.Atoi(l[1])
	if err != nil {
		return 0, 0, err
	}

	start, err := getUInt32ByIP(ip)
	if err != nil {
		return 0, 0, err
	}
	if mask == 32 {
		return start, start, nil
	}
	end := uint32(start | (1<<(32-mask) - 1))
	return start, end, nil
}
