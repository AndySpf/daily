package socket

import (
	"daily/reverse_tunnel/client/cfg"
	"daily/reverse_tunnel/client/task"
	"daily/reverse_tunnel/g"
	"daily/reverse_tunnel/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

var sn = "qijingfqj" // TODO 改为本机sn号

func Run() {
	// 无限尝试重连
	for {
		// 建立连接
		conn, err := getConn()
		if err != nil {
			log.Errorf("获取ws连接失败:%s", err.Error())
			time.Sleep(time.Second * 3)
			continue
		}
		log.Infof("重新获取连接")
		// 发送上线认证
		data, _ := json.Marshal(g.WsClientMsg{SN: sn})
		// TODO TIMEOUT
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Errorf("写入%s消息失败:%s", utils.Bytes2str(data), err.Error())
			time.Sleep(time.Second * 3)
			continue
		}
		log.Infof("已发送:%s", utils.Bytes2str(data))

		retryTimes := 0
		for {
			msgType, data, err := conn.ReadMessage()
			if err != nil {
				if retryTimes >= 3 {
					log.Errorf("尝试读取消息失败超过%d次，重新建立连接", retryTimes)
					break
				}
				retryTimes++
				log.Errorf("读取消息失败:%s", err.Error())
				continue
			}
			retryTimes = 0
			log.Infof("收到消息%s", utils.Bytes2str(data))

			if msgType == websocket.TextMessage {
				log.Infof("收到text消息:%s", utils.Bytes2str(data))
				msg := g.WsServerMsg{}
				err = json.Unmarshal(data, &msg)
				if err != nil {
					log.Errorf("收到客户端无效信息%s", utils.Bytes2str(data))
					continue
				}
				// 执行创建反向ssh的命令
				task.ExecTask(msg)
			}
		}
	}
}

func getConn() (*websocket.Conn, error) {
	conn, res, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s/ws", cfg.Config().ServerAddr), nil)
	if err != nil {
		log.Errorf("拨号失败%s", err.Error())
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("读取响应失败%s", err.Error())
	}
	log.Infof("响应为:%s", utils.Bytes2str(data))

	conn.SetPingHandler(func(appData string) error {
		//log.Infof("收到ping信息:%s", appData)
		if err = conn.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
			return err
		}
		return nil
	})

	return conn, nil
}
