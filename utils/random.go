package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/imnotedmateo/usb/config"
	"github.com/google/uuid"
)

func GenerateRandomPath() (string, error) {
	if config.RandomPath == "GUID" {
		return uuid.New().String(), nil
	}

	numChars, err := strconv.Atoi(config.RandomPath)
	if err != nil || numChars < 1 {
		return "", fmt.Errorf("invalid RandomPathSetting value: must be 'GUID' or a positive integer")
	}

	bytes := make([]byte, (numChars+1)/2)
	_, err = rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes)[:numChars], nil
}
