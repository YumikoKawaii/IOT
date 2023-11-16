package config

import "github.com/alecthomas/kong"

var (
	appConfig *AppConfig
)

type AppConfig struct {
	GRPCPort int `name:"grpc-port" help:"GRPC port" env:"GRPC_PORT" default:"8081"`
	HTTPPort int `name:"http-port" help:"HTTP port" env:"HTTP_PORT" default:"8080"`

	JwtCfg JWTConfig `kong:"help:'jwt config',embed"`

	MysqlCfg MysqlConfig `kong:"help:'Mysql config',embed"`

	HashKey string `name:"hash key" env:"HASH_KEY" default:"YumikoKawaii"`
}

func initAppConfig() (*AppConfig, *kong.Context) {
	cfg := &AppConfig{}
	kongCtx := kong.Parse(cfg)
	return cfg, kongCtx
}

func LoadAppConfig() *AppConfig {
	if appConfig == nil {
		appConfig, _ = initAppConfig()
	}
	return appConfig
}
