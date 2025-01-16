package utils

import "fmt"

func BytesToHumanReadable(bytes int) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	index := 0
	value := float64(bytes)

	for value >= 1024 && index < len(sizes)-1 {
		value /= 1024
		index++
	}

	return fmt.Sprintf("%.2f %s", value, sizes[index])
}

