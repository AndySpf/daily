package task

import (
	"bytes"
	"daily/reverse_tunnel/client/cfg"
	"daily/reverse_tunnel/g"
	"daily/reverse_tunnel/utils"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

var PrivateKeyPath = "./tmp_private_key"

func ExecTask(msg g.WsServerMsg) {
	format := `nohup ssh -p %d -o 'StrictHostKeyChecking=no' -o 'UserKnownHostsFile=/dev/null' -i %s -N -f -R 0.0.0.0:%d:localhost:%d root@%s > /dev/null 2>&1 &`
	res, err := utils.ExecuteCMD(fmt.Sprintf(format,
		msg.SshdPort,
		PrivateKeyPath,
		msg.ServerPort,
		msg.ClientPort,
		strings.Split(cfg.Config().ServerAddr, ":")[0]))
	if err != nil {
		// 执行失败报告
		log.Errorf("执行ssh命令失败:%s,%s", res, err.Error())
		reportToServer(g.FAILED, msg.TaskID)
	}
	log.Infof("执行完毕")
	reportToServer(g.SUCCESS, msg.TaskID)
}

func reportToServer(status g.ProgressStatus, taskID string) {
	data := g.Progress{Status: status, TaskID: taskID}
	bs, _ := json.Marshal(data)

	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s%s", cfg.Config().ServerAddr, cfg.Config().ReportProgressURL), bytes.NewBuffer(bs))
	if err != nil {
		log.Errorf("报告状态失败：%s", err.Error())
		return
	}

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("报告状态失败：%s", err.Error())
		return
	}
	log.Debug(res.Body)
}
