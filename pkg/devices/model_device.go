package devices

type DeviceType string

// currently supported device
const (
	FAN DeviceType = "FAN"
	LED DeviceType = "LED"
)

type Device struct {
	Id    uint32
	Type  string
	Code  string
	Owner string
	Topic string
}
