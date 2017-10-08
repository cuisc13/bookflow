package models

import "github.com/ssdb/gossdb/ssdb"

type Db struct {
}

func (db *Db)Client()*ssdb.Client{
	return _Client
}
