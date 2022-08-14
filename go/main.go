/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/shifty11/dao-dao-notifier/cmd"
	"github.com/shifty11/dao-dao-notifier/database"
	"github.com/shifty11/dao-dao-notifier/log"
)

func main() {
	defer log.SyncLogger()
	defer database.Close()
	cmd.Execute()
}
