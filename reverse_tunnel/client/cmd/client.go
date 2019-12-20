package main

import (
	"daily/reverse_tunnel/client/cfg"
	"daily/reverse_tunnel/client/socket"
	"daily/reverse_tunnel/client/task"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	cfgPath = flag.String("c", "./config.json", "config file path")
)

func main() {
	//cmd := `ssh -o StrictHostKeyChecking=no -fCNR 0.0.0.0:8003:localhost:22 root@134.175.68.57`
	//_, err := utils.ExecuteCMD(cmd)
	//if err != nil{
	//
	//}
	flag.Parse()

	if err := cfg.ParseConfig(*cfgPath); err != nil {
		panic("加载配置文件失败:" + err.Error())
	}

	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	// 从服务端下载客户端应该用的私钥,拿到公钥才能执行ssh反向隧道的命令
	// TODO 服务端公钥目前是刷机时直接刷上的，不需要程序去添加，但这样就限制了服务端不能随意变动，除非从unt或者rtunnel上拷贝秘钥对到新的服务器上。
	for {
		if err := downloadPrivateKey(); err != nil {
			log.Errorf("统一版本的客户端私钥无法获取,等待5s后重试:%s", err.Error())
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	socket.Run()
}

// 下载客户端私钥 并保存为临时文件，建立反向隧道时-i 指定该私钥文件
func downloadPrivateKey() error {
	resp, err := http.Get(fmt.Sprintf("http://%s%s", cfg.Config().ServerAddr, cfg.Config().SshKeyUrl.ClientPrivateKeyX))
	if err != nil {
		return err
	}
	f, err := os.Create(task.PrivateKeyPath)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	if err = os.Chmod(task.PrivateKeyPath, 0600); err != nil {
		return err
	}

	return nil
}
