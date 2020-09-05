package file

import (
	"encoding/json"
	"fmt"
)

type FileHandler interface {
	MakeUp() ([]byte, error)
	AddPriority(v interface{}, key string) error
}

type RunJSON map[string]interface{}

func NewRunJSON() RunJSON {
	return RunJSON{
		"ha":  "",
		"nat": "",
	}
}

func (p RunJSON) AddPriority(v interface{}, key string) error {
	if key == "" {
		p = v.(map[string]interface{})
	}
	p[key] = v
	return nil
}

func (p RunJSON) MakeUp() ([]byte, error) {
	for key := range p {
		if p[key] == nil {
			// DoHttpReq
			fmt.Println(fmt.Sprintf("向boss请求%s的数据", key))
		}
	}
	return p.Bytes(), nil
}

func (p RunJSON) Bytes() []byte {
	bs, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return bs
}
