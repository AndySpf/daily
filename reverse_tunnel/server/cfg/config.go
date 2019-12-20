package cfg

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfig struct {
	SSHConfig  sshConfig  `json:"ssh_config"`
	HttpConfig httpConfig `json:"http_config"`
}

type sshConfig struct {
	ReverseKeys reverseKeys `json:"reverse_keys"` // 保存在固定目录，所有客户端共用一套即可
	ForwardKeys forwardKeys `json:"forward_keys"` // 在.ssh下生成
	SshdPort    uint16      `json:"sshd_port"`
}

type reverseKeys struct {
	MIPSPublicKeyPath  string `json:"mips_public_key_path"`  // mips版本的公钥，应该写入服务端的authorized_keys
	MIPSPrivateKeyPath string `json:"mips_private_key_path"` // mips版本的私钥，提供给客户端
	XPublicKeyPath     string `json:"x_public_key_path"`     // x86版本的公钥，应该写入服务端的authorized_keys
	XPrivateKeyPath    string `json:"x_private_key_path"`    // x86版本的私钥，提供给客户端
}

type forwardKeys struct {
	PublicKeyPath  string `json:"public_key_path"`  // 服务端自己的公钥，提供给客户端写入客户端authorized_keys
	PrivateKeyPath string `json:"private_key_path"` // 客户端自己的私钥，保存在.ssh下即可
}

type httpConfig struct {
	BindAddr string `json:"bind_addr"`
}

var cfg = ServerConfig{}

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

func Config() *ServerConfig {
	return &cfg
}
