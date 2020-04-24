package interactor

import (
	"github.com/UserDatabaseApi/src/interface/repository"
	"github.com/UserDatabaseApi/src/models"
)

type gameInteractor struct {
	GameRepository repository.GameRepository
}

// GameInteractor is inerface for working with game entity
type GameInteractor interface {
	GetGamesSortedByCreated() ([]models.Game, error)
}

func NewGameInteractor(gr repository.GameRepository) GameInteractor {
	return &gameInteractor{gr}
}

func (gi *gameInteractor) GetGamesSortedByCreated() ([]models.Game, error) {
	games, err := gi.GameRepository.FindGamesSortedByCreated()
	if err != nil {
		return nil, err
	}
	return games, nil
}
