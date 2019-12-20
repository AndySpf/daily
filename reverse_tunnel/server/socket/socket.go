package socket

import (
	"daily/reverse_tunnel/g"
	"daily/reverse_tunnel/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

type socketReq struct {
	Sn    string `json:"sn"`
	Token string `json:"token"`
}

type safeMembers struct {
	sync.RWMutex
	Members map[string]*item
}

type item struct {
	Conn     *websocket.Conn
	LastPong int64
	Closed   chan bool // 当前连接的存活状况
}

var onlineMembers = safeMembers{Members: make(map[string]*item)} // 全局map， server端http接口会对其进行数据下发

var (
	upgrader = websocket.Upgrader{
		// 读取存储空间大小
		ReadBufferSize: 1024,
		// 写入存储空间大小
		WriteBufferSize: 1024,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func checkAlive() {
	for {
		now := time.Now().Unix()
		onlineMembers.RLock()
		log.Debug("start check")
		for sn, conn := range onlineMembers.Members {
			if now-conn.LastPong > 20 {
				onlineMembers.Delete(sn)
			}
		}
		onlineMembers.RUnlock()
		time.Sleep(time.Second * 30) // 每半分钟清理一次长时间未收到pong信息的链接，认为这些链接已经断开了
	}
}

func init() {
	go checkAlive()
}

// SocketHandler ws server 服务端代码
func SocketHandler(c *gin.Context) {
	// 将http连接升级
	wbsCon, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error(err.Error()) // 升级ws连接失败直接返回
		c.JSON(500, g.Res{
			RetCode: 1,
			Desc:    "初始化ws失败：" + err.Error(),
		})
		return
	}
	defer wbsCon.Close()

	sn := ""
	retryTimes := 0
	wbsCon.SetPongHandler(func(appData string) error {
		conn := onlineMembers.Get(sn)
		if conn == nil {
			return fmt.Errorf("收到未注册客户端pong消息:%s", sn)
		}
		conn.LastPong = time.Now().Unix()
		return nil
	})
	for {
		// 只接收Text, Pong 类型的数据
		msgType, data, err := wbsCon.ReadMessage()
		if err != nil {
			if retryTimes >= 3 {
				log.Errorf("尝试读取消息失败超过%d次，重新建立连接", retryTimes)
				break
			}
			retryTimes++
			log.Errorf("读取消息失败:%s", err.Error())
			// 跳出循环不再更新pong消息即可，等待定时删除该连接（不直接删除的原因是，如果在定时间隔内客户端重连，则key为sn的对象可复用）
			continue
		}
		retryTimes = 0

		if msgType == websocket.TextMessage {
			msg := g.WsClientMsg{}
			err := json.Unmarshal(data, &msg)
			if err != nil {
				log.Errorf("unexpect msg:%s", utils.Bytes2str(data))
				continue
			}
			sn = msg.SN
			if onlineMembers.CreateOrUpdate(sn, wbsCon) {
				go ping(sn) // 新建连接后对其发送ping信息
			}
			log.Infof("%s上线", sn)
		}
	}
}

func (m *safeMembers) CreateOrUpdate(sn string, conn *websocket.Conn) (isNew bool) {
	m.Lock()
	defer m.Unlock()
	_, ok := m.Members[sn]
	m.Members[sn] = &item{
		Conn:     conn,
		LastPong: time.Now().Unix(),
		Closed:   make(chan bool),
	}
	return !ok
}

func (m *safeMembers) Get(key string) *item {
	m.RLock()
	defer m.RUnlock()
	return m.Members[key]
}

func (m *safeMembers) Delete(key string) {
	m.Members[key].Closed <- true
	log.Infof("%s下线", key)
	onlineMembers.Lock()
	// 直接删除，之前已经从onlineMembers取出来的连接不会因为删除字典项而变为nil
	delete(onlineMembers.Members, key)
	onlineMembers.Unlock()
}

func ping(sn string) {
	conn := onlineMembers.Get(sn)
	if conn == nil {
		// 连接为nil，则直接删除并返回。
		onlineMembers.Lock()
		delete(onlineMembers.Members, sn)
		onlineMembers.Unlock()
		return
	}
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case closed := <-conn.Closed:
			if closed {
				return
			}
		case <-ticker.C:
			if err := conn.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				// 发送失败可能是因为客户端有重连，原来的连接断了，因此需要取新的连接
				log.Errorf("对%s发送ping消息失败，重新获取conn尝试发送:%s", sn, err.Error())
				conn = onlineMembers.Get(sn)
				if conn == nil {
					// 连接为nil，则直接删除并返回。
					onlineMembers.Lock()
					delete(onlineMembers.Members, sn)
					onlineMembers.Unlock()
					return
				}
				continue
			}
			log.Debug("发送了ping")
		}
	}
}

func SendTextMessage(sn string, msg []byte) error {
	conn := onlineMembers.Get(sn)
	if conn == nil {
		return fmt.Errorf("%s设备不在线", sn)
	}
	return conn.Conn.WriteMessage(websocket.TextMessage, msg)
}

func GetOnlineMembers() map[string]interface{} {
	onlineMembers.RLock()
	defer onlineMembers.RUnlock()

	res := map[string]interface{}{}
	for sn, item := range onlineMembers.Members {
		res[sn] = item.LastPong
	}
	return res
}
