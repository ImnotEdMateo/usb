package config

import (
	"log"
	"time"

	"github.com/imnotedmateo/usb/utils" // importa utils
	"gopkg.in/ini.v1"
)

var (
	MaxFileSize        int64
	FileExpirationTime time.Duration
	Doxxing            bool
	Theme              string
	RandomPath         string
)

func LoadConfig(path string) {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	// MaxFileSize desde .ini
	sizeStr := cfg.Section("").Key("MaxFileSize").MustString("1GB")
	sizeBytes, err := utils.HumanReadableToBytes(sizeStr)
	if err != nil {
		log.Fatalf("Error in MaxFileSize: %v", err)
	}
	MaxFileSize = sizeBytes

	// FileExpirationTime desde .ini
	durationStr := cfg.Section("").Key("FileExpirationTime").MustString("1h")
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Fatalf("Error in FileExpirationTime: %v", err)
	}
	FileExpirationTime = duration

	Doxxing = cfg.Section("").Key("Doxxing").MustBool(false)
	Theme = cfg.Section("").Key("Theme").MustString("cidoku.css")
	RandomPath = cfg.Section("").Key("RandomPath").MustString("GUID")
}
