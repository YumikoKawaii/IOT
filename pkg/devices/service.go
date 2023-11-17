package devices

import (
	"context"
	"golang.org/x/xerrors"
	"yumikokawaii.iot.com/pkg/auth"
	"yumikokawaii.iot.com/pkg/mqttpublisher"
)

type Service interface {
	UpsertDevice(*Device) error
	GetDevicesByOwner(string) ([]Device, error)
	GetDeviceById(uint32) (*Device, error)
	ControlDevice(context.Context, uint32, string) error
}

type serviceImpl struct {
	repo       Repository
	mqttClient *mqttpublisher.MQTTClient
}

func NewService(repository Repository, client *mqttpublisher.MQTTClient) Service {
	return &serviceImpl{
		repo:       repository,
		mqttClient: client,
	}
}

func (s *serviceImpl) UpsertDevice(device *Device) error {
	return s.repo.UpsertDevice(device)
}

func (s *serviceImpl) GetDevicesByOwner(owner string) ([]Device, error) {
	return s.repo.GetDevicesByOwner(owner)
}

func (s *serviceImpl) GetDeviceById(id uint32) (*Device, error) {
	return s.repo.GetDeviceById(id)
}

func (s *serviceImpl) ControlDevice(ctx context.Context, deviceId uint32, controlData string) error {
	username := ctx.Value(auth.UsernameKey).(string)
	device, err := s.repo.GetDeviceById(deviceId)
	if err != nil {
		return err
	}
	if device.Owner != username {
		return xerrors.Errorf("error user doesn't have permission on device: %s", device.Code)
	}
	message, err := Deserialize(controlData, GetMessageStructureFromDeviceType(device.Type))
	if err != nil {
		return err
	}
	return s.mqttClient.Send(device.Topic, message)
}
