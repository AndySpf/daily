package data

import "encoding/json"

type DataHandler interface {
	Convert(msg []byte) interface{}
}

type HaConfig struct {
}

func (p *HaConfig) Convert(msg []byte) interface{} {
	var res interface{}
	json.Unmarshal(msg, &res)
	return res
}
