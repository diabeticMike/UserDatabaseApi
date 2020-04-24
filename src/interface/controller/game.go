package controller

import (
	"net/http"

	"github.com/UserDatabaseApi/src/interface/helper"

	"github.com/globalsign/mgo"

	"github.com/UserDatabaseApi/src/interface/interactor"
)

type gameController struct {
	Interactor interactor.GameInteractor
	Helper     helper.GameHelper
}

type GameController interface {
	GetAllGamesInfoSortedByCreated(w http.ResponseWriter, r *http.Request)
}

func NewGameController(gi interactor.GameInteractor, gh helper.GameHelper) GameController {
	return &gameController{gi, gh}
}

func (gc *gameController) GetAllGamesInfoSortedByCreated(w http.ResponseWriter, r *http.Request) {
	games, err := gc.Interactor.GetGamesSortedByCreated()
	if err != nil && err == mgo.ErrNotFound {
		w.WriteHeader(http.StatusNoContent)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := gc.Helper.MarshalAllGames(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
