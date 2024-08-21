package utils

import "os"

func GetEnvOrDefault(key string, default_ string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return default_
	}

	return value
}
