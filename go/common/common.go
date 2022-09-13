package common

import (
	"github.com/shifty11/dao-dao-notifier/log"
	"os"
	"strconv"
)

func GetEnvX(varname string) string {
	value := os.Getenv(varname)
	if value == "" {
		log.Sugar.Panic(varname + " must be set")
	}
	return value
}

func GetEnvAsBoolX(varname string) bool {
	value := os.Getenv(varname)
	if value == "" {
		log.Sugar.Panic(varname + " must be set")
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Sugar.Panic(err)
	}
	return boolValue
}
