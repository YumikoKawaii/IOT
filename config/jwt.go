package config

import "time"

type JWTConfig struct {
	SecretKey     string        `help:"jwt-secret-key" env:"JWT_SECRET" default:"YumikoSekaiDeIchibanKawaii"`
	TokenDuration time.Duration `help:"jwt-token-duration" env:"JWT_DURATION" default:"262800h"`
}
