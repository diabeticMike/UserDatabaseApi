package controller

import (
	"net/http"

	"github.com/globalsign/mgo"

	"github.com/UserDatabaseApi/src/interface/helper"

	"github.com/UserDatabaseApi/src/interface/interactor"
)

type userController struct {
	Interactor interactor.UserInteractor
	Helper     helper.UserHelper
}

type UserController interface {
	GetAllUsersInfo(w http.ResponseWriter, r *http.Request)
}

func NewUserController(ui interactor.UserInteractor, uh helper.UserHelper) UserController {
	return &userController{ui, uh}
}

func (uc *userController) GetAllUsersInfo(w http.ResponseWriter, r *http.Request) {
	usersInfo, err := uc.Interactor.GetAllUserStatistics()
	if err != nil && err == mgo.ErrNotFound {
		w.WriteHeader(http.StatusNoContent)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := uc.Helper.MarshalAllUsersStats(usersInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
