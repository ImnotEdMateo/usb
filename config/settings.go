package config

import "time"

// DEFAULT CONFIG
const (
  MaxFileSize = 1 << 30 // 1GB
  FileExpirationTime = 1 * time.Hour // 1 Hour
  Doxxing = false // By default: FALSE, it shows the IPs in logs 
  Theme = "dark.css" // Themes located in /static/themes/ 
  RandomPath = "GUID" // GUID or an entire number 
)
