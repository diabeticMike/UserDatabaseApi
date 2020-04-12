package controller

import (
	"fmt"
	"net/http"

	"github.com/UserDatabaseApi/src/interface/interactor"
)

type userController struct {
	UserInteractor interactor.UserInteractor
}

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{ui}
}

func (uc *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
