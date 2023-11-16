//go:build wireinject
// +build wireinject

package iot

import (
	"github.com/google/wire"
	"yumikokawaii.iot.com/config"
	"yumikokawaii.iot.com/db"
	"yumikokawaii.iot.com/pkg/auth"
	"yumikokawaii.iot.com/pkg/userinfo"
)

func Initialize(cfg *config.AppConfig) ServiceServer {
	panic(wire.Build(NewServiceServer, userinfo.NewService, userinfo.NewRepository, db.NewMySQLDB, auth.NewJWTResolver))
}
