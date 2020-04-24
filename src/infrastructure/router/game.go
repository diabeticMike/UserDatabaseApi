package router

import (
	"github.com/UserDatabaseApi/src/interface/controller"
	"github.com/gorilla/mux"
)

func ApplyGameRoutes(router *mux.Router, controller controller.GameController) {
	router.Methods("GET").Path("/getGamesSortedByCreate").HandlerFunc(controller.GetAllGamesInfoSortedByCreated)
}
