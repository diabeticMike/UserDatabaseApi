package router

import (
	"github.com/UserDatabaseApi/src/interface/controller"
	"github.com/gorilla/mux"
)

func ApplyUserRoutes(router *mux.Router, controller controller.UserController) {
	router.Methods("GET").Path("/getAllUsersInfo").HandlerFunc(controller.GetAllUsersInfo)
}
