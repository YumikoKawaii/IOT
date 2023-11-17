package config

import (
	"fmt"
	"github.com/alecthomas/kong"
)

var (
	appConfig *AppConfig
)

type AppConfig struct {
	Migrate struct {
		Command string `kong:"arg,name:'command',enum:'up,down,force,create'"`
		Option  string `kong:"arg,optional,name:'option'"`
	} `kong:"cmd"`

	Serve struct{} `kong:"cmd"`

	GRPCPort int `name:"grpc-port" help:"GRPC port" env:"GRPC_PORT" default:"9901"`
	HTTPPort int `name:"http-port" help:"HTTP port" env:"HTTP_PORT" default:"8080"`

	MigrationFolder string `name:"migration-folder" help:"Migration folder" env:"PUBLISHER_SERVICE_MIGRATE_PATH" default:"migration"`

	JwtCfg JWTConfig `kong:"help:'jwt config',embed"`

	MysqlCfg MysqlConfig `kong:"help:'Mysql config',embed"`

	MqttCfg MQTTConfig `kong:"help:'mqtt config',embed"`

	HashKey string `name:"hash key" env:"HASH_KEY" default:"YumikoKawaii"`
}

func (c *AppConfig) GetMigrationFolder() string {
	return fmt.Sprintf("file://%s", c.MigrationFolder)
}

func (c *AppConfig) GetDatabaseURL() string {
	return c.MysqlCfg.String()
}

func InitAppConfig() (*AppConfig, *kong.Context) {
	cfg := &AppConfig{}
	kongCtx := kong.Parse(cfg)
	return cfg, kongCtx
}

func LoadAppConfig() *AppConfig {
	if appConfig == nil {
		appConfig, _ = InitAppConfig()
	}
	return appConfig
}
