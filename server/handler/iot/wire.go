//go:build wireinject
// +build wireinject

package iot

import (
	"github.com/google/wire"
	"yumikokawaii.iot.com/config"
	"yumikokawaii.iot.com/db"
	"yumikokawaii.iot.com/pkg/auth"
	"yumikokawaii.iot.com/pkg/devices"
	"yumikokawaii.iot.com/pkg/mqttpublisher"
	"yumikokawaii.iot.com/pkg/userinfo"
)

func Initialize(cfg *config.AppConfig) *ServiceServer {
	panic(wire.Build(NewServiceServer, userinfo.NewService, userinfo.NewRepository, devices.NewService, devices.NewRepository, mqttpublisher.NewMQTTClient, db.NewMySQLDB, auth.NewJWTResolver))
}
