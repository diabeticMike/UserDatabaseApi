package controller

import (
	"net/http"

	"github.com/UserDatabaseApi/src/interface/interactor"
)

type userController struct {
	UserInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers() error
}

func (uc *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
}
