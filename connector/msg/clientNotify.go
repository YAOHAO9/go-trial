package msg

import "encoding/json"

// ClientNotify 主动推送的客户端通知
type ClientNotify struct {
	Route string
	Data  interface{}
}

// ToBytes To []byte
func (m ClientNotify) ToBytes() (data []byte) {
	data, _ = json.Marshal(m)
	return
}
