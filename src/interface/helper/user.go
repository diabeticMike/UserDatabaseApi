package helper

import (
	"encoding/json"

	"github.com/UserDatabaseApi/src/models"
)

type userHelper struct{}

type UserHelper interface {
	MarshalAllUsers(users []models.User) ([]byte, error)
}

func NewUserHelper() UserHelper {
	return &userHelper{}
}

func (*userHelper) MarshalAllUsers(users []models.User) ([]byte, error) {
	resp, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
