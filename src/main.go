package main

import (
	baseLog "log"

	"github.com/UserDatabaseApi/src/config"
	"github.com/UserDatabaseApi/src/infrastructure/datastore"
	"github.com/UserDatabaseApi/src/logger"
	"github.com/globalsign/mgo"
)

func main() {
	configFilePath := "config.json"
	var (
		Config config.Configuration
		log    logger.Logger
		err    error
	)

	if err, Config = config.Load(configFilePath); err != nil {
		baseLog.Fatal(err)
	}

	if err, log = logger.Load(Config.Log); err != nil {
		baseLog.Fatal(err)
	}

	db, err := datastore.NewDB(Config.MongoURL)
	if err != nil {
		log.Errorf("Incorrect db connection, err: %s", err.Error())
	}
	defer db.Close()
	db.SetMode(mgo.Monotonic, true)
}
