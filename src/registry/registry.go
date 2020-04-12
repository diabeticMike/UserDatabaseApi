package registry

import "github.com/globalsign/mgo"

type registry struct {
	db *mgo.Session
}

type Registry interface {
}

func NewRegistry(db *mgo.Session) Registry {
	return &registry{db}
}

//func (r *registry) NewAppController() controller.AppController {
//
//}
