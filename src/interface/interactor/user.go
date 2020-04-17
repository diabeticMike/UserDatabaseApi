package interactor

import (
	"github.com/UserDatabaseApi/src/interface/repository"
	"github.com/UserDatabaseApi/src/models"
	"github.com/globalsign/mgo/bson"
)

type userInteractor struct {
	UserRepository     repository.UserRepository
	UserGameRepository repository.UserGameRepository
}

// UserInteractor is inerface for working with user entity
type UserInteractor interface {
	GetByID(id bson.ObjectId) (*models.User, error)
}

func NewUserInteractor(ur repository.UserRepository, ugr repository.UserGameRepository) UserInteractor {
	return &userInteractor{ur, ugr}
}

func (ui *userInteractor) GetByID(id bson.ObjectId) (*models.User, error) {
	u, err := ui.UserRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ui *userInteractor) GetUserStatistics(id bson.ObjectId) (*models.User, error) {
	u, err := ui.UserRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
