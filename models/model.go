package models

import (
	"github.com/ssdb/gossdb/ssdb"
	"cn/bkread/booktrans/conf"
	"log"
)

var (
	_Client *ssdb.Client
)

func init(){
	var err error
	_Client, err = ssdb.Connect(conf.DBCONFIG.Host, conf.DBCONFIG.Port)
	if err != nil {
		panic(err)
	}
	log.Println("Connetct to ssdb successed!")
}
