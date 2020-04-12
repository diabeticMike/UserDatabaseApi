package repository

import (
	"github.com/UserDatabaseApi/src/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type userRepository struct {
	users *mgo.Collection
}

func NewUserRepository(db *mgo.Session, databaseName string) UserRepository {
	return &userRepository{db.DB(databaseName).C("users")}
}

// UserRepository is interface for user entity
type UserRepository interface {
	FindUserByID(id bson.ObjectId) (*models.User, error)
}

// FindUserByID find user by id
func (ur *userRepository) FindUserByID(id bson.ObjectId) (*models.User, error) {

	var user models.User

	err := ur.users.FindId(id).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
