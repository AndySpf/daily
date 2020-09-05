package convert

import (
	"daily/reverse_tunnel/convert/data"
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	tp := "ha" // type为HA
	msg := []byte(`{"agencyName":"杭州分公司","agencyId":303,"vip":"192.168.1.2","heartbeat_ip":"172.16.1.2", "mode":"router","heartbeat_local":"","heartbeat_remote":"","heartbeat_dev":"","vip_dev":""}`)
	ha := data.HaConfig{}
	for index := range Rel[tp] {
		Rel[tp][index].AddPriority(ha.Convert(msg), tp)
		result, err := Rel[tp][index].MakeUp()
		if err != nil {
			fmt.Println(err.Error()) // 整合文件失败是否重新处理该数据
		}
		fmt.Println(string(result))
	}
}
