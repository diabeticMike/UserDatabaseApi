package helper

import (
	"encoding/json"

	"github.com/UserDatabaseApi/src/models"
)

type gameHelper struct{}

type GameHelper interface {
	MarshalAllGames(games []models.Game) ([]byte, error)
}

func NewGameHelper() GameHelper {
	return &gameHelper{}
}

func (*gameHelper) MarshalAllGames(games []models.Game) ([]byte, error) {
	resp, err := json.Marshal(games)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
