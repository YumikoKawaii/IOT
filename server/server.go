package server

import (
	"google.golang.org/grpc"
	"log"
	"yumikokawaii.iot.com/config"
	"yumikokawaii.iot.com/pkg/auth"
	"yumikokawaii.iot.com/server/handler/iot"
	"yumikokawaii.iot.com/services"
)

func ServeServer(cfg *config.AppConfig) {

	jwtResolver := auth.NewJWTResolver(cfg)

	authInterceptor := auth.NewAuthInterceptor(jwtResolver)

	serverCfg := services.NewConfig(cfg.GRPCPort, cfg.HTTPPort)

	sv := services.NewServer(serverCfg, grpc.ChainUnaryInterceptor(authInterceptor.Unary()))

	iotServer := iot.Initialize(cfg)

	if err := sv.Register(iotServer); err != nil {
		log.Fatalf("error register server: %w", err)
	}

	if err := sv.Serve(); err != nil {
		log.Fatalf("error failed to serve: %w", err)
	}

}
