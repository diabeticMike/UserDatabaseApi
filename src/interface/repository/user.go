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
	FindUser(incomingUser models.User) (*models.User, error)
	FindAllUsers() ([]models.User, error)
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

// FindUser find user by parameters except id
func (ur *userRepository) FindUser(incomingUser models.User) (*models.User, error) {
	var user models.User
	err := ur.users.Find(bson.M{"email": incomingUser.Email,
		"last_name":  incomingUser.LastName,
		"country":    incomingUser.Country,
		"city":       incomingUser.City,
		"gender":     incomingUser.Gender,
		"birth_date": incomingUser.BirthDate}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAllUsers find all users
func (ur *userRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := ur.users.Find(bson.M{}).Limit(100).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
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
