package config

import "time"
// DEFAULT CONFIG
const MaxFileSize = 1 << 30 // 1 GB
const FileExpirationTime = 1 * time.Hour // 1 Hour
const Doxxing = false // Show IP in Logs. Default: false
