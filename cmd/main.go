package main

import (
	"yumikokawaii.iot.com/config"
	"yumikokawaii.iot.com/migration/migrate"
	"yumikokawaii.iot.com/server"
)

func main() {
	cfg, kongCtx := config.InitAppConfig()

	databaseUrl := cfg.GetDatabaseURL()
	cfg.MigrationFolder = cfg.GetMigrationFolder()

	switch kongCtx.Command() {
	case "serve":
		server.ServeServer(cfg)

	case "migrate <command>":
		switch cfg.Migrate.Command {
		case "up":
			migrate.Up(databaseUrl, cfg.MigrationFolder)
		}

	case "migrate <command> <option>":
		switch cfg.Migrate.Command {
		case "create":
			migrate.New(cfg.MigrationFolder, cfg.Migrate.Option)
		case "down":
			migrate.Down(databaseUrl, cfg.MigrationFolder, cfg.Migrate.Option)
		case "force":
			migrate.Force(databaseUrl, cfg.MigrationFolder, cfg.Migrate.Option)
		}
	}

}
