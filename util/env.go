package util

import "os"

func Getenv(key string, def string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return def
}
