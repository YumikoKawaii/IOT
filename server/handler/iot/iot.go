package iot

import (
	"context"
	pb "yumikokawaii.iot.com/pb"
	"yumikokawaii.iot.com/pkg/auth"
	"yumikokawaii.iot.com/pkg/userinfo"
)

type ServiceServer struct {
	pb.UnimplementedIOTServiceServer
	userinfoService userinfo.Service
	jwtResolver     auth.JWTResolver
}

func NewServiceServer(service userinfo.Service, resolver auth.JWTResolver) ServiceServer {
	return ServiceServer{
		userinfoService: service,
		jwtResolver:     resolver,
	}
}

func (s *ServiceServer) Register(ctx context.Context, request *pb.RegisterRequest) (error, *pb.RegisterResponse) {
	err := s.userinfoService.Register(request.Username, request.Password)
	if err != nil {
		return err, nil
	}
	token, _ := s.jwtResolver.GenerateJWTToken(request.Username)
	return nil, &pb.RegisterResponse{
		Code:    200,
		Message: "success",
		Token:   token,
	}
}

func (s *ServiceServer) Login(ctx context.Context, request *pb.LoginRequest) (error, *pb.LoginResponse) {
	err := s.userinfoService.Login(request.Username, request.Password)
	if err != nil {
		return err, nil
	}
	token, _ := s.jwtResolver.GenerateJWTToken(request.Username)
	return nil, &pb.LoginResponse{
		Code:    200,
		Message: "success",
		Token:   token,
	}
}

func (s *ServiceServer) Control(ctx context.Context, request *pb.ControlRequest) (error, *pb.ControlResponse) {
	return nil, nil
}
