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
	InsertUsers(users []models.User) error
	InsertUser(user models.User) error
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

// InsertUsers create user objects inside db
func (ur *userRepository) InsertUsers(users []models.User) error {
	for _, u := range users {
		if err := ur.users.Insert(u); err != nil {
			return err
		}
	}

	return nil
}

// InsertUser create user object inside db
func (ur *userRepository) InsertUser(user models.User) error {
	if err := ur.users.Insert(user); err != nil {
		return err
	}

	return nil
}
