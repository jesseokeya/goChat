package database

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// DB holds all methods that can be applied to the database
type DB struct{}

// ConnectToDatabase returns a pointer to mgo.Session
// which is where the database function reside
func (d *DB) ConnectToDatabase() *mgo.Session {
	session, err := mgo.Dial(d.GetURL())
	if err != nil {
		log.Fatal(err)
	}
	return session
}

// GetURL returns a string which contains the url of the mongoDb database
func (d *DB) GetURL() string {
	url := os.Getenv("DB_URL")
	if len(url) == 0 {
		url = ""
	}
	return url
}

// GetDbName gets name if its an environment variable or returns default name
func (d *DB) GetDbName() string {
	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		dbName = "goChat"
	}
	return dbName
}
