package common

import (
	"github.com/shifty11/cosmos-notifier/log"
	"os"
	"strconv"
)

// GetEnvX returns the value of the environment variable named by the key or panics if it is not set.
func GetEnvX(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Sugar.Panic(key + " must be set")
	}
	return value
}

// GetEnvAsBoolX returns the value of the environment variable named by the key as a bool and panics if it is not set.
func GetEnvAsBoolX(key string) bool {
	value := os.Getenv(key)
	if value == "" {
		log.Sugar.Panic(key + " must be set")
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Sugar.Panic(err)
	}
	return boolValue
}
