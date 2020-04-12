package main

import (
	baseLog "log"
	"net/http"

	"github.com/UserDatabaseApi/src/infrastructure/router"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/UserDatabaseApi/src/interface/controller"
	"github.com/UserDatabaseApi/src/interface/interactor"
	"github.com/UserDatabaseApi/src/interface/repository"

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

	// Create service configuration
	if err, Config = config.Load(configFilePath); err != nil {
		baseLog.Fatal(err)
	}

	// Create service logger
	if err, log = logger.Load(Config.Log); err != nil {
		baseLog.Fatal(err)
	}

	// Set up database session
	db, err := datastore.NewDB(Config.MongoURL)
	if err != nil {
		log.Errorf("Incorrect db connection, err: %s", err.Error())
	}
	defer db.Close()
	db.SetMode(mgo.Monotonic, true)

	userController := controller.NewUserController(
		interactor.NewUserInteractor(
			repository.NewUserRepository(db, Config.DatabaseName),
		),
	)
	mainRouter := mux.NewRouter().StrictSlash(true)
	router.ApplyUserRoutes(mainRouter.PathPrefix("/users").Subrouter(), userController)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(Config.ListenPort, handlers.CORS(headers, methods, origins)(mainRouter)))
}
