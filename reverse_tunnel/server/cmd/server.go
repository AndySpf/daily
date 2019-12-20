package main

import (
	"bytes"
	"daily/reverse_tunnel/server/cfg"
	"daily/reverse_tunnel/server/http_server"
	"daily/reverse_tunnel/utils"
	"flag"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var (
	cfgPath = flag.String("c", "./config.json", "config file path")
)

func main() {

	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe("localhost:10000", nil))
	}()
	err := cfg.ParseConfig(*cfgPath)
	if err != nil {
		panic(err.Error())
	}

	checkAndHandleSSHKey() //检测ssh秘钥文件是否正常

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)

	http_server.Run()
}

func checkAndHandleSSHKey() {
	mipsPrivateKeyPath := cfg.Config().SSHConfig.ReverseKeys.MIPSPrivateKeyPath
	mipsPublicKeyPath := cfg.Config().SSHConfig.ReverseKeys.MIPSPublicKeyPath
	xPrivateKeyPath := cfg.Config().SSHConfig.ReverseKeys.XPrivateKeyPath
	xPublicKeyPath := cfg.Config().SSHConfig.ReverseKeys.XPublicKeyPath
	privateKeyPath := cfg.Config().SSHConfig.ForwardKeys.PrivateKeyPath
	publicKeyPath := cfg.Config().SSHConfig.ForwardKeys.PublicKeyPath

	if !utils.FileExist(publicKeyPath) {
		panic("正向ssh公钥文件不存在")
	}
	if !utils.FileExist(privateKeyPath) {
		panic("正向ssh私钥文件不存在")
	}

	if !utils.FileExist(mipsPublicKeyPath) {
		panic("反向mips版本ssh公钥文件不存在")
	}
	if !utils.FileExist(mipsPrivateKeyPath) {
		panic("反向mips版本ssh私钥文件不存在")
	}

	if !utils.FileExist(xPublicKeyPath) {
		panic("反向x86版本ssh公钥文件不存在")
	}
	if !utils.FileExist(xPrivateKeyPath) {
		panic("反向x86版本ssh私钥文件不存在")
	}

	// 将反向公钥写入服务端authorized_keys
	var authorizedKeys = "/root/.ssh/authorized_keys"
	if !utils.FileExist(authorizedKeys) {
		if _, err := os.Create(authorizedKeys); err != nil {
			panic("创建authorized_keys失败" + err.Error())
		}
	}

	mipsbs, err := ioutil.ReadFile(mipsPublicKeyPath)
	if err != nil {
		panic("读取mips客户端公钥失败" + err.Error())
	}
	xbs, err := ioutil.ReadFile(xPublicKeyPath)
	if err != nil {
		panic("读取normal客户端公钥失败" + err.Error())
	}

	bs, err := ioutil.ReadFile(authorizedKeys)
	if err != nil {
		panic("读取authorized_keys失败" + err.Error())
	}
	if !bytes.Contains(bs, mipsbs) {
		if err = utils.AppendFile(authorizedKeys, mipsbs); err != nil {
			panic("authorized_keys追加mips版公钥" + err.Error())
		}
	}
	if !bytes.Contains(bs, xbs) {
		if err = utils.AppendFile(authorizedKeys, xbs); err != nil {
			panic("authorized_keys追加normal版公钥" + err.Error())
		}
	}
}
