package devices

import (
	"context"
	"fmt"
	mqttPkg "github.com/eclipse/paho.mqtt.golang"
	"golang.org/x/xerrors"
	"yumikokawaii.iot.com/pkg/auth"
	"yumikokawaii.iot.com/pkg/mqttresolver"
)

type Service interface {
	UpsertDevice(*Device) error
	GetDevicesByOwner(string) ([]Device, error)
	GetDeviceById(uint32) (*Device, error)
	ControlDevice(context.Context, uint32, string) error

	UpsertDeviceStat(*DeviceStat) error
	GetDeviceStatById(uint32) (*DeviceStat, error)
	HandleStatMessage(mqttPkg.Message) error
}

type serviceImpl struct {
	repo      Repository
	publisher *mqttresolver.Publisher
}

func NewService(repository Repository, publisher *mqttresolver.Publisher) Service {

	return &serviceImpl{
		repo:      repository,
		publisher: publisher,
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
	return s.publisher.Send(device.Topic, message)
}

func (s *serviceImpl) UpsertDeviceStat(stat *DeviceStat) error {
	return s.repo.UpsertDeviceStat(stat)
}

func (s *serviceImpl) GetDeviceStatById(id uint32) (*DeviceStat, error) {
	return s.repo.GetDeviceStatById(id)
}

func (s *serviceImpl) HandleStatMessage(message mqttPkg.Message) error {
	fmt.Println(message)
	return nil
}
