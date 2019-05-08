package infra

import (
	"errors"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/matryer/try.v1"
	"gopkg.in/mgo.v2"
)

var mongoGLOBAL MongoStore

// MongoStore TODO: NEEDS COMMENT INFO
type MongoStore struct {
	Database     string
	AuthDatabase string
	Session      *mgo.Session
	User         string
	Password     string
	Host         string
	Port         int
	Logger       Logger
}

// Mongo TODO: NEEDS COMMENT INFO
func Mongo(db string) {
	if viper.GetString("env") == "homolog" {
		db = "homolog"
	}
	mongoGLOBAL = MongoStore{
		Database:     db,
		AuthDatabase: "admin",
		User:         viper.GetString("MONGOUSER"),
		Password:     viper.GetString("MONGOPASS"),
		Host:         viper.GetString("MONGOHOST"),
		Port:         viper.GetInt("MONGOPORT"),
		Logger:       NewLogger("MONGO"),
	}
	mongoGLOBAL.retry()
}

// MongoDev TODO: NEEDS COMMENT INFO
func MongoDev(db string) {
	mongoGLOBAL = MongoStore{
		Database:     "onni",
		AuthDatabase: "admin",
		User:         "",
		Password:     "",
		Host:         "localhost",
		Port:         27017,
		Logger:       NewLogger("MONGO"),
	}
	mongoGLOBAL.retry()
}

//Cloner clones the local MONGO object
func Cloner() (*MongoStore, error) {

	tryFunc := func(attempt int) (bool, error) {
		if attempt == 3 {
			mongoGLOBAL.retry()
		}
		err := mongoGLOBAL.Session.Ping()
		if err != nil {
			mongoGLOBAL.Logger.Printf(" Failed ping on attempt %d\n", attempt)
			return attempt < 5, err // try 5 times
		}
		return true, nil // try 5 times
	}
	err := try.Do(tryFunc)
	if err != nil {
		return nil, errors.New("Failed to get database")
	}
	return &MongoStore{
		Database: mongoGLOBAL.Database,
		Logger:   mongoGLOBAL.Logger,
		Session:  mongoGLOBAL.Session.Clone(),
	}, nil
}

//InitMongo set MongoDB connection
func (store MongoStore) retry() {
	tryFunc := func(attempt int) (bool, error) {
		var err error
		mongoGLOBAL.Session, err = store.connect()
		if err != nil {
			mongoGLOBAL.Logger.Printf(" Failed on attempt %d\n", attempt)
			return attempt < 5, err // try 5 times
		}
		mongoGLOBAL.Session.SetMode(mgo.Monotonic, true)
		mongoGLOBAL.Logger.Printf(" Connected on  attempt %d\n", attempt)
		return true, nil // try 5 times
	}
	err := try.Do(tryFunc)
	if err != nil {
		mongoGLOBAL.Logger.Printf("error: %s", err)
	}
}

//Connect is a function that connects to MOngoDB
func (store MongoStore) connect() (*mgo.Session, error) {
	addr := []string{store.Host + ":" + strconv.Itoa(store.Port)}
	user := store.User
	password := store.Password
	if viper.GetString("env") == "homolog" || viper.GetString("env") == "production" {
		addr = []string{"mongo.default.svc.cluster.local"}
		user = "admin"
		password = "descriptor8"
	}
	inf := &mgo.DialInfo{
		Addrs:    addr,
		Database: store.AuthDatabase,
		Username: user,
		Password: password,
		Timeout:  20 * time.Second}
	session, err := mgo.DialWithInfo(inf)
	if err != nil {
		return session, err
	}
	session.SetSafe(&mgo.Safe{})
	if err = session.Ping(); err != nil {
		return session, err
	}
	return session, nil
}
