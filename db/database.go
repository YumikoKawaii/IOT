package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"yumikokawaii.iot.com/config"
)

var (
	mysqlDB *gorm.DB
)

func NewMySQLDB(cfg *config.AppConfig) *gorm.DB {
	if mysqlDB == nil {
		dns := cfg.MysqlCfg.GetDns()
		db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			log.Fatalf("error connect to database: %w", err)
		}
		mysqlDB = db
		log.Println("connect to mysql success")
	}
	return mysqlDB
}
