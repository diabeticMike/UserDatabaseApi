package helper

import (
	"encoding/json"

	"github.com/UserDatabaseApi/src/models"
)

type userHelper struct{}

type UserHelper interface {
	MarshalAllUsersStats(users []models.UserStats) ([]byte, error)
}

func NewUserHelper() UserHelper {
	return &userHelper{}
}

func (*userHelper) MarshalAllUsersStats(users []models.UserStats) ([]byte, error) {
	resp, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
