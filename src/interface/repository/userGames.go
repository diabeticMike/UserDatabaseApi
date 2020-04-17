package repository

import (
	"github.com/UserDatabaseApi/src/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type userGameRepository struct {
	userGames *mgo.Collection
}

func NewUserGameRepository(db *mgo.Session, databaseName string) UserGameRepository {
	return &userGameRepository{db.DB(databaseName).C("user_games")}
}

// UserGamesRepository is interface for userGames entity
type UserGameRepository interface {
	FindUserGameByID(id bson.ObjectId) (*models.UserGame, error)
	InsertUserGames(userGames []models.UserGame) error
}

// FindUserGameByID find userGame by id
func (ugr *userGameRepository) FindUserGameByID(id bson.ObjectId) (*models.UserGame, error) {
	var userGame models.UserGame

	err := ugr.userGames.FindId(id).One(&userGame)
	if err != nil {
		return nil, err
	}

	return &userGame, nil
}

// InsertUserGames create userGame objects inside db
func (ugr *userGameRepository) InsertUserGames(userGames []models.UserGame) error {
	for _, u := range userGames {
		u.ID = bson.NewObjectId()
		if err := ugr.userGames.Insert(u); err != nil {
			return err
		}
	}

	return nil
}
