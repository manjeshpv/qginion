package dbconfig

import (
	"gopkg.in/mgo.v2"
	"time"
)

//const MongoDb details
const (
	hosts      = "ds063439.mongolab.com:63439"
	database   = "tsc"
	username   = "qnotify"
	password   = "quezx123"
	collection = "messages"
)

type DB struct {
	Session *mgo.Session
}

func (db *DB) DoDial() (s *mgo.Session, err error) {
	//	dburl := os.Getenv("MONGOHQ_URL")
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	return mgo.DialWithInfo(info)
}

func (db *DB) Name() string {
	return "tsc"
}
