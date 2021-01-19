package ipisp

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type mysqlIpIsp struct{}

func (p *mysqlIpIsp) DelOld() {
	_, err := store.DB.Exec("truncate table ipaddress")
	if err != nil {
		log.Errorf("delete old ipisp table error", err.Error())
		return
	}
}

func (p *mysqlIpIsp) AddData(lines []string) error {
	results := make([]string, len(lines))
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

		start, end, _, err := unpack(ipMask)
		if err != nil {
			log.Errorf("invalid ip format: %s, unpack error:%s", ipMask, err.Error())
			return err
		}
		v := fmt.Sprintf("(%d,%d,\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")", start, end, ipMask, ispCode, country, province, path)
		results[i] = v
	}

	// 批量写入可能中间失败，不过目前这个问题可以不作考虑
	values := strings.Join(results, ",")
	query := fmt.Sprintf("insert into ipaddress(`start`, `end`, `ip`, `isp`, `country`, `province`, `path`) values%s", values)
	_, err := store.DB.Exec(query)
	if err != nil {
		log.Errorf(results[0], err.Error())
		return err
	}
	return nil
}

func (p *mysqlIpIsp) GetIspInfo(ip string) (interface{}, error) {
	v, err := IPV42Uint32(ip)
	if err != nil {
		log.Errorf("invalid ip:%s, convert to uint32 error:%s", ip, err.Error())
		return nil, err
	}
	ipIsp := new(Ipaddress)
	exist, err := store.DB.Table(ipIsp).Where("start<=? and end>=?", v, v).Select("isp,country,province,path").Get(ipIsp)
	if err != nil {
		log.Errorf("select ip %s error: %s", ip, err.Error())
		return nil, err
	}
	if !exist {
		log.Warnf("can'tree find ip %s", ip)
		return nil, errors.New("can'tree find ip " + ip)
	}
	return ipIsp, nil
}

type Ipaddress struct {
	Id       int    `xorm:"not null pk autoincr INT(11)" json:"-"`
	Start    int    `xorm:"not null INT(11)" json:"-"`
	End      int    `xorm:"not null INT(11)" json:"-"`
	Ip       string `xorm:"not null VARCHAR(40)" json:"-"`
	Isp      string `xorm:"not null VARCHAR(40)" json:"isp"`
	Province string `xorm:"not null VARCHAR(40)" json:"province"`
	Country  string `xorm:"not null VARCHAR(40)" json:"country"`
	Path     string `xorm:"not null VARCHAR(40)" json:"path"`
}
