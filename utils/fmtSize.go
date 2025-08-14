package utils

import (
	"fmt"
	"strings"
)

func BytesToHumanReadable(bytes int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	index := 0
	value := float64(bytes)

	for value >= 1024 && index < len(sizes)-1 {
		value /= 1024
		index++
	}

	return fmt.Sprintf("%.2f %s", value, sizes[index])
}

func HumanReadableToBytes(s string) (int64, error) {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	units := map[string]int64{
		"B":  1,
		"KB": 1 << 10,
		"MB": 1 << 20,
		"GB": 1 << 30,
		"TB": 1 << 40,
	}

	var value float64
	var unit string
	n, err := fmt.Sscanf(s, "%f%s", &value, &unit)
	if err != nil || n != 2 {
		return 0, fmt.Errorf("formato inválido: %s", s)
	}

	multiplier, ok := units[unit]
	if !ok {
		return 0, fmt.Errorf("unidad inválida: %s", unit)
	}

	return int64(value * float64(multiplier)), nil
}
