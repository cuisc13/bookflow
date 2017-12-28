package models

import (
	"github.com/ssdb/gossdb/ssdb"
	"cn/bkread/booktrans/conf"
	"log"
	"gopkg.in/mgo.v2"
	"fmt"
)

var (
	_Client *ssdb.Client
	_MgoSession *mgo.Session
)

func init(){
	var err error
	// _Client, err = ssdb.Connect(conf.DBCONFIG.Host, conf.DBCONFIG.Port)
	// if err != nil {
	// 	//panic(err)
	// }
	// log.Println("Connetct to ssdb successed!")

	_MgoSession, err = mgo.Dial(fmt.Sprintf("%s:%d", conf.MGOCONF.Host, conf.MGOCONF.Port))
	if err != nil {
		//panic(err)
	}
	log.Println("Connect to mongo successed!")

}
