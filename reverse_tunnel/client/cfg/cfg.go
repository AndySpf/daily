package cfg

import (
	"encoding/json"
	"io/ioutil"
)

type ClientConfig struct {
	LogConfig         logConfig `json:"log_config"`
	ServerAddr        string    `json:"server_addr"`
	ReportProgressURL string    `json:"report_progress_url"`
	SshKeyUrl         sshKeyUrl `json:"ssh_key_url"`
}

type logConfig struct {
	LogPath string `json:"log_path"`
}

type sshKeyUrl struct {
	ClientPrivateKeyMips string `json:"client_private_key_mips"`
	ClientPrivateKeyX    string `json:"client_private_key_x"`
	ServerPublicKey      string `json:"server_public_key"`
}

var cfg = ClientConfig{}

func ParseConfig(path string) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, &cfg)
	if err != nil {
		return err
	}
	return nil
}

func Config() *ClientConfig {
	return &cfg
}
