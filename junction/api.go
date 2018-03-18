package main

import (
	"strings"

	"github.com/jesseokeya/goChat/database"
	"github.com/jesseokeya/goChat/models"
	"gopkg.in/mgo.v2/bson"
)

var (
	api API
)

// API creates an api struct which connects logics with the database
type API struct {
	db database.DB
}

func main() {
	db := api.db
	s := db.ConnectToDatabase()
	dbName := strings.ToLower(db.GetDbName())
	// userDataTest creates new test user for database

	userDataTest := models.User{
		ID:             bson.NewObjectId(),
		FirstName:      "Jesse",
		LastName:       "Okeya",
		Description:    "Software Developer Intern At Qlik",
		AccountCreated: bson.Now(),
	}
	s.DB(dbName).C(dbName).Insert(userDataTest)
}
