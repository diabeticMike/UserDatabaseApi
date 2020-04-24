package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID        bson.ObjectId `json:"-" bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	LastName  string        `json:"last_name" bson:"last_name"`
	Country   string        `json:"country" bson:"country"`
	City      string        `json:"city" bson:"city"`
	Gender    string        `json:"gender" bson:"gender"`
	BirthDate time.Time     `json:"birth_date" bson:"birth_date"`
}

type UserStats struct {
	GamesCount int  `json:"games_count"`
	UserInfo   User `json:"user_info"`
}
