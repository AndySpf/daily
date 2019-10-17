package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"time"
	"daily/hooks"

	"daily/g"

	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

type myFormat struct {
}

const (
	red    = 31
	green  = 32
	yellow = 33
	blue   = 36
	gray   = 37
)

func (m *myFormat) Format(entry *log.Entry) ([]byte, error) {

	b := &bytes.Buffer{}
	if entry.Buffer != nil {
		b = entry.Buffer
	}

	colorMap := map[log.Level]int{
		log.DebugLevel: green,
		log.WarnLevel:  yellow,
		log.InfoLevel:  blue,
		log.ErrorLevel: red,
		log.FatalLevel: red,
	}

	fmt.Fprintf(b, "%s\x1b[%dm[%s]\x1b[0m %s:%d %-44s ", entry.Time.Format("2006/01/02 15:04:05"), colorMap[entry.Level], entry.Level, entry.Caller.File, entry.Caller.Line, entry.Message)
	for key, value := range entry.Data {
		b.WriteString(key)
		b.WriteString("=")
		b.WriteString(value.(string))
		b.WriteByte(' ')
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func main() {
	err := g.InitConfig("/Users/qijing.fqj/go/src/uselog/config.json")
	if err != nil {
		panic(err.Error())
	}

	logger = log.New()
	logger.SetLevel(log.DebugLevel)
	logger.Formatter = &myFormat{}
	logger.ReportCaller = true
	file, err := os.OpenFile(path.Join(g.Cfg.LogConfig.LogDir, "log.1"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(err.Error())
	}
	logger.Out = file

	logger.AddHook(&hooks.MyHookCut{})
	//logger.Warn("111111111ewueo7819hqjwnfjdknjkfbjkebjkn")
	for {
		go func() {
			logger.Warn("111111111ewueo7819hqjwnfjdknjkfbjkebjkn")
		}()
		go func() {
			logger.Warn("222222222ewueo7819hqjwnfjdknjkfbjkebjkn")
		}()
		go func() {
			logger.Warn("333333333ewueo7819hqjwnfjdknjkfbjkebjkn")
		}()
		go func() {
			logger.Warn("444444444ewueo7819hqjwnfjdknjkfbjkebjkn")
		}()
		go func() {
			logger.Warn("555555555ewueo7819hqjwnfjdknjkfbjkebjkn")
		}()
		time.Sleep(time.Second * 1)
	}

}
