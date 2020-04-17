package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type UserGame struct {
	ID           bson.ObjectId `json:"-" bson:"_id"`
	UserID       bson.ObjectId `json:"-,omitempty" bson:"userId"`
	PointsGained int           `json:"points_gained" bson:"points_gained"`
	WinStatus    int           `json:"win_status" bson:"win_status"`
	GameType     int           `json:"game_type" bson:"game_type"`
	Created      time.Time     `json:"created" bson:"created"`
}
