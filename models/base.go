package models

import (
	"github.com/ssdb/gossdb/ssdb"
	"gopkg.in/mgo.v2"
)

type Db struct {
}

func (db *Db)Client()*ssdb.Client{
	return _Client
}
func (db *Db)MgoClient()*mgo.Session{
	return _MgoSession
}
