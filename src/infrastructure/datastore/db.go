package datastore

import (
	"github.com/globalsign/mgo"
)

func NewDB(mongoURL string) (*mgo.Session, error) {
	db, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, err
	}
	db.SetMode(mgo.Monotonic, true)

	return db, nil
}
