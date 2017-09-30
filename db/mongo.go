package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/batazor/go-bookmarks/utils"
	"gopkg.in/mgo.v2"
)

var log = logrus.New()

var (
	//	Session stores mongodb session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

func Connect() {
	// Get configuration ======================================================
	MONGO_URL := utils.getenv("MONGO_URL", "localhost/bookmarks")

	s, err := mgo.Dial(MONGO_URL)
	if err != nil {
		log.Panic("Fail connect to Mongo")
		panic(err)
	}
	s.SetSafe(&mgo.Safe{})
	log.Info("Success connect to MongoDB")
	Session = s
}
