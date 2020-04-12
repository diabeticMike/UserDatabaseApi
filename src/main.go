package main

import (
	basicLog "log"

	"github.com/UserDatabaseApi/src/logger"

	"github.com/UserDatabaseApi/src/config"
)

func main() {
	configFilePath := "config.json"
	var (
		Config config.Configuration
		log    logger.Logger
		err    error
	)

	if err, Config = config.Load(configFilePath); err != nil {
		basicLog.Fatal(err)
	}

	if err, log = logger.Load(Config.Log); err != nil {
		basicLog.Fatal(err)
	}

	log.Error("Kolya lisiy")
}
