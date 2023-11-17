package iot

import (
	"fmt"
	"github.com/google/uuid"
	pb "yumikokawaii.iot.com/pb"
	"yumikokawaii.iot.com/pkg/devices"
)

const (
	topicPrefix = "yumiko.kawaii."
)

func ConvertUpsertDeviceRequestToDeviceModel(request *pb.UpsertDeviceRequest) *devices.Device {
	return &devices.Device{
		Type:  request.Type,
		Code:  request.Code,
		Owner: request.Owner,
		Topic: fmt.Sprintf("%s%s", topicPrefix, uuid.New()),
	}
}

func ConvertFromDeviceModelToProtoDevice(device devices.Device) *pb.GetDevicesResponse_Device {
	return &pb.GetDevicesResponse_Device{
		Id:   device.Id,
		Type: device.Type,
		Code: device.Code,
	}
}
