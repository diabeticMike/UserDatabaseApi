package models

import "github.com/globalsign/mgo/bson"

// UserGame used as relation between user and games
type UserGame struct {
	ID      bson.ObjectId   `json:"-" bson:"_id"`
	GameIDs []bson.ObjectId `json:"-" bson:"gameIDs"`
	UserID  bson.ObjectId   `json:"-" bson:"userID"`
}
