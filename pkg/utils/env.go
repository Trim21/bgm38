package utils

import "os"

//GetEnv Get environment variable, return fallback if not exist
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
