package mqttpublisher

import "encoding/json"

type Serializable interface {
	ToBytes() []byte
}

type FanControlMessage struct {
	Mode int32
}

func (f *FanControlMessage) ToBytes() []byte {
	data, _ := json.Marshal(f)
	return data
}

type LedControlMessage struct {
	Mode int32
}

func (l *LedControlMessage) ToBytes() []byte {
	data, _ := json.Marshal(l)
	return data
}
