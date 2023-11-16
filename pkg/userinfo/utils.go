package userinfo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"yumikokawaii.iot.com/config"
)

func HashString(message string) string {
	appConfig := config.LoadAppConfig()
	key := appConfig.HashKey
	keyBytes := []byte(key)
	messageBytes := []byte(message)

	// Create an HMAC using SHA-256
	h := hmac.New(sha256.New, keyBytes)
	h.Write(messageBytes)
	hmacInBytes := h.Sum(nil)

	// Encode the result as a hexadecimal string
	return hex.EncodeToString(hmacInBytes)
}
