package lib

import (
	"encoding/hex"
	"crypto/hmac"
	"crypto/sha256"
)

func HashByKey(message  string,key string) string {
	keyBytes := []byte(key)
	messageBytes := []byte(message)
	mac := hmac.New(sha256.New, keyBytes)
	mac.Write(messageBytes)
	hmacBytes := mac.Sum(nil)
	return hex.EncodeToString(hmacBytes)
}