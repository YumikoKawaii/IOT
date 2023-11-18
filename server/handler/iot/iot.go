package iot

import (
	"context"
	pb "yumikokawaii.iot.com/pb"
	"yumikokawaii.iot.com/pkg/auth"
	"yumikokawaii.iot.com/pkg/devices"
	"yumikokawaii.iot.com/pkg/userinfo"
)

type ServiceServer struct {
	pb.UnimplementedIOTServiceServer
	userinfoService userinfo.Service
	deviceService   devices.Service
	jwtResolver     auth.JWTResolver
}

func NewServiceServer(userinfoService userinfo.Service, deviceService devices.Service, resolver auth.JWTResolver) *ServiceServer {
	return &ServiceServer{
		userinfoService: userinfoService,
		deviceService:   deviceService,
		jwtResolver:     resolver,
	}
}

func (s *ServiceServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := s.userinfoService.Register(ConvertFromRegisterRequestToRegisterModel(request))
	if err != nil {
		return nil, err
	}
	token, _ := s.jwtResolver.GenerateJWTToken(request.Username)
	return &pb.RegisterResponse{
		Code:    200,
		Message: "success",
		Token:   token,
	}, nil
}

func (s *ServiceServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	err := s.userinfoService.Login(request.Username, HashString(request.Password))
	if err != nil {
		return nil, err
	}
	token, _ := s.jwtResolver.GenerateJWTToken(request.Username)
	return &pb.LoginResponse{
		Code:    200,
		Message: "success",
		Token:   token,
	}, nil
}

func (s *ServiceServer) Control(ctx context.Context, request *pb.ControlRequest) (*pb.ControlResponse, error) {
	if err := s.deviceService.ControlDevice(ctx, request.DeviceId, request.ControlData); err != nil {
		return nil, err
	}
	return &pb.ControlResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (s *ServiceServer) UpsertDevice(ctx context.Context, request *pb.UpsertDeviceRequest) (*pb.UpsertDeviceResponse, error) {
	if err := s.deviceService.UpsertDevice(ConvertUpsertDeviceRequestToDeviceModel(request)); err != nil {
		return nil, err
	}
	return &pb.UpsertDeviceResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (s *ServiceServer) GetDevices(ctx context.Context, request *pb.GetDevicesRequest) (*pb.GetDevicesResponse, error) {
	deviceModels, err := s.deviceService.GetDevicesByOwner(request.Username)
	if err != nil {
		return nil, err
	}

	devicesRes := make([]*pb.GetDevicesResponse_Device, 0)

	for _, device := range deviceModels {
		devicesRes = append(devicesRes, ConvertFromDeviceModelToProtoDevice(device))
	}

	return &pb.GetDevicesResponse{
		Code:    200,
		Message: "success",
		Devices: devicesRes,
	}, nil
}
