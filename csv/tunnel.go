package main

import (
	"encoding/csv"
	"fmt"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
	"os"
	"strings"
)

var AllRelation = ResData{}

func main() {
	f, err := os.OpenFile("/Users/qijing.fqj/go/src/daily/csv/pathinfo.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://139.196.142.168:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	var pathInfos = map[string][]string{}
	//q := client.NewQuery(`select miss, pathname from "1w".pathinfo where tid='2846' group by "pathname" limit 3`, "steplog", "")
	q := client.NewQuery(`show tag values on steplog from "1w".pathinfo with key = tid`, "steplog", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		//fmt.Printf("%+v", response.Results)
		length := len(response.Results[0].Series[0].Values)
		fmt.Println(length)
		for index, tidInfo := range response.Results[0].Series[0].Values {
			tid := tidInfo[1].(string)
			fmt.Println(index)
			q := client.NewQuery(fmt.Sprintf(`select miss,pathname from "1w".pathinfo where tid='%s' and pathname =~ /^->/ group by pathname`, tid), "steplog", "")
			if res, err := c.Query(q); err == nil && response.Error() == nil {
				//fmt.Println(res.Results[0].Series)
				if len(res.Results[0].Series) != 0 {
					// 分组后res.Results[0].Series会有多个，取每个中的第一个即可
					for _, series := range res.Results[0].Series {
						pathInfo := series.Values[0][2].(string)
						if !isFiltered(pathInfo) {
							if _, ok := pathInfos[tid]; ok {
								pathInfos[tid] = append(pathInfos[tid], pathInfo)
							} else {
								pathInfos[tid] = []string{pathInfo}
							}
						}
					}
					//break
				}

			}
		}
		genAllRelations(pathInfos)
		writeRes(f)
	}
}

var filterPaths = []string{
	"103.76.66.6",
	"103.76.66.2",
	"103.76.66.38",
	"103.76.66.90",
	"103.76.66.18",
	"103.76.66.10",
	"103.76.66.22",
	"103.76.66.14",
	"103.76.66.30",
	"103.76.66.74",
	"103.76.66.26",
	"103.76.66.66",
	"103.76.66.54",
}

func isFiltered(path string) bool {
	for _, filter := range filterPaths {
		if strings.Contains(path, filter) {
			return false
		}
	}
	return true
}

func writeRes(f *os.File) {
	w := csv.NewWriter(f)
	err := w.WriteAll([][]string{})
	if err != nil {
		panic(err.Error())
	}
	var tid string
	for sn, infos := range AllRelation {
		for _, info := range infos {
			tid = info.TID
			err = w.Write([]string{sn, info.PathInfo, info.Name})
			if err != nil {
				panic(err.Error())
			}
		}
	}
	w.Flush()
	fmt.Println(tid, "已写入")
	AllRelation = ResData{}
	//if err := w.Error(); err != nil {
	//	panic(err.Error())
	//}
}

type ResData map[string][]ResItem

type ResItem struct {
	PathInfo string
	Name     string
	TID      string
}

func genAllRelations(pathInfos map[string][]string) {
	f, err := os.Open("/Users/qijing.fqj/go/src/daily/csv/link_id_sn_company.csv")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	res, err := r.ReadAll()
	if err != nil {
		panic(err.Error())
	}
	for _, item := range res {
		if _, ok := AllRelation[item[1]]; ok {
			if paths, ok := pathInfos[item[0]]; ok {
				for _, path := range paths {
					AllRelation[item[1]] = append(AllRelation[item[1]], ResItem{
						PathInfo: path,
						Name:     item[2],
						TID:      item[0],
					})
				}
			}
			//}else{
			//	AllRelation[item[1]] = append(v, ResItem{
			//		Name:     item[2],
			//		TID:      item[0],
			//	})
			//}
		} else {
			if paths, ok := pathInfos[item[0]]; ok {
				AllRelation[item[1]] = []ResItem{}
				for _, path := range paths {
					AllRelation[item[1]] = append(AllRelation[item[1]], ResItem{
						PathInfo: path,
						Name:     item[2],
						TID:      item[0],
					})
				}
			}
		}
	}
}
