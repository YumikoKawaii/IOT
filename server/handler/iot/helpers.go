package iot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"yumikokawaii.iot.com/config"
	pb "yumikokawaii.iot.com/pb"
	"yumikokawaii.iot.com/pkg/devices"
	"yumikokawaii.iot.com/pkg/userinfo"
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

func ConvertFromRegisterRequestToRegisterModel(request *pb.RegisterRequest) *userinfo.Account {
	return &userinfo.Account{
		Username: request.Username,
		Password: HashString(request.Password),
	}
}

func HashString(message string) string {
	appConfig := config.LoadAppConfig()
	key := appConfig.HashKey
	keyBytes := []byte(key)
	messageBytes := []byte(message)

	// Create an HMAC using SHA-256
	h := hmac.New(sha256.New, keyBytes)
	h.Write(messageBytes)
	hmacInBytes := h.Sum(nil)

	// Encode the result as a hexadecimal string
	return hex.EncodeToString(hmacInBytes)
}
