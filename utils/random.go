package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomPath() (string, error) {
	bytes := make([]byte, 2) // 2 bytes = 4 caracteres hexadecimales
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes)[:3], nil
}

