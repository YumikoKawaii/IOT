package config

import "fmt"

type MysqlConfig struct {
	Host     string `help:"hlidskjalf DB host" env:"HLIDSKJALF_DB_HOST" default:"localhost"`
	Port     int    `help:"hlidskjalf DB port" env:"HLIDSKJALF_DB_PORT" default:"3307"`
	DBName   string `help:"hlidskjalf DB name" env:"HLIDSKJALF_DB_NAME" default:"hlidskjalf"`
	Username string `help:"hlidskjalf DB username" env:"HLIDSKJALF_DB_USERNAME" default:"yumiko"`
	Password string `help:"hlidskjalf DB password" env:"HLIDSKJALF_DB_PASSWORD" default:"Yumiko1@"`
}

func (c *MysqlConfig) String() string {
	return fmt.Sprintf("mysql://%s", c.GetDns())
}

func (c *MysqlConfig) GetDns() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", c.Username, c.Password, c.Host, c.Port, c.DBName)
}

func NewMysqlConfig() *MysqlConfig {
	return &MysqlConfig{}
}
