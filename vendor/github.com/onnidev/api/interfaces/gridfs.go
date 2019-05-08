package interfaces

import (
	"github.com/onnidev/api/infra"
	"gopkg.in/mgo.v2"
)

// GridFSRepo is a struc that hold a mongo collection
type GridFSRepo struct {
	Session *mgo.Session
	FS      *mgo.GridFS
}

// NewGridFS creates a new CardDAO
func NewGridFS(store *infra.MongoStore) (GridFSRepo, error) {
	repo := GridFSRepo{
		Session: store.Session,
		FS:      store.Session.DB(store.Database).GridFS("onniFs"),
	}
	return repo, nil
}
