package file

import (
	"encoding/json"
)

type TunnelInfoJSON []map[string]interface{}

func NewTunnelInfoJSON() TunnelInfoJSON {
	return []map[string]interface{}{}
}

func (p TunnelInfoJSON) AddPriority(v interface{}, key string) error {
	if key == "" {
		p = v.([]map[string]interface{})
	}
	return nil
}

func (p TunnelInfoJSON) MakeUp() ([]byte, error) {
	return p.Bytes(), nil
}

func (p TunnelInfoJSON) Bytes() []byte {
	bs, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return bs
}
