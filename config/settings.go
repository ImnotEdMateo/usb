package config

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

// Variables que se llenarán desde el .ini
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
		log.Fatalf("Error cargando config: %v", err)
	}

	MaxFileSize = cfg.Section("").Key("MaxFileSize").MustInt64(1 << 30)
	FileExpirationTime = time.Duration(cfg.Section("").Key("FileExpirationTimeHours").MustInt(1)) * time.Hour
	Doxxing = cfg.Section("").Key("Doxxing").MustBool(false)
	Theme = cfg.Section("").Key("Theme").MustString("cidoku.css")
	RandomPath = cfg.Section("").Key("RandomPath").MustString("GUID")
}
