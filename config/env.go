package config

import (
	"os"
)

func GetENV(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable " + key + " not set")
	}
	return value
}
