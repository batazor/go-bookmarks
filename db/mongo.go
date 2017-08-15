package db

import (
	"gopkg.in/mgo.v2"
	"github.com/spf13/viper"
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

var (
	//	Session stores mongodb session
	MongoSession *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

func Connect() {
	MONGO_URL := viper.GetString("database.mongo.url")
	session, err := mgo.Dial(MONGO_URL)
	if err != nil {
		log.Panic("Fail connect to Mongo")
		panic(err)
	}

	MongoSession = session
}