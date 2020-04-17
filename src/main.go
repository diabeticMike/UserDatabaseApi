package main

import (
	"fmt"
	baseLog "log"
	"net/http"

	"github.com/UserDatabaseApi/src/seeds"

	"github.com/UserDatabaseApi/src/interface/helper"

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

	userRepository := repository.NewUserRepository(db, Config.DatabaseName)
	userGameRepository := repository.NewUserGameRepository(db, Config.DatabaseName)
	userController := controller.NewUserController(
		interactor.NewUserInteractor(
			userRepository,
			userGameRepository,
		),
		helper.NewUserHelper(),
	)

	fmt.Println(Config)
	if users, err := seeds.RunUserSeeds(userRepository, Config.SeedsFilePaths.Users); err != nil {
		log.Fatalf("Error while providing user seeds, err: %s", err.Error())
	} else {
		if err := seeds.RunUserGamesSeeds(userGameRepository, users, Config.SeedsFilePaths.UserGames); err != nil {
			log.Fatalf("Error while providing userGames seeds, err: %s", err.Error())
		}
	}

	mainRouter := mux.NewRouter().StrictSlash(true)
	router.ApplyUserRoutes(mainRouter.PathPrefix("/users").Subrouter(), userController)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(Config.ListenPort, handlers.CORS(headers, methods, origins)(mainRouter)))
}
