package devices

type DeviceStat struct {
	Id   uint32
	Info string
}

func (DeviceStat) GetTableName() string {
	return "device_stats"
}
