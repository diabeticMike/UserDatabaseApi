package controller

import (
	"net/http"

	"github.com/UserDatabaseApi/src/models"

	"github.com/UserDatabaseApi/src/interface/helper"

	"github.com/UserDatabaseApi/src/interface/interactor"
)

type userController struct {
	Interactor interactor.UserInteractor
	Helper     helper.UserHelper
}

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

func NewUserController(ui interactor.UserInteractor, uh helper.UserHelper) UserController {
	return &userController{ui, uh}
}

func (uc *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := uc.Helper.MarshalAllUsers([]models.User{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
