package devices

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"reflect"
	"yumikokawaii.iot.com/pkg/mqttpublisher"
)

func Deserialize(data string, structure interface{}) (mqttpublisher.Serializable, error) {

	message := reflect.New(reflect.TypeOf(structure)).Interface()
	err := json.Unmarshal([]byte(data), message)
	if err != nil {
		return nil, xerrors.Errorf("incorrect data structure")
	}
	return message.(mqttpublisher.Serializable), nil
}

func GetMessageStructureFromDeviceType(deviceType string) interface{} {
	switch deviceType {
	case string(FAN):
		return mqttpublisher.FanControlMessage{}
	case string(LED):
		return mqttpublisher.LedControlMessage{}
	}
	return nil
}
