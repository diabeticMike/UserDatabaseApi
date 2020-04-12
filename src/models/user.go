package models

import "github.com/globalsign/mgo/bson"

type User struct {
	ID        bson.ObjectId `json:"-" bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	LastName  string        `json:"last_name" bson:"last_name"`
	Country   string        `json:"country" bson:"country"`
	City      string        `json:"city" bson:"city"`
	Gender    string        `json:"gender" bson:"gender"`
	BirthDate string        `json:"birth_date" bson:"birth_date"`
}
