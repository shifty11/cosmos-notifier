package main

import (
	"github.com/shifty11/cosmos-notifier/cmd"
	"github.com/shifty11/cosmos-notifier/database"
	"github.com/shifty11/cosmos-notifier/log"
)

func main() {
	defer log.SyncLogger()
	defer database.Close()

	defer func() {
		if err := recover(); err != nil {
			log.Sugar.Panic(err)
			return
		}
	}()

	cmd.Execute()
}

