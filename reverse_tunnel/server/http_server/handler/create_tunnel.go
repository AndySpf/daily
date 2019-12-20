package handler

import (
	"daily/reverse_tunnel/g"
	"daily/reverse_tunnel/server/cfg"
	"daily/reverse_tunnel/server/socket"
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type createTunnelReq struct {
	SN         string `json:"sn"`
	ServerPort uint16 `json:"server_port"`
	ClientPort uint16 `json:"client_port"`
}

// CreateTunnelHandler 操作者请求创建隧道，通过socket将命令下发给对应设备
func CreateTunnelHandler(c *gin.Context) {
	req := &createTunnelReq{
		ClientPort: 22, // 默认22
	}
	if err := c.BindJSON(req); err != nil {
		log.Errorf(err.Error())
		c.JSON(400, g.Res{RetCode: 1, Desc: err.Error()})
		return
	}

	// 等待客户端上报进度，待客户端执行命令成功，且netstat可以看到server_port端口被sshd程序占用，则证明创建反向隧道成功
	taskID, _ := uuid.NewV4()
	progressMap.InitChan(taskID.String())
	if err := sendTask(req, taskID.String()); err != nil {
		log.Errorf("发送命令失败:%s", err.Error())
		c.JSON(500, g.Res{RetCode: 1, Desc: err.Error()})
		return
	}

	res := <-progressMap.Get(taskID.String())
	progressMap.Delete(taskID.String())

	if !res {
		c.JSON(500, g.Res{RetCode: 1, Desc: "创建反向隧道失败，请检查日志"})
		return
	}

	c.JSON(200, g.Res{RetCode: 0, Desc: "OK"})
	return
}

func sendTask(req *createTunnelReq, taskID string) error {
	msg := g.WsServerMsg{
		ServerPort: req.ServerPort,
		ClientPort: req.ClientPort,
		SshdPort:   cfg.Config().SSHConfig.SshdPort,
		TaskID:     taskID,
	}
	bs, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return socket.SendTextMessage(req.SN, bs)
}
