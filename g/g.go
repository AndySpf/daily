package g

import (
	"encoding/json"
	"io/ioutil"
)

// 全局可读对象
var Cfg *config

type config struct {
	LogConfig LogConfig `json:"log_config"`
}

type LogConfig struct {
	LogDir            string `json:"log_dir"`
	LogLevel          string `json:"log_level"`
	CompressFileCount int    `json:"compress_file_count"`
	LogFileCount      int    `json:"log_file_count"`
	LogFileUnit       string `json:"log_file_unit"`
	LogFileSize       int    `json:"log_file_size"`
}

func InitConfig(logPath string) error {
	bs, err := ioutil.ReadFile(logPath)
	if err != nil {
		return err
	}
	c := &config{}
	err = json.Unmarshal(bs, c)
	if err != nil {
		return err
	}
	Cfg = c
	return nil
}
