package ipisp

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type ipIspItem struct {
	CIDR     string `json:"CIDR"`
	Isp      string `json:"isp"`
	Province string `json:"province"`
	Country  string `json:"country"`
	Path     string `json:"path"`
}

var ipIspArray []binaryItem

type binaryItem struct {
	StartIp uint32
	EndIp   uint32
	Value   ipIspItem
}

type binaryIpIsp struct{}

func (b binaryIpIsp) AddData(lines []string) error {
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		l := strings.Split(lines[i], string([]byte{9})) // txt文件以水平定位符作为分隔符了，ascii码十进制表示为9
		ipMask := l[0]
		country := l[1]
		province := l[2]
		isp := strings.Replace(l[3], "\n", "", -1)
		ispCode := getIspCode(isp)
		path := getPath(province)

		// 构建二分查找切片
		start, end, _, err := unpack(ipMask)
		if err != nil {
			log.Errorf("invalid ip:%s", ipMask)
			continue
		}

		ipIspArray = append(ipIspArray, binaryItem{
			StartIp: start,
			EndIp:   end,
			Value: ipIspItem{
				CIDR:     ipMask,
				Isp:      ispCode,
				Province: province,
				Country:  country,
				Path:     path,
			},
		})
	}
	return nil
}

func (b binaryIpIsp) GetIspInfo(ip string) (interface{}, error) {
	v, err := IPV42Uint32(ip)
	if err != nil {
		log.Errorf("invalid ip:%s, convert to uint32 error:%s", ip, err.Error())
		return nil, err
	}

	left := 0
	right := len(ipIspArray) - 1
	for {
		mid := (right-left)/2 + left
		if isMatch(v, left) {
			return ipIspArray[left].Value, nil
		}
		if isMatch(v, right) {
			return ipIspArray[right].Value, nil
		}
		if isMatch(v, mid) {
			return ipIspArray[mid].Value, nil
		}
		if v > ipIspArray[mid].EndIp {
			left = mid + 1
			continue
		}
		if v < ipIspArray[mid].StartIp {
			right = mid - 1
			continue
		}
		if left >= right {
			return nil, errors.New("not find ip:" + ip)
		}
	}
}

func (b binaryIpIsp) DelOld() {
	ipIspArray = ipIspArray[:0]
}

func isMatch(ip uint32, index int) bool {
	if ip >= ipIspArray[index].StartIp && ip <= ipIspArray[index].EndIp {
		return true
	}
	return false
}

func unpack(ipMask string) (uint32, uint32, int, error) {
	l := strings.Split(ipMask, "/")
	if len(l) < 2 {
		ip, err := IPV42Uint32(ipMask)
		if err != nil {
			return 0, 0, 0, err
		}
		return ip, ip, 32, nil
	}
	ip := l[0]

	mask, err := strconv.Atoi(l[1])
	if err != nil {
		return 0, 0, 0, err
	}

	start, err := IPV42Uint32(ip)
	if err != nil {
		return 0, 0, 0, err
	}
	if mask == 32 {
		return start, start, 32, nil
	}
	end := uint32(start | (1<<(32-mask) - 1))
	return start, end, mask, nil
}

func IPV42Uint32(ip string) (uint32, error) {
	sList := strings.Split(ip, ".")
	if len(sList) < 4 {
		return 0, errors.New("invalid ip")
	}
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
