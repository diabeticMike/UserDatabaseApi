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
	GetAllUserStatistics() ([]models.UserStats, error)
}

func NewUserInteractor(ur repository.UserRepository, ugr repository.UserGameRepository) UserInteractor {
	return &userInteractor{ur, ugr}
}

func (ui *userInteractor) GetAllUserStatistics() ([]models.UserStats, error) {
	userGames, err := ui.UserGameRepository.FindAllUserGames()
	if err != nil {
		return nil, err
	}

	users, err := ui.UserRepository.FindAllUsers()
	if err != nil {
		return nil, err
	}

	userMap := make(map[bson.ObjectId]models.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	usersStats := make([]models.UserStats, 0, len(users))
	for _, userGame := range userGames {
		usersStats = append(usersStats, models.UserStats{UserInfo: userMap[userGame.UserID], GamesCount: len(userGame.GameIDs)})
	}

	return usersStats, nil
}
