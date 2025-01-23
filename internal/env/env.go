package env

import (
	"os"
	"strconv"
)

func GetEnvString(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func GetEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return valueInt
}
