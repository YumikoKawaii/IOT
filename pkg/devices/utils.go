package devices

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"reflect"
	"yumikokawaii.iot.com/pkg/mqttresolver"
)

func Deserialize(data string, structure interface{}) (mqttresolver.Serializable, error) {

	message := reflect.New(reflect.TypeOf(structure)).Interface()
	err := json.Unmarshal([]byte(data), message)
	if err != nil {
		return nil, xerrors.Errorf("incorrect data structure")
	}
	return message.(mqttresolver.Serializable), nil
}

func GetMessageStructureFromDeviceType(deviceType string) interface{} {
	switch deviceType {
	case string(FAN):
		return mqttresolver.FanControlMessage{}
	case string(LED):
		return mqttresolver.LedControlMessage{}
	}
	return nil
}
